package gue

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/tdewolff/parse/css"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateId() string {
	b := make([]byte, 6)
	b[0] = alphanumeric[rand.Intn(len(alphanumeric)-10)]
	for i := 1; i < len(b); i++ {
		b[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}
	return string(b)
}

func getScript(doc *html.Node) string {
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode &&
			node.DataAtom == atom.Script {
			var buf bytes.Buffer
			html.Render(&buf, node.FirstChild)
			return buf.String()
		}
	}
	return ""
}

func getStyle(doc *html.Node) string {
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode &&
			node.DataAtom == atom.Style {
			var buf bytes.Buffer
			html.Render(&buf, node.FirstChild)
			return buf.String()
		}
	}
	return ""
}

func processTemplate(doc *html.Node, scopeId string) {
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode &&
			node.DataAtom == atom.Template {
			appendClassNames(node, scopeId)
			for n := range node.Descendants() {
				appendClassNames(n, scopeId)
			}
			html.Render(os.Stdout, node)
			fmt.Println()
			break
		}
	}
}

func appendClassNames(node *html.Node, scopeId string) {
	foundClassAttr := false
	for i, a := range node.Attr {
		if a.Key == "class" {
			foundClassAttr = true
			classNames := strings.Fields(a.Val)
			for i, cName := range classNames {
				classNames[i] = cName + "-" + scopeId
			}
			classNames = append(classNames, scopeId)
			node.Attr[i].Val = strings.Join(classNames, " ")
		} else if a.Key == "id" {
			node.Attr[i].Val = a.Val + "-" + scopeId
		}
	}
	if !foundClassAttr {
		node.Attr = append(node.Attr, html.Attribute{Key: "class", Val: scopeId})
	}
}

func processCss(doc *html.Node, scopeId string, compName string) {
	styles := getStyle(doc)

	l := css.NewLexer(strings.NewReader(styles))
	var output string
	inBraces := false
	isClass := false
	for {
		tt, bytes := l.Next()
		data := string(bytes)
		if tt == css.ErrorToken {
			break
		}

		switch tt {
		case css.LeftBraceToken:
			inBraces = true
			output += data
		case css.RightBraceToken:
			inBraces = false
			output += data
		case css.HashToken:
			output += data + "." + scopeId
		case css.DelimToken:
			if !inBraces && data == "." {
				isClass = true
			}
			output += data
		case css.IdentToken:
			if !inBraces {
				if isClass {
					output += data + "-" + scopeId
					isClass = false
				} else if data == "comp" {
					output += compName
				} else {
					output += data + "." + scopeId
				}
			} else {
				output += data
			}
		default:
			output += data
		}
	}
	fmt.Println(output)
}

func getCompName(filePath string) string {
	lastSlashIndex := strings.LastIndex(filePath, "/")

	if lastSlashIndex == -1 {
		lastSlashIndex = 0
	} else if lastSlashIndex < len(filePath)-1 {
		lastSlashIndex++
	}

	filePath = filePath[lastSlashIndex:]

	lastDotIndex := strings.LastIndex(filePath, ".")

	if lastDotIndex == -1 {
		return filePath
	}

	return filePath[0:lastDotIndex]
}

func SplitFile(path string) {
	// Access the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the file
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal("Error parsing HTML.")
	}

	// Variables that will be needed
	scopeId := generateId()
	compName := getCompName(file.Name())

	// Process the sections of the file
	processTemplate(doc, scopeId)
	processCss(doc, scopeId, compName)
	script := getScript(doc)
	fmt.Println(script)
}
