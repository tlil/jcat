package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/titanous/json5"
)

func main() {
	input, err := readInput(os.Args[1:], os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	output, ok := formatJSON5(input)
	if !ok {
		fmt.Println(string(input))
		return
	}

	fmt.Println(string(output))
}

func readInput(args []string, stdin io.Reader) ([]byte, error) {
	if len(args) > 0 {
		return os.ReadFile(args[0])
	}

	return io.ReadAll(stdin)
}

func formatJSON5(input []byte) ([]byte, bool) {
	var value any
	if err := json5.Unmarshal(input, &value); err != nil {
		return nil, false
	}

	output, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return nil, false
	}

	return output, true
}
