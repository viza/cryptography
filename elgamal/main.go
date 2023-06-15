package main

import (
	"flag"
	"fmt"

	"github.com/viza/cryptography/elgamal/elgamal"
)

func main() {
	// Options flags
	encrypt_bool := flag.Bool("encrypt", true, "Option to encrypt.")
	decrypt_bool := flag.Bool("decrypt", true, "Option to decrypt.")
	sign_bool := flag.Bool("sign", true, "Option to sign.")
	verify_bool := flag.Bool("verify", true, "Option to verify.")
	debug_bool := flag.Bool("debug", true, "Allow debugging.")
	// Kpub, Kpriv flags
	alfa := flag.Int("alfa", 1234, "Parameter alfa.")
	p := flag.Int("p", 1234, "Parameter p.")
	B := flag.Int("B", 1234, "Parameter B (public)")
	b := flag.Int("b", 1234, "Parameter b (secret)")
	// Values sent flags
	Ke := flag.Int("Ke", 1234, "Ephemeral key")
	m := flag.Int("m", 1234, "Message (to encrypt or decrypt)")
	// Signature flags
	r := flag.Int("r", 1234, "Parameter r (of the signature)")
	s := flag.Int("s", 1234, "Parameter s (of the signature)")
	flag.Parse()

	if *encrypt_bool == true {
		fmt.Println("Encrypting...")
		Ke, y := elgamal.Encrypt(*p, *alfa, *B, *m, *debug_bool)
		fmt.Printf("(Ke, y) = (%v,%v)\n", Ke, y)
	} else if *decrypt_bool == true {
		fmt.Println("Decrypting...")
		x := elgamal.Decrypt(*p, *alfa, *b, *m, *Ke, *debug_bool)
		fmt.Println(x)
	} else if *sign_bool == true {
		fmt.Println("Signing...")
		r, s := elgamal.Sign(*p, *alfa, *b, *m, *debug_bool)
		fmt.Printf("(r, s) = (%v,%v)\n", r, s)
	} else if *verify_bool == true {
		fmt.Println("Verifying...")
		verified := elgamal.Verify(*p, *alfa, *B, *m, *r, *s, *debug_bool)
		fmt.Printf("Verification =  %v\n", verified)
	}

}
