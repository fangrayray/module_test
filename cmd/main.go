package main

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/fangrayray/module_test/parser"
)

func main() {
	// input := "[file:hashes.'SHA-256' = '56cca56e39431187a2bd95e53eece8f11d3cbe2ea7ee692fa891875f40f233f5']"
	// input := "[file:hashes.'sha1' = '0011e13490dac077425863667eeaa97bb40449cf2d5752dc0b010c82200a0720']"
	input := "[file:hashes.'md5' = 'd41d8cd98f00b204e9800998ecf8427e']"
	// input := "[(directory:path MATCHES '^C:\\\\Windows\\\\w+$') OR (file:parent_directory_ref.path = 'C:\\\\Windows\\\\System32')]"
	// GG input := "[[directory:path MATCHES '^C:\\\\Windows\\\\w+$'] OR [file:parent_directory_ref.path = 'C:\\\\Windows\\\\System32']]"
	// input := "[file:name = 'abc.exe'] AND [file:parent_directory_ref.path = 'C:\\\\Windows\\\\System32']"
	// input := " [file:name = 'abc.exe'] AND [ file:name = 'cc']"
	// input := " [ file:name = 'abcd.exe' AND file:name = 'abc.exe'] "
	// input := " [file:name = 'abc.exe']"

	printYara(input)
}

func printYara(s string) {

	// Setup the input
	is := antlr.NewInputStream(s)

	// Create the Lexer
	lexer := parser.NewSTIXPatternLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTIXPatternParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Pattern()

	// Finally parse the expression by walking the tree
	var listener PocListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
	root := listener.expressionNode

	forEachNode(root)

}

func forEachNode(n *ExpressionNode) {
	if n == nil {
		return
	}
	if n.Type == "JOIN" {
		ok := isFilePathCase(n)
		if !ok {
			forEachNode(n.FirstNode)
			forEachNode(n.SlibingNode)
		}

	} else {
		ok := handleFileChecksum(n)
		if !ok {
			fmt.Println(
				fmt.Sprintf(
					"%s - %s:%s %s %s",
					n.Type,
					n.Data.Object,
					n.Data.NestedProperty,
					n.Data.Operator,
					n.Data.Value,
				))
		}

	}
}

func handleFileChecksum(n *ExpressionNode) bool {
	if n.Type == NODETYPE &&
		n.Data.Object == "file" {
		data := n.Data
		prop := strings.ToLower(data.NestedProperty)
		if prop == "hashes.'sha-256'" || prop == "hashes.'sha1'" || prop == "hashes.'md5'" {
			val := strings.ReplaceAll(data.Value, "'", "")
			var buffer bytes.Buffer
			buffer.WriteString("import \"hash\"\n")
			buffer.WriteString("\n")
			buffer.WriteString("rule poc {\n")
			buffer.WriteString("  condition:\n")
			switch prop {
			case "hashes.'sha-256'":
				buffer.WriteString("     hash.sha256(0, filesize) == \"" + val + "\"\n")
			case "hashes.'sha1'":
				buffer.WriteString("     hash.sha1(0, filesize) == \"" + val + "\"\n")
			case "hashes.'md5'":
				buffer.WriteString("     hash.md5(0, filesize) == \"" + val + "\"\n")

			}

			buffer.WriteString("}\n")

			fmt.Println(buffer.String())
			return true
		}
	}
	return false
}

func isFilePathCase(n *ExpressionNode) bool {
	// if true, merge 2 node as 1
	if n.Type == JOINTYPE &&
		n.FirstNode != nil && n.FirstNode.Type == NODETYPE &&
		n.SlibingNode != nil && n.SlibingNode.Type == NODETYPE {
		fData := n.FirstNode.Data
		sData := n.SlibingNode.Data
		ok := handleFileNameCase(fData, sData)
		if !ok {
			ok = handleFileNameCase(sData, fData)
		}
		return ok

	}
	return false
}

func handleFileNameCase(pre *AtomExp, pos *AtomExp) bool {
	if pre.Object == "file" && pre.NestedProperty == "name" {
		if pos.Object == "file" && pos.NestedProperty == "parent_directory_ref.path" {
			// Print yara rule for file name
			YaraRuleFileName(pre, pos)
			return true
		}
	}
	return false
}

// YaraRuleFileName -
func YaraRuleFileName(fileExp *AtomExp, pathExp *AtomExp) {
	fullpath := filepath.Join(strings.ReplaceAll(pathExp.Value, "'", ""),
		strings.ReplaceAll(fileExp.Value, "'", ""))
	var buffer bytes.Buffer
	buffer.WriteString("rule poc {\n")
	buffer.WriteString("  string:\n")
	buffer.WriteString("    $path = \"" + fullpath + "\"\n")
	buffer.WriteString("  condition:\n")
	buffer.WriteString("     $path\n")
	buffer.WriteString("}\n")

	fmt.Println(buffer.String())
}

func printTokens(s string) {
	// Setup the input
	is := antlr.NewInputStream(s)
	// Create the Lexer
	lexer := parser.NewSTIXPatternLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
