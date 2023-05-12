package main

import (
	"fmt"

	"github.com/viza/cryptography/num"
)

func main() {

	//byteArr, err := num.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	//if err != nil {
	//	panic("Eror string")
	//} else {
	//	fmt.Printf("Array of bytes(hex): % x\n", byteArr)
	//	fmt.Printf("Array of bytes(bin): % b\n", byteArr)
	//}
	//var res num.Nat
	//(byteArr)

	//str := num.GetHex(byteArr)
	//fmt.Printf("Hex string: 0x% s\n", str)

	byteArray := []byte("1234567890")
	fmt.Println("byteArray: ", byteArray)

	num.GetHex(byteArray)

	//num.SetBytes(byteArray)
	num.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")

}
