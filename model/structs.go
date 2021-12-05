package model

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

func Do(db, collection, path string) {

	log.Println(path)

	var all = read(path)

	var fSet = token.NewFileSet()

	f, err := parser.ParseFile(fSet, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var res [][]byte

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

		var n, ok = node.(*ast.StructType)
		if !ok {
			return true
		}

		// log.Println(f.Scope.Objects)

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

		res = append(res, create(packageName, structName, db, collection, mongoIDName, mongoIDType))

		return true
	})

	addImportWithoutAnyImport(f)

	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fSet, f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(res); i++ {
		buffer.Write(res[i])
	}

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

func addImportWithoutAnyImport(file *ast.File) {
	var gList = []ast.Decl{&ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote("github.com/lemoyxk/longo/longo"),
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote("github.com/lemoyxk/longo/model"),
				},
			},
		},
	}}

	file.Decls = append(gList, file.Decls...)
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
