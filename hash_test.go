package ntest

/**
赎回金
*/

func canConstruct(ransomNote string, magazine string) bool {
	// 字母和次数
	m := make(map[int32]int, 26)
	for _, b := range magazine {
		m[b]++
	}
	for _, b := range ransomNote {
		m[b]--
		if m[b] < 0 {
			return false
		}
	}
	return true
}

/*
*
同构字符串
*/
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapsS, mapsT := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s); i++ {
		if mapsS[s[i]] != mapsT[t[i]] {
			return false
		}
		mapsS[s[i]] = i + 1
		mapsT[t[i]] = i + 1
	}
	return true
}
