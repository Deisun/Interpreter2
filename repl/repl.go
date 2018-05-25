// Read Evaluate Print Loop (console)
package repl

import (
	"io"
	"bufio"
	"fmt"
	"Interpreter2/lexer"
	"Interpreter2/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner :=  bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
