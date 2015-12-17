// Reads from stdin and converts all characters to lower case.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	for r.Scan() {
		fmt.Println(strings.ToLower(r.Text()))
	}
}
