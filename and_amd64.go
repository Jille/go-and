package and

func and(dst, a, b []byte) {
	if hasAVX2() {
		l := uint64(len(a)) >> 8
		if l != 0 {
			andAVX2(&dst[0], &a[0], &b[0], l)
		}
		l <<= 8
		dst = dst[l:]
		a = a[l:]
		b = b[l:]
	}
	andGeneric(dst, a, b)
}

func or(dst, a, b []byte) {
	if hasAVX2() {
		l := uint64(len(a)) >> 8
		if l != 0 {
			orAVX2(&dst[0], &a[0], &b[0], l)
		}
		l <<= 8
		dst = dst[l:]
		a = a[l:]
		b = b[l:]
	}
	orGeneric(dst, a, b)
}

func andNot(dst, a, b []byte) {
	if hasAVX2() {
		l := uint64(len(a)) >> 8
		if l != 0 {
			andNotAVX2(&dst[0], &a[0], &b[0], l)
		}
		l <<= 8
		dst = dst[l:]
		a = a[l:]
		b = b[l:]
	}
	andNotGeneric(dst, a, b)
}
