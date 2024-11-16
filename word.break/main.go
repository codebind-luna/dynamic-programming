package main

func canBreak(idx int, s string, m map[string]bool) bool {
	if idx == len(s) {
		return true
	}

	if m[s[idx:]] {
		return true
	}

	for j := 0; idx+j < len(s); j++ {
		if m[s[idx:idx+j+1]] && canBreak(idx+j+1, s, m) {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, v := range wordDict {
		words[v] = true
	}
	return canBreak(0, s, words)
}
