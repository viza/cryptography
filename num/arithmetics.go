package num

import (
	"math/bits"
	"sync"
)

const (
	_S = _W / 8 // word size in bytes

	_W = bits.UintSize // word size in bits
	_B = 1 << _W       // digit base
	_M = _B - 1        // digit mask
)

// A Word represents a single digit of a multi-precision unsigned integer.
type Word uint

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

type Modulus struct {
	nat Nat
	// the number of leading zero bits
	leading int
	// The inverse of the least significant limb, modulo W
	m0inv Word
	// If true, then this modulus is even
	even bool
}

// Choice represents a constant-time boolean.
//
// The value of Choice is always either 1 or 0.
//
// We use a separate type instead of bool, in order to be able to make decisions without leaking
// which decision was made.
//
// You can easily convert a Choice into a bool with the operation c == 1.
//
// In general, logical operations on bool become bitwise operations on choice:
//
//	a && b => a & b
//	a || b => a | b
//	a != b => a ^ b
//	!a     => 1 ^ a
type Choice Word

func PutNat(x *Nat) {
	natPool.Put(x)
}

var natPool sync.Pool
