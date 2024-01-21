package repl

import (
	"bufio"
	"fmt"
	"github.com/inancgumus/screen"
	"io"
	"os"
	"tensei/lexer"
	"tensei/token"
)

const PROMPT = ">> "

const HELP = `
  Commands:
  exit/quit: quit the program
  clear: clears screen
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
		case "clear":
			screen.Clear()
			screen.MoveTopLeft()
			continue program
		case "help":
			println(HELP)
			continue program
		default:
			if len(line) > 7 && line[0:7] == "--file " {
				bytes, err := os.ReadFile(line[7:])
				if err != nil {
					fmt.Println(err)
				}
				l = lexer.New(string(bytes))
			} else {
				l = lexer.New(line)
			}
		}

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
