//go:build (!amd64 && !arm64) || purego

package and

func and(dst, a, b []byte) {
	andGeneric(dst, a, b)
}

func or(dst, a, b []byte) {
	orGeneric(dst, a, b)
}

func andNot(dst, a, b []byte) {
	andNotGeneric(dst, a, b)
}

func popcnt(a []byte) int {
	return popcntGeneric(a)
}
