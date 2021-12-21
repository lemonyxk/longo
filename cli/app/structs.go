package app

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type createInfo struct {
	packageName string
	structName  string
	db          string
	collection  string
	mongoIDName string
	mongoIDType string
}

func Do(db, collection, path string) {

	log.Println(path)

	var all = read(path)

	var fSet = token.NewFileSet()

	f, err := parser.ParseFile(fSet, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var ci *createInfo

	ast.Inspect(f, func(node ast.Node) bool {

		// if fun, ok := node.(*ast.FuncDecl); ok {
		// 	// 4. Validate that method is exported and has a receiver
		// 	if fun.Name.IsExported() && fun.Recv != nil && len(fun.Recv.List) == 1 {
		// 		// 5. Check that the receiver is actually the struct we want
		// 		if r, rok := fun.Recv.List[0].Type.(*ast.StarExpr); rok {
		// 			// we found it!
		// 			log.Println(r.X.(*ast.Ident).Name)
		// 		}
		// 	}
		// }

		// switch node.(type) {
		// case *ast.File:
		// 	file := node.(*ast.File)
		// 	if len(file.Imports) == 0 {
		// 		addImportWithoutAnyImport(file, "github.com/lemoyxk/longo/longo")
		// 		addImportWithoutAnyImport(file, "github.com/lemoyxk/longo/model")
		// 	}
		// case *ast.GenDecl:
		// 	dec := node.(*ast.GenDecl)
		// 	if dec.Tok == token.IMPORT {
		// 		addImport(dec, "github.com/lemoyxk/longo/longo")
		// 		addImport(dec, "github.com/lemoyxk/longo/model")
		// 	}
		// }

		var n, ok = node.(*ast.StructType)
		if !ok {
			return true
		}

		var packageName = f.Name.Name
		if db == "" {
			db = packageName
		}

		var structName = getStructName(all, int(n.Pos()))
		if collection == "" {
			collection = structName
		}

		if !canExport(structName) {
			return true
		}

		var mongoIDName = ""
		var mongoIDType = ""

		for i := 0; i < len(n.Fields.List); i++ {

			if n.Fields.List[i].Tag == nil {
				continue
			}

			if isMongoID(n.Fields.List[i].Tag.Value) {
				mongoIDName = n.Fields.List[i].Names[0].Name
				mongoIDType = fmt.Sprintf("%s", n.Fields.List[i].Type)
			}
		}

		if mongoIDName == "" || mongoIDType == "" {
			return true
		}

		ci = &createInfo{packageName, structName, db, collection, mongoIDName, mongoIDType}

		return false
	})

	if ci == nil {
		return
	}

	ast.FilterFile(f, func(s string) bool {
		if s != ci.structName {
			return false
		}
		return true
	})

	addImportWithoutAnyImport(f, "github.com/lemoyxk/longo")
	// addImportWithoutAnyImport(f, "github.com/lemoyxk/longo/model")

	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fSet, f)
	if err != nil {
		panic(err)
	}

	var res = create(*ci)

	buffer.Write(res)

	write(path, buffer)

	goFmt(path)
}

func getStructName(all []byte, pos int) string {
	var start = -1
	var end = -1
	var repeat = 0
	for i := pos; i >= 1; i-- {
		if all[i] == ' ' && repeat == 0 {
			end = i
			repeat++
			continue
		}

		if all[i] == ' ' && repeat == 1 {
			start = i + 1
			repeat++
			break
		}
	}

	return string(all[start:end])
}

func canExport(name string) bool {
	return isBigWord(name[0])
}

func isBigWord(char byte) bool {
	return char >= 65 && char <= 90
}

func isMongoID(v string) bool {
	return strings.Index(v, `bson:"_id"`) != -1
}

func addImportWithoutAnyImport(file *ast.File, packageName string) {
	var gList = []ast.Decl{&ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote(packageName),
				},
			},
		},
	}}

	file.Decls = append(gList, file.Decls...)
}

func addImport(dec *ast.GenDecl, importName string) {
	hasImport := false
	for _, value := range dec.Specs {
		importSpec := value.(*ast.ImportSpec)
		if importSpec.Path.Value == strconv.Quote(importName) {
			hasImport = true
		}
	}
	if hasImport {
		return
	}
	if !hasImport {
		dec.Specs = append(dec.Specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(importName),
			},
		})
	}
}

func write(path string, buf *bytes.Buffer) {

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	defer func() { _ = f.Close() }()

	_, err = f.WriteString(buf.String())
	if err != nil {
		panic(err)
	}

}

func read(path string) []byte {

	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer func() { _ = f.Close() }()

	all, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return all
}
