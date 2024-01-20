package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"tensei/lexer"
	"tensei/token"
)

const PROMPT = ">> "

const HELP = `
  Commands:
  exit/quit: quit the program
  --file <filePath>: read the file and run source code from file 
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

program:
	for {
		var l *lexer.Lexer
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		switch line {
		case "exit":
			fallthrough
		case "quit":
			break program
		case "help":
			println(HELP)
			continue program
		default:
			if len(line) > 7 && line[0:7] == "--file " {
				bytes, err := os.ReadFile(line[7:])
				if err != nil {
					fmt.Println(err)
				}
				l = lexer.NewLexer(string(bytes))
			} else {
				l = lexer.NewLexer(line)
			}
		}

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
