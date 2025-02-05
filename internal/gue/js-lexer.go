package gue

import (
	"fmt"
	"log"
	"unicode"
)

// Token Types
const (
	NotSpecified   = "NOT_SPECIFIED"   //
	Eof            = "EOF"             // \x00
	Identifier     = "IDENTIFIER"      //
	Number         = "NUMBER"          //
	Quote          = "QUOTE"           // " ' `
	Boolean        = "BOOLEAN"         // true false
	Whitespace     = "WHITESPACE"      //
	LineComment    = "LINE_COMMENT"    // //
	BlockComment   = "BLOCK_COMMENT"   // /*
	VarDeclarer    = "VAR_DECLARER"    // var let const
	FuncDeclarer   = "FUNC_DECLARER"   // function
	Setter         = "SETTER"          // =
	Equality       = "EQUALITY"        // == ===
	Colon          = "COLON"           // :
	Semicolon      = "SEMICOLON"       // ;
	Comma          = "COMMA"           // ,
	Dot            = "DOT"             // .
	Plus           = "PLUS"            // +
	Minus          = "MINUS"           // -
	Divide         = "DIVIDE"          // /
	MultipLy       = "MULTIPLY"        // *
	OpeningParen   = "OPENING_PAREN"   // (
	ClosingParen   = "CLOSING_PAREN"   // )
	OpeningBrace   = "OPENING_BRACE"   // {
	ClosingBrace   = "CLOSING_BRACE"   // }
	OpeningBracket = "OPENING_BRACKET" // [
	ClosingBracket = "CLOSING_BRACKET" // ]
)

type Token struct {
	Type  string
	Value string
}

type Tokenizer struct {
	position     int
	readPosition int
	char         rune
	input        string
}

func NewTokenizer(input string) Tokenizer {
	var t = Tokenizer{
		position:     0,
		readPosition: 0,
		input:        input,
	}
	t.readChar()
	return t
}

func (t *Tokenizer) GetNextToken() Token {

	if unicode.IsSpace(t.char) {
		return Token{Whitespace, t.readWhitespace()}
	}

	if isStartOfWord((t.char)) {
		identifier := t.readIdentifier()

		token, exists := KeywordMap[identifier]
		if exists {
			// t.readChar()
			return token
		}

		return Token{Identifier, identifier}
	}

	if isDigit((t.char)) {
		return Token{Number, t.readNumber()}
	}

	if isQuote(t.char) {
		return Token{Quote, t.readQuote()}
	}

	var token Token
	switch t.char {
	case '[':
		token = Token{OpeningBracket, string(t.char)}
	case ']':
		token = Token{ClosingBracket, string(t.char)}
	case '{':
		token = Token{OpeningBrace, string(t.char)}
	case '}':
		token = Token{ClosingBrace, string(t.char)}
	case '(':
		token = Token{OpeningParen, string(t.char)}
	case ')':
		token = Token{ClosingParen, string(t.char)}
	case ',':
		token = Token{Comma, string(t.char)}
	case ';':
		token = Token{Semicolon, string(t.char)}
	case ':':
		token = Token{Colon, string(t.char)}
	case '*':
		token = Token{MultipLy, string(t.char)}
	case '=':
		if t.nextChar() == '=' {
			token = Token{Equality, t.readEquality()}
		} else {
			token = Token{Setter, string(t.char)}
		}
	case '/':
		next := t.nextChar()
		if next == '/' {
			token = Token{LineComment, t.readLineComment()}
		} else if next == '*' {
			token = Token{BlockComment, t.readBlockComment()}
		} else {
			token = Token{Divide, string(t.char)}
		}
	case '.':
		if isDigit(t.nextChar()) {
			token = Token{Number, t.readNumber()}
		} else {
			token = Token{Dot, string(t.char)}
		}
	case '+':
		if isDigit(t.nextChar()) {
			token = Token{Number, t.readNumber()}
		} else {
			token = Token{Plus, string(t.char)}
		}
	case '-':
		if isDigit(t.nextChar()) {
			token = Token{Number, t.readNumber()}
		} else {
			token = Token{Minus, string(t.char)}
		}
	case '\x00':
		token = Token{Eof, "eof"}
		return token
	default:
		token = Token{NotSpecified, string(t.char)}
	}

	t.readChar()
	return token
}

func (t *Tokenizer) nextChar() rune {
	return rune(t.input[t.readPosition])
}

var KeywordMap = map[string]Token{
	"true":     {Boolean, "true"},
	"false":    {Boolean, "false"},
	"var":      {VarDeclarer, "var"},
	"let":      {VarDeclarer, "let"},
	"const":    {VarDeclarer, "const"},
	"function": {FuncDeclarer, "function"},
}

func isStartOfWord(r rune) bool {
	return r >= 'a' && r <= 'z' ||
		r >= 'A' && r <= 'Z' ||
		r == '_' || r == '$'
}

func isIdentifierChar(r rune) bool {
	return r >= 'a' && r <= 'z' ||
		r >= 'A' && r <= 'Z' ||
		r >= '0' && r <= '9' ||
		r == '_' || r == '$'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isQuote(r rune) bool {
	return r == '"' || r == '\'' || r == '`'
}

func (t *Tokenizer) readChar() {
	if t.readPosition >= len(t.input) {
		t.char = '\x00'
		t.position = t.readPosition
		return
	}

	t.char = rune(t.input[t.readPosition])
	t.position = t.readPosition
	t.readPosition++
}

func (t *Tokenizer) readWhitespace() string {
	start := t.position
	t.readChar()

	for unicode.IsSpace(t.char) {
		t.readChar()
	}

	return t.input[start:t.position]
}

func (t *Tokenizer) readIdentifier() string {
	start := t.position
	t.readChar()

	for isIdentifierChar(t.char) {
		t.readChar()
	}

	return t.input[start:t.position]
}

func (t *Tokenizer) readNumber() string {
	start := t.position
	t.readChar()

	for !unicode.IsSpace(t.char) {
		if t.char == '\x00' {
			break
		}
		t.readChar()
	}

	return t.input[start:t.position]
}

func (t *Tokenizer) readQuote() string {
	start := t.position
	quoteMark := t.char
	fmt.Println("quoteMark:", string(quoteMark))

	t.readChar()

	num := 1
	for t.char != quoteMark {
		fmt.Println("LOOP:", num)
		fmt.Println("t.char:", string(t.char))
		t.readChar()
		num++
		if num > 3 {
			log.Fatal("force quit")
		}
	}
	t.readChar()
	return t.input[start:t.position]
}

func (t *Tokenizer) readLineComment() string {
	start := t.position
	t.readChar()

	for t.char != '\n' && t.char != '\r' {
		t.readChar()
	}

	return t.input[start:t.position]
}

func (t *Tokenizer) readBlockComment() string {
	start := t.position
	t.readChar()

	for string(t.char)+string(t.input[t.readPosition]) != "*/" {
		t.readChar()
	}

	return t.input[start:t.position]
}

func (t *Tokenizer) readEquality() string {
	start := t.position
	t.readChar()

	for t.char == '=' {
		t.readChar()
	}

	return t.input[start:t.position]
}
