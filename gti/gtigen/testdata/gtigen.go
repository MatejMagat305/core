// Code generated by "gtigen.test -test.paniconexit0 -test.timeout=10m0s"; DO NOT EDIT.

package testdata

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

// PersonType is the [gti.Type] for [Person]
var PersonType = gti.AddType(&gti.Type{
	Name:      "goki.dev/gti/gtigen/testdata.Person",
	ShortName: "testdata.Person",
	IDName:    "person",
	Doc:       "Person represents a person and their attributes.\nThe zero value of a Person is not valid.",
	Directives: gti.Directives{
		{Tool: "ki", Directive: "flagtype", Args: []string{"NodeFlags", "-field", "Flag"}},
		{Tool: "goki", Directive: "embedder"},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Name", &gti.Field{Name: "Name", Type: "string", LocalType: "string", Doc: "Name is the name of the person"}},
		{"Age", &gti.Field{Name: "Age", Type: "int", LocalType: "int", Doc: "Age is the age of the person"}},
		{"Type", &gti.Field{Name: "Type", Type: "*goki.dev/gti.Type", LocalType: "*gti.Type", Doc: "Type is the type of the person"}},
		{"Nicknames", &gti.Field{Name: "Nicknames", Type: "[]string", LocalType: "[]string", Doc: "Nicknames are the nicknames of the person"}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"RGBA", &gti.Field{Name: "RGBA", Type: "image/color.RGBA", LocalType: "color.RGBA", Doc: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{
		{"Introduction", &gti.Method{Name: "Introduction", Doc: "Introduction returns an introduction for the person.\nIt contains the name of the person and their age.", Directives: gti.Directives{
			{Tool: "gi", Directive: "toolbar", Args: []string{"-name", "ShowIntroduction", "-icon", "play", "-show-result", "-confirm"}},
			{Tool: "gti", Directive: "add"},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"string", &gti.Field{Name: "string", Type: "string", LocalType: "string", Doc: ""}},
		})}},
	}),
	Instance: &Person{},
})

func (t *Person) MyCustomFuncForStringers(a any) error {
	return nil
}

// SetName sets the [Person.Name]:
// Name is the name of the person
func (t *Person) SetName(v string) *Person {
	t.Name = v
	return t
}

// SetAge sets the [Person.Age]:
// Age is the age of the person
func (t *Person) SetAge(v int) *Person {
	t.Age = v
	return t
}

// SetType sets the [Person.Type]:
// Type is the type of the person
func (t *Person) SetType(v *gti.Type) *Person {
	t.Type = v
	return t
}

// SetNicknames sets the [Person.Nicknames]:
// Nicknames are the nicknames of the person
func (t *Person) SetNicknames(v ...string) *Person {
	t.Nicknames = v
	return t
}

// SetR sets the [Person.R]
func (t *Person) SetR(v uint8) *Person {
	t.R = v
	return t
}

// SetG sets the [Person.G]
func (t *Person) SetG(v uint8) *Person {
	t.G = v
	return t
}

// SetB sets the [Person.B]
func (t *Person) SetB(v uint8) *Person {
	t.B = v
	return t
}

// SetA sets the [Person.A]
func (t *Person) SetA(v uint8) *Person {
	t.A = v
	return t
}

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.Alert",
	Doc:        "Alert prints an alert with the given message",
	Directives: gti.Directives{},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"msg", &gti.Field{Name: "msg", Type: "string", LocalType: "string", Doc: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs0",
	Doc:        "",
	Directives: gti.Directives{},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"x", &gti.Field{Name: "x", Type: "float32", LocalType: "float32", Doc: ""}},
		{"y", &gti.Field{Name: "y", Type: "float32", LocalType: "float32", Doc: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs1",
	Doc:        "",
	Directives: gti.Directives{},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"x", &gti.Field{Name: "x", Type: "int", LocalType: "int", Doc: ""}},
		{"y", &gti.Field{Name: "y", Type: "struct{}", LocalType: "struct{}", Doc: ""}},
		{"z", &gti.Field{Name: "z", Type: "struct{}", LocalType: "struct{}", Doc: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs2",
	Doc:        "",
	Directives: gti.Directives{},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"x", &gti.Field{Name: "x", Type: "int", LocalType: "int", Doc: ""}},
		{"y", &gti.Field{Name: "y", Type: "int", LocalType: "int", Doc: ""}},
		{"z", &gti.Field{Name: "z", Type: "int", LocalType: "int", Doc: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs3",
	Doc:        "",
	Directives: gti.Directives{},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"x", &gti.Field{Name: "x", Type: "int", LocalType: "int", Doc: ""}},
		{"y", &gti.Field{Name: "y", Type: "bool", LocalType: "bool", Doc: ""}},
		{"z", &gti.Field{Name: "z", Type: "bool", LocalType: "bool", Doc: ""}},
		{"w", &gti.Field{Name: "w", Type: "float32", LocalType: "float32", Doc: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
})

var _ = gti.AddFunc(&gti.Func{
	Name:       "goki.dev/gti/gtigen/testdata.TypeOmittedArgs4",
	Doc:        "",
	Directives: gti.Directives{},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"x", &gti.Field{Name: "x", Type: "string", LocalType: "string", Doc: ""}},
		{"y", &gti.Field{Name: "y", Type: "string", LocalType: "string", Doc: ""}},
		{"z", &gti.Field{Name: "z", Type: "string", LocalType: "string", Doc: ""}},
		{"w", &gti.Field{Name: "w", Type: "bool", LocalType: "bool", Doc: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
})
