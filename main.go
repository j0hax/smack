package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/j0hax/smack/interpreter"
	"github.com/j0hax/smack/lexer"
	"github.com/jwalton/gchalk"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result := interpreter.Process(scanner.Text())

		out := result.String()
		switch result.Type {
		case lexer.Number:
			fmt.Println(gchalk.Green(out))
		default:
			fmt.Println(out)
		}
	}
}
