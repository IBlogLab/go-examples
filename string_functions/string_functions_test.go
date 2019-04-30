package string_functions_test

import (
	"fmt"
	"strings"
	"testing"
)

var p = fmt.Println

func TestStringFunctions(t *testing.T) {
	p("Contains: ", strings.Contains("test", "es"))
	p("Count: ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasPrefix: ", strings.HasPrefix("test", "st"))
	p("Index: ", strings.Index("test", "e"))
	p("Join: ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", strings.Repeat("a", 5))
	p("Replace: ", strings.Replace("foo", "o", "0", -1))
	p("Replace: ", strings.Replace("foo", "o", "0", 1))
	p("Split: ", strings.Split("a-b-c", "-"))
	p("ToLower: ", strings.ToLower("TEST"))
	p("ToUpper: ", strings.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char: ", "hello"[1])
}
