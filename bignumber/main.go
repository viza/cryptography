package main

import (
	"encoding/hex"
	"fmt"

	"github.com/viza/cryptography/num"
)

func main() {

	number := num.Nat{}
	var n, n2 *num.Nat

	decodedByteArray, err := hex.DecodeString("36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80")
	if err != nil {
		fmt.Println("Unable to convert hex to byte ", err)
	} else {
		n = number.SetBytes(decodedByteArray)
		n.SetHex(string(decodedByteArray))
		fmt.Println(n)
	}

	decodedByteArray2, err := hex.DecodeString("70983d692f648185febe6d6fa607630ae68649f7e6fc45b94680096c06e4fadb")
	if err != nil {
		fmt.Println("Unable to convert hex to byte ", err)
	} else {
		n2 = number.SetBytes(decodedByteArray2)
		n.SetHex(string(decodedByteArray2))
		fmt.Println(n2)
	}

	res2, err := number.Inv(n, n2)
	fmt.Println(res2, err)

	res, _ := number.Add(n, n2, -1)
	fmt.Println(res)
	fmt.Printf("%p\n", res)
}
