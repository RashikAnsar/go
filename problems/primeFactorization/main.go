package primefactor

// PrimeFactor returns the int slice of prime factors of a given number
func PrimeFactor(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}
	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// PrimeFactor func return factors slice (f) and`
// no.of time each factor used (exponent e) slice
func primeFactor(n int) (f []int, e []int) {
	var (
		d   = 2
		len = 0
	)
	for d*d <= n && n > 1 {
		k := 0
		for n%d == 0 {
			k++
			n /= d
		}
		if k > 0 {
			len++
			f = append(f, d)
			e = append(e, k)
		}
		d++
	}
	if n > 1 {
		len++
		f = append(f, n)
		e = append(e, 1)
	}
	return
}
