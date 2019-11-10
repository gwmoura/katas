package primefactors

func generate(n int) []int {
	primes := []int{}

	for candidate := 2; n > 1; candidate++ {
		for ; n%candidate == 0; n /= candidate {
			primes = append(primes, candidate)
		}
	}

	return primes
}
