package array

func FindPrimes(n int) []int {
	if n < 2 {
		return []int{}
	}
	if n < 3 {
		return []int{2}
	}
	primes := make([]bool, n+1)
	for i := 0; i < n+1; i += 1 {
		primes[i] = false
	}
	prime_nums := make([]int, 0)
	for i := 2; i < n+1; i += 1 {
		if !primes[i] {
			prime_nums = append(prime_nums, i)
			for j := 2 * i; j < n+1; j += i {
				primes[j] = true
			}
		}
	}
	return prime_nums
}
