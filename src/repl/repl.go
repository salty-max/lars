package repl

import (
	"bufio"
	"fmt"
	"io"
	"os/user"

	"github.com/salty-max/lars/src/lexer"
	"github.com/salty-max/lars/src/log"
	"github.com/salty-max/lars/src/parser"
)

const PROMPT = ">> "

// Start starts the REPL.
func Start(in io.Reader, out io.Writer, user *user.User) {
	logger := log.NewLogger(out)
	logger.Info("Lars REPL v0.1.0")
	logger.Info(fmt.Sprintf("Hello %s!", user.Username))

	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Fprintf(out, "%+v\n", tok.Debug())
		// }

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []parser.ParserError) {
	io.WriteString(out, log.Colorize(log.RED, "Woops! Something went awry!\n"))

	io.WriteString(out, log.Colorize(log.RED, fmt.Sprintf("Parser has %d error(s)\n", len(errors))))
	for _, err := range errors {
		io.WriteString(out, log.Colorize(log.RED, fmt.Sprintf("\t(%d:%d) -> %s\n", err.Line, err.Col, err.Msg)))
	}
}
