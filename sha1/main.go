package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {

	printHelp()

	args := os.Args

	if len(args) > 1 && args[1] != "" {

		SHA1FromFile(args[1])
	} else {
		str := "This is a sha1 test string"
		sha1 := SHA1FromStr(str)
		fmt.Printf("%s: %s\n", str, sha1)
	}
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
