// +build ignore

package main

import (
	"os"
	"text/template"
)

// Borrowed from the core reflect package
// Ignoring: Array, Chan, Func, Map, Ptr, Slice, Struct
var types = map[string]string{
	"Bool":          "bool",
	"Int":           "int",
	"Int8":          "int8",
	"Int16":         "int16",
	"Int32":         "int32",
	"Int64":         "int64",
	"Uint":          "uint",
	"Uint8":         "uint8",
	"Uint16":        "uint16",
	"Uint32":        "uint32",
	"Uint64":        "uint64",
	"Uintptr":       "uintptr",
	"Float32":       "float32",
	"Float64":       "float64",
	"Complex64":     "complex64",
	"Complex128":    "complex128",
	"Interface":     "interface{}",
	"String":        "string",
	"UnsafePointer": "unsafe.Pointer",
}

var pattern = `// This file has been generated with: go run gen.go; DO NOT EDIT!

package ternary

import "unsafe"
{{range $name, $type := .}}
// {{$name}} provides the ternary operator for {{$type}} types
func {{$name}}(ok bool, truthy, falsey {{$type}}) {{$type}} {
	if ok {
		return truthy
	}
	return falsey
}
{{end}}`

func do(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	tpl := template.Must(template.New("").Parse(pattern))
	file, err := os.Create("ternary.go")
	do(err)
	do(tpl.Execute(file, types))
	do(file.Close())
}
