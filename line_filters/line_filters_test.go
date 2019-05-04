package line_filters_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLineFilters(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdout)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
