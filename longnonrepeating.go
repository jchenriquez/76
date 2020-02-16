package longestnonrepeating

import (
	"math"
)

// LengthOfLongestSubstring will return the longest non repeating string.
func LengthOfLongestSubstring(s string) (maxLen int) {
	if len(s) == 0 {
		return
	}

	mp := make(map[byte]int, len(s))
	var ln int

	mp[s[0]] = 0
	ln++
	maxLen = ln

	for i := 1; i < len(s); i++ {
		index, in := mp[s[i]]

		if in {
			maxLen = int(math.Max(float64(ln), float64(maxLen)))
			for val, ind := range mp {
				if ind <= index {
					delete(mp, val)
				}
			}
			ln = len(mp)
		}
		mp[s[i]] = i
		ln++
	}

	maxLen = int(math.Max(float64(ln), float64(maxLen)))

	return
}
