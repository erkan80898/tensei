package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"tensei/lexer"
	"tensei/parser"
	"tensei/token"
	"tensei/util"

	"github.com/inancgumus/screen"
)

const PROMPT = ">> "

const HELP = `
  Commands:
  exit/quit: quit the program
  clear: clears screen
  --set <mode>: set mode of repl l-lexer, p-parser (DEFAULT: parser)
  --file <filePath>: read the file and run source code from file 
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	mode := "p"
program:
	for {
		var l *lexer.Lexer
		var p *parser.Parser
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if len(line) > 7 && line[0:7] == "--file " {
			bytes, err := os.ReadFile(line[7:])
			if err != nil {
				fmt.Println(err)
			}
			line = string(bytes)
		} else if len(line) > 6 && line[0:6] == "--set " {
			mode = string(line[6])
			text := mode
			if mode == "l" {
				text = "Lexer"
			} else {
				text = "Parser"
			}

			fmt.Printf("Mode changed to: %s", text+util.NEW_LINE)
			continue program
		}

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
			l = lexer.New(line)
		}

		switch mode {
		case "l":
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}
		case "p":
			p = parser.New(l)
			program := p.ParseProgram()
			println(program.Statements[0].Display())
		}
	}
}
