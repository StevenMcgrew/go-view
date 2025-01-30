package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// func ProcessFile(filename string) {
// 	// Access the file
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	defer file.Close()

// 	const (
// 		none = iota
// 		outsideTag
// 		tagStart
// 		closingTag
// 		attributes
// 		possibleVarStart
// 		variable
// 		possibleVarEnd
// 		script
// 		style
// 		exclamation
// 		comment
// 		doctype
// 	)

// 	doc := ""
// 	state := 1
// 	tag := ""
// 	stateBeforeVar := 0
// 	_var := ""

// 	// Scan the file
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		for _, r := range line {

// 			// This has to be checked for every state, so just check it here first
// 			if r == '{' {
// 				stateBeforeVar = state
// 				state = possibleVarStart
// 				continue
// 			}

// 			if state == possibleVarStart {

// 				// checks
// 				if r == '{' {
// 					state = variable
// 					continue
// 				}

// 				// process
// 				stateBeforeVar = 1
// 				doc += "{" + string(r)
// 				continue
// 			}

// 			if state == variable {

// 				// checks
// 				if r == '}' {
// 					state = possibleVarEnd
// 					continue
// 				}

// 				// process
// 				_var += string(r)
// 				continue
// 			}

// 			if state == possibleVarEnd {

// 				// checks
// 				if r != '}' {
// 					doc += "}" + string(r)
// 					state = variable
// 					continue
// 				}

// 				// process
// 				// TODO: eval variable
// 				// doc += evalExpr(expression)
// 				state = stateBeforeVar
// 				continue
// 			}

// 			if state == outsideTag {

// 				// checks
// 				if r == '<' {
// 					state = tagStart
// 					continue
// 				}
// 				if r == '{' {
// 					state = possibleVarStart
// 					continue
// 				}

// 				// process
// 				doc += string(r)
// 				continue
// 			}

// 			if state == tagStart {
// 				if r == ' ' {
// 					state = attributes
// 					continue
// 				}
// 				if r == '>' {
// 					state = outsideTag
// 					continue
// 				}
// 				if r == '!' {
// 					state = exclamation
// 					continue
// 				}
// 				if r == '/' {
// 					state = closingTag
// 					continue
// 				}

// 				// process
// 				tag += string(r)
// 				continue
// 			}
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "reading input:", err)
// 	}
// }

// func ProcessJS(js string) string {

// 	const (
// 		none = iota
// 	)

// 	return ""
// }

// func ProcessHtml(html string) {
// 	// Get working directory
// 	wdir, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Start by reading root file (index.html)
// 	bytes, err := os.ReadFile(wdir + "/frontend/index.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	const (
// 		none = iota
// 		outsideTag
// 		tagStart
// 		closingTag
// 		attributes
// 		possibleExprStart
// 		expr
// 		possibleExprEnd
// 		script
// 		style
// 		exclamation
// 		comment
// 		doctype
// 	)

// 	doc := ""
// 	state := 1
// 	tag := ""
// 	stateBeforeExpr := 0
// 	expression := ""

// 	for i := 0; i < len(bytes); i++ {
// 		b := rune(bytes[i])

// 		// This has to be checked for every state, so just check it here first
// 		if b == '{' {
// 			stateBeforeExpr = state
// 			state = possibleExprStart
// 			continue
// 		}

// 		if state == possibleExprStart {

// 			// checks
// 			if b == '{' {
// 				state = expr
// 				continue
// 			}

// 			// process
// 			stateBeforeExpr = 0
// 			doc += "{" + string(b)
// 			continue
// 		}

// 		if state == expr {

// 			// checks
// 			if b == '}' {
// 				state = possibleExprEnd
// 				continue
// 			}

// 			// process
// 			expression += string(b)
// 			continue
// 		}

// 		if state == possibleExprEnd {

// 			// checks
// 			if b != '}' {
// 				doc += "}" + string(b)
// 				state = expr
// 				continue
// 			}

