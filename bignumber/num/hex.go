package num

import (
	"encoding/hex"
	"fmt"
	"math/bits"
)

// SetHex modifies the value of z to hold a hex string, returning z
//
// The hex string must be in big endian order. If it contains characters
// other than 0..9, A..F, the value of z will be undefined, and an error will
// be returned.
//
// The value of the string shouldn't be leaked, except in the case where the string
// contains invalid characters.
func (z *Nat) SetHex(hex string) (*Nat, error) {
	z.reduced = nil
	z.announced = 4 * len(hex)
	z.limbs = z.resizedLimbs(z.announced)
	hexI := len(hex) - 1
	for i := 0; i < len(z.limbs) && hexI >= 0; i++ {
		z.limbs[i] = 0
		for shift := 0; shift < _W && hexI >= 0; shift += 4 {
			nibble, valid := nibbleFromASCII(byte(hex[hexI]))
			if valid != 1 {
				return nil, fmt.Errorf("invalid hex character: %c", hex[hexI])
			}
			z.limbs[i] |= Word(nibble) << shift
			hexI--
		}
	}
	return z, nil
}

func (z *Nat) GetHex(data []byte) string {
	encodedString := hex.EncodeToString(data)
	return encodedString
}

func (z *Nat) SetBytes(buf []byte) *Nat {
	z.reduced = nil
	z.announced = 8 * len(buf)
	z.limbs = z.resizedLimbs(z.announced)
	bufI := len(buf) - 1
	for i := 0; i < len(z.limbs) && bufI >= 0; i++ {
		z.limbs[i] = 0
		for shift := 0; shift < _W && bufI >= 0; shift += 8 {
			z.limbs[i] |= Word(buf[bufI]) << shift
			bufI--
		}
	}
	return z
}

// convert an ASCII value into a 4 bit value, returning whether or not this value is valid.
func nibbleFromASCII(ascii byte) (byte, Choice) {
	w := Word(ascii)
	inFirstRange := ctGt(w, Word('0')-1) & (1 ^ ctGt(w, Word('9')))
	inSecondRange := ctGt(w, Word('A')-1) & (1 ^ ctGt(w, Word('F')))
	valid := inFirstRange | inSecondRange
	nibble := ctIfElse(inFirstRange, w-Word('0'), w-Word('A')+0xA)
	return byte(nibble), valid
}

// ctGt checks x > y, returning 1 or 0
//
// This doesn't leak any information about either of them
func ctGt(x, y Word) Choice {
	_, b := bits.Sub(uint(y), uint(x), 0)
	return Choice(b)
}

// ctIfElse selects x if v = 1, and y otherwise
//
// This doesn't leak the value of any of its inputs
func ctIfElse(v Choice, x, y Word) Word {
	// mask should be all 1s if v is 1, otherwise all 0s
	mask := -Word(v)
	return y ^ (mask & (y ^ x))
}

// _WShift can be used to multiply or divide by _W
//
// This assumes that _W = 64, 32
const _WShift = 5 + (_W >> 6)
const _WMask = _W - 1

func (z *Nat) resizedLimbs(bits int) []Word {
	size := limbCount(bits)
	z.ensureLimbCapacity(size)
	res := z.limbs[:size]
	// Make sure that the expansion (if any) is cleared
	for i := len(z.limbs); i < size; i++ {
		res[i] = 0
	}
	maskEnd(res, bits)
	return res
}

// limbCount returns the number of limbs needed to accomodate bits.
func limbCount(bits int) int {
	return (bits + _W - 1) >> _WShift
}

// limbMask returns the mask used for the final limb of a Nat with this number of bits.
//
// Note that this function will leak the number of bits. For our library, this isn't
// a problem, since we always call this function with announced sizes.
func limbMask(bits int) Word {
	remaining := bits & _WMask
	allOnes := ^Word(0)
	if remaining == 0 {
		return allOnes
	}
	return ^(allOnes << remaining)
}

func (z *Nat) ensureLimbCapacity(size int) {
	if cap(z.limbs) < size {
		newLimbs := make([]Word, len(z.limbs), size)
		copy(newLimbs, z.limbs)
		z.limbs = newLimbs
	}
}

// maskEnd applies the correct bit mask to some limbs
func maskEnd(limbs []Word, bits int) {
	if len(limbs) <= 0 {
		return
	}
	limbs[len(limbs)-1] &= limbMask(bits)
}
