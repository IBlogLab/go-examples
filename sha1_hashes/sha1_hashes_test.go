package sha1_hashes_test

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestSha1Hashes(t *testing.T) {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(bs)
	fmt.Printf("%x\n", bs)
}
