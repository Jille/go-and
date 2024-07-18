// and implements bitwise and for two byte-slices.
package and

import (
	"encoding/binary"
	"math/bits"
)

// Writes bitwise and of a and b to dst.
//
// Panics if len(a) ≠ len(b), or len(dst) ≠ len(a).
func And(dst, a, b []byte) {
	if len(a) != len(b) || len(b) != len(dst) {
		panic("lengths of a, b and dst must be equal")
	}

	and(dst, a, b)
}

func andGeneric(dst, a, b []byte) {
	i := 0

	for ; i <= len(a)-8; i += 8 {
		binary.LittleEndian.PutUint64(
			dst[i:],
			binary.LittleEndian.Uint64(a[i:])&binary.LittleEndian.Uint64(b[i:]),
		)
	}

	for ; i < len(a); i++ {
		dst[i] = a[i] & b[i]
	}
}

// Writes bitwise or of a and b to dst.
//
// Panics if len(a) ≠ len(b), or len(dst) ≠ len(a).
func Or(dst, a, b []byte) {
	if len(a) != len(b) || len(b) != len(dst) {
		panic("lengths of a, b and dst must be equal")
	}

	or(dst, a, b)
}

func orGeneric(dst, a, b []byte) {
	i := 0

	for ; i <= len(a)-8; i += 8 {
		binary.LittleEndian.PutUint64(
			dst[i:],
			binary.LittleEndian.Uint64(a[i:])|binary.LittleEndian.Uint64(b[i:]),
		)
	}

	for ; i < len(a); i++ {
		dst[i] = a[i] | b[i]
	}
}

// Writes bitwise and of not(a) and b to dst.
//
// Panics if len(a) ≠ len(b), or len(dst) ≠ len(a).
func AndNot(dst, a, b []byte) {
	if len(a) != len(b) || len(b) != len(dst) {
		panic("lengths of a, b and dst must be equal")
	}

	andNot(dst, a, b)
}

func andNotGeneric(dst, a, b []byte) {
	i := 0

	for ; i <= len(a)-8; i += 8 {
		binary.LittleEndian.PutUint64(
			dst[i:],
			(^binary.LittleEndian.Uint64(a[i:]))&binary.LittleEndian.Uint64(b[i:]),
		)
	}

	for ; i < len(a); i++ {
		dst[i] = (^a[i]) & b[i]
	}
}

// Writes bitwise and of not(a) and b to dst.
//
// Panics if len(a) ≠ len(b), or len(dst) ≠ len(a).
func Popcnt(a []byte) int {
	return popcnt(a)
}

func popcntGeneric(a []byte) int {
	var ret int
	i := 0

	for ; i <= len(a)-8; i += 8 {
		ret += bits.OnesCount64(binary.LittleEndian.Uint64(a[i:]))
	}

	for ; i < len(a); i++ {
		ret += bits.OnesCount8(a[i])
	}
	return ret
}
