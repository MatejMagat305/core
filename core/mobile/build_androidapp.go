// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mobile

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"cogentcore.org/core/core/config"
	"cogentcore.org/core/core/mobile/binres"
	"cogentcore.org/core/core/rendericon"
	"cogentcore.org/core/grog"
	"cogentcore.org/core/xe"
	"golang.org/x/tools/go/packages"
)

const (
	MinAndroidSDK           = 23
	DefaultAndroidTargetSDK = 29
)

// GoAndroidBuild builds the given package for the given Android targets.
func GoAndroidBuild(c *config.Config, pkg *packages.Package, targets []config.Platform) (map[string]bool, error) {
	ndkRoot, err := NDKRoot(c, targets...)
	if err != nil {
		return nil, err
	}
	libName := AndroidPkgName(c.Name)

	// TODO(hajimehoshi): This works only with Go tools that assume all source files are in one directory.
	// Fix this to work with other Go tools.
	dir := filepath.Dir(pkg.GoFiles[0])

	manifestPath := filepath.Join(dir, "AndroidManifest.xml")
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}

		buf := new(bytes.Buffer)
		buf.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
		err := ManifestTmpl.Execute(buf, ManifestTmplData{
			JavaPkgPath: c.ID,
			Name:        c.Name,
			LibName:     libName,
		})
		if err != nil {
			return nil, err
		}
		manifestData = buf.Bytes()
		grog.PrintfDebug("generated AndroidManifest.xml:\n%s\n", manifestData)
	} else {
		libName, err = ManifestLibName(manifestData)
		if err != nil {
			return nil, fmt.Errorf("error parsing %s: %v", manifestPath, err)
		}
	}

	libFiles := []string{}
	nmpkgs := make(map[string]map[string]bool) // map: arch -> extractPkgs' output

	for _, t := range targets {
		toolchain := NDK.Toolchain(t.Arch)
		libPath := "lib/" + toolchain.ABI + "/lib" + libName + ".so"
		libAbsPath := filepath.Join(TmpDir, libPath)
		if err := xe.MkdirAll(filepath.Dir(libAbsPath), 0755); err != nil {
			return nil, err
		}
		err = GoBuild(
			c,
			pkg.PkgPath,
			AndroidEnv[t.Arch],
			"-buildmode=c-shared",
			"-o", libAbsPath,
		)
		if err != nil {
			return nil, err
		}
		nmpkgs[t.Arch], err = ExtractPkgs(c, toolchain.Path(c, ndkRoot, "nm"), libAbsPath)
		if err != nil {
			return nil, err
		}
		libFiles = append(libFiles, libPath)
	}

	block, _ := pem.Decode([]byte(DebugCert))
	if block == nil {
		return nil, errors.New("no debug cert")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	if c.Build.Output == "" {
		c.Build.Output = filepath.Join(".core", "bin", "android", c.Name+".apk")
	}
	if !strings.HasSuffix(c.Build.Output, ".apk") {
		return nil, fmt.Errorf("output file name %q does not end in '.apk'", c.Build.Output)
	}
	err = os.MkdirAll(filepath.Dir(c.Build.Output), 0777)
	if err != nil {
		return nil, err
	}
	var out io.Writer
	if !c.Build.PrintOnly {
		f, err := os.Create(c.Build.Output)
		if err != nil {
			return nil, err
		}
		defer func() {
			if cerr := f.Close(); err == nil {
				err = cerr
			}
		}()
		out = f
	}

	var apkw *Writer
	if !c.Build.PrintOnly {
		apkw = NewWriter(out, privKey)
	}
	apkwCreate := func(name string) (io.Writer, error) {
		grog.PrintfInfo("apk: %s\n", name)
		if c.Build.PrintOnly {
			return io.Discard, nil
		}
		return apkw.Create(name)
	}
	apkwWriteFile := func(dst, src string) error {
		w, err := apkwCreate(dst)
		if err != nil {
			return err
		}
		if !c.Build.PrintOnly {
			f, err := os.Open(src)
			if err != nil {
				return err
			}
			defer f.Close()
			if _, err := io.Copy(w, f); err != nil {
				return err
			}
		}
		return nil
	}

	// TODO: do we need this writer stuff?
	w, err := apkwCreate("classes.dex")
	if err != nil {
		return nil, err
	}
	dexData, err := base64.StdEncoding.DecodeString(dexStr)
	if err != nil {
		log.Fatalf("internal error: bad dexStr: %v", err)
	}
	if _, err := w.Write(dexData); err != nil {
		return nil, err
	}

	for _, libFile := range libFiles {
		if err := apkwWriteFile(libFile, filepath.Join(TmpDir, libFile)); err != nil {
			return nil, err
		}
	}

	// TODO: what should we do about OpenAL?
	for _, t := range targets {
		toolchain := NDK.Toolchain(t.Arch)
		if nmpkgs[t.Arch]["cogentcore.org/core/mobile/exp/audio/al"] {
			dst := "lib/" + toolchain.ABI + "/libopenal.so"
			src := filepath.Join(GoMobilePath, dst)
			if _, err := os.Stat(src); err != nil {
				return nil, errors.New("the Android requires the golang.org/x/mobile/exp/audio/al, but the OpenAL libraries was not found. Please run gomobile init with the -openal flag pointing to an OpenAL source directory")
			}
			if err := apkwWriteFile(dst, src); err != nil {
				return nil, err
			}
		}
	}

	// Add the icon. 512 is the largest icon size on Android
	// (for the Google Play Store icon).
	ic, err := rendericon.Render(512)
	if err != nil {
		return nil, err
	}

	bxml, err := binres.UnmarshalXML(bytes.NewReader(manifestData), true, c.Build.AndroidMinSDK, c.Build.AndroidTargetSDK)
	if err != nil {
		return nil, err
	}

	// generate resources.arsc identifying single xxxhdpi icon resource.
	pkgname, err := bxml.RawValueByName("manifest", xml.Name{Local: "package"})
	if err != nil {
		return nil, err
	}
	tbl, name := binres.NewMipmapTable(pkgname)
	iw, err := apkwCreate(name)
	if err != nil {
		return nil, err
	}
	err = png.Encode(iw, ic)
	if err != nil {
		return nil, err
	}
	resw, err := apkwCreate("resources.arsc")
	if err != nil {
		return nil, err
	}
	rbin, err := tbl.MarshalBinary()
	if err != nil {
		return nil, err
	}
	if _, err := resw.Write(rbin); err != nil {
		return nil, err
	}

	w, err = apkwCreate("AndroidManifest.xml")
	if err != nil {
		return nil, err
	}
	bin, err := bxml.MarshalBinary()
	if err != nil {
		return nil, err
	}
	if _, err := w.Write(bin); err != nil {
		return nil, err
	}

	// TODO: add gdbserver to apk?

	if !c.Build.PrintOnly {
		if err := apkw.Close(); err != nil {
			return nil, err
		}
	}

	// TODO: return nmpkgs
	return nmpkgs[targets[0].Arch], nil
}

