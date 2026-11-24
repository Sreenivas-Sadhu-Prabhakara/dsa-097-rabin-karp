package main

import "fmt"

// Search returns every index where pattern occurs in text, found with
// Rabin–Karp: a rolling polynomial hash compares windows in O(1), and only
// hash hits are verified character by character.
func Search(text, pattern string) []int {
	out := []int{}
	n, m := len(text), len(pattern)
	if m == 0 || m > n {
		return out
	}
	const base = 256
	const mod = 1_000_000_007
	// pow = base^(m-1) % mod, the weight of the outgoing character.
	pow := 1
	for i := 0; i < m-1; i++ {
		pow = pow * base % mod
	}
	ph, th := 0, 0
	for i := 0; i < m; i++ {
		ph = (ph*base + int(pattern[i])) % mod
		th = (th*base + int(text[i])) % mod
	}
	for i := 0; ; i++ {
		if th == ph && text[i:i+m] == pattern {
			out = append(out, i)
		}
		if i+m == n {
			break
		}
		// Slide: remove text[i], append text[i+m].
		th = ((th-int(text[i])*pow%mod+mod)%mod*base + int(text[i+m])) % mod
	}
	return out
}

func main() {
	fmt.Println(Search("ababcababc", "abc"))
}
