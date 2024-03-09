// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"

	"cogentcore.org/core/core/config"
	"cogentcore.org/core/xe"
	"github.com/mitchellh/go-homedir"
)

// Setup installs platform-specific dependencies for the current platform.
// It only needs to be called once per system.
func Setup(c *config.Config) error { //gti:add
	vc := xe.Verbose().SetBuffer(false)
	switch runtime.GOOS {
	case "darwin":
		p, err := xe.Output("xcode-select", "-p")
		if err != nil || p == "" {
			err := vc.Run("xcode-select", "--install")
			if err != nil {
				return err
			}
		}
		err = vc.Run("curl", "-O", "https://sdk.lunarg.com/sdk/download/1.3.261.1/mac/vulkansdk-macos-1.3.261.1.dmg")
		if err != nil {
			return err
		}
		err = xe.Run("sudo", "hdiutil", "attach", "vulkansdk-macos-1.3.261.1.dmg")
		if err != nil {
			return err
		}
		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		root := filepath.Join(home, "VulkanSDK", "1.3.261.1")
		err = vc.Run("sudo", "/Volumes/vulkansdk-macos-1.3.261.1/InstallVulkan.app/Contents/MacOS/InstallVulkan", "--root", root, "--accept-licenses", "--default-answer", "--confirm-command", "install", "com.lunarg.vulkan.core", "com.lunarg.vulkan.usr", "com.lunarg.vulkan.sdl2", "com.lunarg.vulkan.glm", "com.lunarg.vulkan.volk", "com.lunarg.vulkan.vma")
		if err != nil {
			return err
		}
		return nil
	case "linux":
		_, err := exec.LookPath("apt-get")
		if err == nil {
			return vc.Run("sudo", "apt-get", "install", "libgl1-mesa-dev", "xorg-dev")
		}
		_, err = exec.LookPath("dnf")
		if err == nil {
			return vc.Run("sudo", "dnf", "install", "libX11-devel", "libXcursor-devel", "libXrandr-devel", "libXinerama-devel", "mesa-libGL-devel", "libXi-devel", "libXxf86vm-devel")
		}
		return fmt.Errorf("unknown Linux distro (apt-get and dnf not found); file an issue at https://github.com/cogentcore/core/issues")
	}
	return fmt.Errorf("platform %q not supported for core setup", runtime.GOOS)
}