// 			// process
// 			// TODO: eval variable
// 			// doc += evalExpr(expression)
// 			state = stateBeforeExpr
// 			continue
// 		}

// 		if state == outsideTag {

// 			// checks
// 			if b == '<' {
// 				state = tagStart
// 				continue
// 			}
// 			if b == '{' {
// 				state = possibleExprStart
// 				continue
// 			}

// 			// process
// 			doc += string(b)
// 			continue
// 		}

// 		if state == tagStart {
// 			if b == ' ' {
// 				state = attributes
// 				continue
// 			}
// 			if b == '>' {
// 				state = outsideTag
// 				continue
// 			}
// 			if b == '!' {
// 				state = exclamation
// 				continue
// 			}
// 			if b == '/' {
// 				state = closingTag
// 				continue
// 			}

// 			// process
// 			tag += string(b)
// 			continue
// 		}
// 	}

// 	// This has potential
// 	// f := func(c rune) bool {
// 	// 	if c == '{' || c == '}' {
// 	// 		return true
// 	// 	}
// 	// 	return false
// 	// }
// 	// fmt.Printf("Fields are: %q", strings.FieldsFunc(" AAA { test } BBB {{test2}}", f))
// 	// Output: Fields are: [" AAA " " test " " BBB " "test2"]

// 	// Could use Index to slice the string at index of {}'s

// 	// This could work, not fully fleshed out yet.
// 	// str := "aaa{{test1}}bbb{{test2}}ccc{{test3}}ddd"
// 	// slice1 := strings.Split(str, "{{")
// 	// fmt.Printf("%q\n", slice1)

// 	// var b strings.Builder

// 	// b.WriteString(slice1[0])

// 	// for i := 1; i < len(slice1); i++ {
// 	// 	section := slice1[i]
// 	// 	slice2 := strings.Split(section, "}}")
// 	// 	fmt.Printf("%q\n", slice2)
// 	// }
// 	// Output: ["aaa" "test1}}bbb" "test2}}ccc" "test3}}ddd"]
// 	// ["test1" "bbb"]
// 	// ["test2" "ccc"]
// 	// ["test3" "ddd"]
// }

// func ProcessCss(css string) {
// 	fmt.Println("Start ProcessCss")
// 	r, err := os.Open("file.scss")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	comp, err := libsass.New(os.Stdout, r)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := comp.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

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

func getScript(doc *html.Node) {
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode &&
			node.DataAtom == atom.Script {
			html.Render(os.Stdout, node.FirstChild)
			fmt.Println()
			return
		}
	}
}

func processTemplate(doc *html.Node, compId string) {
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode &&
			node.DataAtom == atom.Template {
			appendClassNames(node, compId)
			for n := range node.Descendants() {
				appendClassNames(n, compId)
			}
			html.Render(os.Stdout, node)
			fmt.Println()
			break
		}
	}
}

func appendClassNames(node *html.Node, compId string) {
	for i, a := range node.Attr {
		if a.Key == "class" {
			classNames := strings.Fields(a.Val)
			for i, cName := range classNames {
				classNames[i] = cName + "-" + compId
			}
			classNames = append(classNames, compId)
			node.Attr[i].Val = strings.Join(classNames, " ")
			return
		}
	}

	// "class" attribute not found, therefore add the compId scope class
	node.Attr = append(node.Attr, html.Attribute{Key: "class", Val: compId})
}

func main() {
	// Access the file
	file, err := os.Open("ModalBox.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal("Error parsing HTML.")
	}

	compId := generateId()

	// getScript(doc)

	processTemplate(doc, compId)

	// // Create goquery.Document
	// doc, err := goquery.NewDocumentFromReader(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var sel = &goquery.Selection{}

	// // Get script contents
	// sel = doc.Find("script").First()
	// fmt.Println(sel.Html())

	// // Get template contents
	// sel = doc.Find("template").First()
	// fmt.Println(sel.Html())

	// // Get style contents
	// sel = doc.Find("style").First()
	// fmt.Println(sel.Html())
}