// AndroidPkgName sanitizes the go package name to be acceptable as a android
// package name part. The android package name convention is similar to the
// java package name convention described in
// https://docs.oracle.com/javase/specs/jls/se8/html/jls-6.html#jls-6.5.3.1
// but not exactly same.
func AndroidPkgName(name string) string {
	var res []rune
	for _, r := range name {
		switch {
		case 'a' <= r && r <= 'z', 'A' <= r && r <= 'Z', '0' <= r && r <= '9':
			res = append(res, r)
		default:
			res = append(res, '_')
		}
	}
	if len(res) == 0 || res[0] == '_' || ('0' <= res[0] && res[0] <= '9') {
		// Android does not seem to allow the package part starting with _.
		res = append([]rune{'g', 'o'}, res...)
	}
	s := string(res)
	// Look for Java keywords that are not Go keywords, and avoid using
	// them as a package name.
	//
	// This is not a problem for normal Go identifiers as we only expose
	// exported symbols. The upper case first letter saves everything
	// from accidentally matching except for the package name.
	//
	// Note that basic type names (like int) are not keywords in Go.
	switch s {
	case "abstract", "assert", "boolean", "byte", "catch", "char", "class",
		"do", "double", "enum", "extends", "final", "finally", "float",
		"implements", "instanceof", "int", "long", "native", "private",
		"protected", "public", "short", "static", "strictfp", "super",
		"synchronized", "this", "throw", "throws", "transient", "try",
		"void", "volatile", "while":
		s += "_"
	}
	return s
}

