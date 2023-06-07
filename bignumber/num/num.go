package num

type Nat struct {
	// The exact number of bits this number claims to have.
	//
	// This can differ from the actual number of bits needed to represent this number.
	announced int
	// If this is set, then the value of this Nat is in the range 0..reduced - 1.
	//
	// This value should get set based only on statically knowable things, like what
	// functions have been called. This means that we will have plenty of false
	// negatives, where a value is small enough, but we don't know statically
	// that this is the case.
	//
	// Invariant: If reduced is set, then announced should match the announced size of
	// this modulus.
	reduced *Modulus
	// The limbs representing this number, in little endian order.
	//
	// Invariant: The bits past announced will not be set. This includes when announced
	// isn't a multiple of the limb size.
	//
	// Invariant: two Nats are not allowed to share the same slice.
	// This allows us to use pointer comparison to check that Nats don't alias eachother
	limbs []Word
}

type Number interface {
	// work with hex string and unsigned big number
	SetHex(string) (*Nat, error)
	GetHex([]byte) string
	SetBytes([]byte) *Nat

	//Bitwise operations
	Inv(*Nat, *Nat) (*Nat, error)
	Xor(*Nat, *Nat) (*Nat, error)
	Or(*Nat, *Nat) (*Nat, error)
	And(*Nat, *Nat) (*Nat, error)
	ShiftL(*Nat, int) (*Nat, error)
	ShiftR(*Nat, int) (*Nat, error)

	//Arifmetic operations
	Add(*Nat, *Nat, int) (*Nat, error)
	Sub(*Nat, *Nat, int) (*Nat, error)
	Mul(*Nat, *Nat, int) (*Nat, error)
	Div(*Nat, *Nat, int) (*Nat, error)
	PowMod(*Nat, int, int) (*Nat, error)
}
