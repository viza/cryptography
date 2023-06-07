package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/viza/cryptography/sha1/sha1impl"
)

/* See https://www.rfc-editor.org/rfc/rfc3174.html */

func main() {

	printHelp()

	args := os.Args
	if len(args) > 1 && args[1] != "" {
		// standard lib, test data from file
		SHA1FromFile(args[1])
	} else {
		str := "This is a sha1 test string"
		//standard lib, test data as string
		sha1 := SHA1FromStr(str)
		fmt.Printf("%s: %s\n", str, sha1)
	}

	byteArr := []byte("This is a sha1 test string")
	s := sha1impl.New()
	s.Write(byteArr)
	sumArr := s.SHA1Sum(byteArr)
	fmt.Printf("This is a sha1 impl test string: %x\n", sumArr)

}

func SHA1FromStr(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func SHA1FromFile(filePath string) (string, error) {

	var sha1Str string

	file, err := os.Open(filePath)
	if err != nil {
		return sha1Str, err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return sha1Str, err
	}
	//Get the 20 bytes hash
	hashInBytes := hash.Sum(nil)[:20]

	//Convert the bytes to a string
	sha1Str = hex.EncodeToString(hashInBytes)

	fmt.Printf("SHA1 from file: %s\n", sha1Str)
	return sha1Str, nil
}

func printHelp() {
	fmt.Println("Making sha1 hash...")
	fmt.Printf("\n")
	fmt.Println("Make sha1 from test string - use no args in command line")
	fmt.Println("Make sha1 from test file - use filename as argument in command line")
	fmt.Printf("\n")
}
