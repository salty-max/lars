package repl

import (
	"bufio"
	"fmt"
	"io"
	"os/user"

	"github.com/salty-max/lars/src/lexer"
	"github.com/salty-max/lars/src/log"
	"github.com/salty-max/lars/src/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok.Debug())
		}
	}
}