// A random uninteresting private key.
// Must be consistent across builds so newer app versions can be installed.
const DebugCert = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAy6ItnWZJ8DpX9R5FdWbS9Kr1U8Z7mKgqNByGU7No99JUnmyu
NQ6Uy6Nj0Gz3o3c0BXESECblOC13WdzjsH1Pi7/L9QV8jXOXX8cvkG5SJAyj6hcO
LOapjDiN89NXjXtyv206JWYvRtpexyVrmHJgRAw3fiFI+m4g4Qop1CxcIF/EgYh7
rYrqh4wbCM1OGaCleQWaOCXxZGm+J5YNKQcWpjZRrDrb35IZmlT0bK46CXUKvCqK
x7YXHgfhC8ZsXCtsScKJVHs7gEsNxz7A0XoibFw6DoxtjKzUCktnT0w3wxdY7OTj
9AR8mobFlM9W3yirX8TtwekWhDNTYEu8dwwykwIDAQABAoIBAA2hjpIhvcNR9H9Z
BmdEecydAQ0ZlT5zy1dvrWI++UDVmIp+Ve8BSd6T0mOqV61elmHi3sWsBN4M1Rdz
3N38lW2SajG9q0fAvBpSOBHgAKmfGv3Ziz5gNmtHgeEXfZ3f7J95zVGhlHqWtY95
JsmuplkHxFMyITN6WcMWrhQg4A3enKLhJLlaGLJf9PeBrvVxHR1/txrfENd2iJBH
FmxVGILL09fIIktJvoScbzVOneeWXj5vJGzWVhB17DHBbANGvVPdD5f+k/s5aooh
hWAy/yLKocr294C4J+gkO5h2zjjjSGcmVHfrhlXQoEPX+iW1TGoF8BMtl4Llc+jw
lKWKfpECgYEA9C428Z6CvAn+KJ2yhbAtuRo41kkOVoiQPtlPeRYs91Pq4+NBlfKO
2nWLkyavVrLx4YQeCeaEU2Xoieo9msfLZGTVxgRlztylOUR+zz2FzDBYGicuUD3s
EqC0Wv7tiX6dumpWyOcVVLmR9aKlOUzA9xemzIsWUwL3PpyONhKSq7kCgYEA1X2F
f2jKjoOVzglhtuX4/SP9GxS4gRf9rOQ1Q8DzZhyH2LZ6Dnb1uEQvGhiqJTU8CXxb
7odI0fgyNXq425Nlxc1Tu0G38TtJhwrx7HWHuFcbI/QpRtDYLWil8Zr7Q3BT9rdh
moo4m937hLMvqOG9pyIbyjOEPK2WBCtKW5yabqsCgYEAu9DkUBr1Qf+Jr+IEU9I8
iRkDSMeusJ6gHMd32pJVCfRRQvIlG1oTyTMKpafmzBAd/rFpjYHynFdRcutqcShm
aJUq3QG68U9EAvWNeIhA5tr0mUEz3WKTt4xGzYsyWES8u4tZr3QXMzD9dOuinJ1N
+4EEumXtSPKKDG3M8Qh+KnkCgYBUEVSTYmF5EynXc2xOCGsuy5AsrNEmzJqxDUBI
SN/P0uZPmTOhJIkIIZlmrlW5xye4GIde+1jajeC/nG7U0EsgRAV31J4pWQ5QJigz
0+g419wxIUFryGuIHhBSfpP472+w1G+T2mAGSLh1fdYDq7jx6oWE7xpghn5vb9id
EKLjdwKBgBtz9mzbzutIfAW0Y8F23T60nKvQ0gibE92rnUbjPnw8HjL3AZLU05N+
cSL5bhq0N5XHK77sscxW9vXjG0LJMXmFZPp9F6aV6ejkMIXyJ/Yz/EqeaJFwilTq
Mc6xR47qkdzu0dQ1aPm4XD7AWDtIvPo/GG2DKOucLBbQc2cOWtKS
-----END RSA PRIVATE KEY-----
`