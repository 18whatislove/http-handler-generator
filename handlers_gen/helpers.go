package main

import (
	"fmt"
	"strconv"
	"strings"
)

func format2Slice(enum, sliceType string) string {
	// "1|2|3", "int" -> "[]int{1, 2, 3}"
	const sep string = "|"
	var tpl string
	b := &strings.Builder{}
	b.Grow(len(enum) + strings.Count(enum, sep))

	switch sliceType {
	case "int":
		b.WriteString(strings.ReplaceAll(enum, sep, ", "))
		tpl = "[]%s{%s}"
	case "string":
		b.WriteString(strings.ReplaceAll(enum, sep, "\", \""))
		tpl = "[]%s{\"%s\"}"
	}
	return fmt.Sprintf(tpl, sliceType, b.String())
}

func tagValue2Struct(tagValue string, f *Field) {
	var name, val string
	// tags = [paramname, default, min, max, required, enum]

	for _, option := range strings.Split(tagValue, ",") {
		parsedOption := strings.Split(option, "=")

		name = parsedOption[0]
		if name == "required" {
			f.Required = true
			continue
		}
		val = parsedOption[1]

		switch name {
		case "paramname":
			f.ParamName = val
		case "default":
			f.Default.IsSet = true
			if f.Type == "string" {
				f.Default.Value = strconv.Quote(val)
			} else if f.Type == "int" {
				f.Default.Value = val
			}
		case "max":
			f.Max.IsSet = true
			f.Max.Value = val
		case "min":
			f.Min.IsSet = true
			f.Min.Value = val
		case "enum":
			f.Enum.IsSet = true
			f.Enum.Value = format2Slice(val, f.Type)
			// f.Enum.Value = val
		}
	}

}

// func getImports(f *ast.File) []string {
// 	imports := make([]string, 0, 10)
// 	for _, decl := range f.Decls {
// 		decl, ok := decl.(*ast.GenDecl)
// 		if !ok {
// 			continue
// 		}
// 		if decl.Tok == token.IMPORT {
// 			for _, spec := range decl.Specs {
// 				spec := spec.(*ast.ImportSpec)
// 				// spike
// 				if spec.Path.Value == "context" || spec.Path.Value == "sync" {
// 					continue
// 				}
// 				imports = append(imports, spec.Path.Value)
// 			}
// 		}
// 	}
// 	return imports
// }
