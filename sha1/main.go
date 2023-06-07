package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {

	str := "This is a sha1 test string"

	SHA1FromStr(str)

}

func SHA1FromStr(s string) {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	fmt.Printf("%s: %s\n", s, sha1_hash)
}
