// Package docToHtml is read stdin and create Html file from data.
// Write it to new file.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// Struct for read package and import names from Comments.
type Names struct {
	PackageValue string // package name
	ImportValue  string // import name
}

// Main goroutine
func main() {
	strByte := ReadStdin()
	var packagesNames Names
	packagesNames.NamesScan(string(strByte))
	fileName := "doc.p_" + packagesNames.PackageValue + "_.i_" + packagesNames.ImportValue + "_.html"
	os.WriteFile(fileName, BuildHtmlPage(&packagesNames, strByte), 0744)
}

// Read data from Stdin and write it to strByte([]byte)
func ReadStdin() (strByte []byte) {
	strByte, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Can not scan data")
		os.Exit(1)
	}
	return
}

// Outside parser MD
func MdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

// Building HTML page
func BuildHtmlPage(p *Names, strByte []byte) []byte {
	fileBody := "<!doctype html><html lang='ru'><head><meta charset='utf-8'><title>" +
		"Documentation for package " + p.PackageValue + " imported as " +
		p.ImportValue + "</title></head><body><div align=\"left\" >" + string(MdToHTML(strByte)) +
		"</div></body></html>"

	return []byte(fileBody)
}

// Scan string for package and import names
func (n *Names) NamesScan(str string) {
	var tempImportValue, tempPackageValue, tempImport, tempPackage string
	var tempSlash string
	_, err := fmt.Sscan(str, &tempPackage, &tempPackageValue, &tempSlash, &tempImport, &tempImportValue)
	if err != nil {
		fmt.Println("Can not scan data")
		os.Exit(1)
	}
	if strings.EqualFold(tempImport, "import") && !(strings.EqualFold(tempImportValue, "")) {
		n.ImportValue = tempImportValue[1 : len(tempImportValue)-1]
	} else {
		n.ImportValue = "unknownImport"
	}
	if strings.EqualFold(tempPackage, "package") && !(strings.EqualFold(tempPackageValue, "")) {
		n.PackageValue = tempPackageValue
	} else {
		n.PackageValue = "unknownPackage"
	}
}
