package num

import (
	"fmt"
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

// Add calculates z <- x + y, modulo 2^cap
//
// The capacity is given in bits, and also controls the size of the result.
//
// If cap < 0, the capacity will be max(x.AnnouncedLen(), y.AnnouncedLen()) + 1
func (z *Nat) Add(x *Nat, y *Nat, cap int) (*Nat, error) {
	if cap < 0 {
		cap = x.maxAnnounced(y) + 1
	}
	xLimbs := x.resizedLimbs(cap)
	yLimbs := y.resizedLimbs(cap)
	z.limbs = z.resizedLimbs(cap)
	addVV(z.limbs, xLimbs, yLimbs)
	// Mask off the final bits
	z.limbs = z.resizedLimbs(cap)
	z.announced = cap
	z.reduced = nil
	return z, nil //TODO: error handling
}

// maxAnnounced returns the larger announced length of z and y
func (z *Nat) maxAnnounced(y *Nat) int {
	maxBits := z.announced
	if y.announced > maxBits {
		maxBits = y.announced
	}
	return maxBits
}

func addVV(z, x, y []Word) (c Word) {
	return addVV_g(z, x, y)
}

// The resulting carry c is either 0 or 1.
func addVV_g(z, x, y []Word) (c Word) {
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x) && i < len(y); i++ {
		zi, cc := bits.Add(uint(x[i]), uint(y[i]), uint(c))
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

func (z *Nat) Sub(x *Nat, y *Nat, cap int) (*Nat, error) {

	fmt.Println("TODO: Implement")
	return nil, nil
}

func (z *Nat) Mod(x *Nat, cap int) (*Nat, error) {

	fmt.Println("TODO: Implement")
	return nil, nil
}

func (z *Nat) Mul(x *Nat, y *Nat, cap int) (*Nat, error) {

	fmt.Println("TODO: Implement")
	return nil, nil
}

func (z *Nat) Div(x *Nat, y *Nat, cap int) (*Nat, error) {

	fmt.Println("TODO: Implement")
	return nil, nil
}

func (z *Nat) PowMod(x *Nat, y int, cap int) (*Nat, error) {

	fmt.Println("TODO: Implement")
	return nil, nil
}
