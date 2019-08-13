package Bytes

import "fmt"

func IndexOf(needle []byte, haystack []byte) int {
	indicator := 0
	res := -1
	for i, haystackByte := range haystack {
		if haystackByte == needle[indicator] {
			indicator++
			res = i - indicator
			if indicator == len(needle) {
				return res + 1
			}
		} else {
			indicator = 0
			res = -1
		}
	}

	return res
}

func IndexOfWildcard(needle []int16, haystack []byte) int {
	indicator := 0
	res := -1
	var p uint8

	for needle[indicator] < 0 || needle[indicator] > 255 {
		needle = needle[1:]
	}

	for i, haystackByte := range haystack {

		if needle[indicator] >= 0 && needle[indicator] <= 255 {
			p = uint8(needle[indicator])
			fmt.Println(indicator, haystackByte, p)
			if haystackByte == p {
				indicator++
				res = i - indicator
				if indicator == len(needle) {
					return i - indicator + 1
				}
			} else {
				indicator = 0
				res = -1
			}
		} else {
			fmt.Println(indicator, haystackByte, needle[indicator], "pass")
			indicator++
			res = i - indicator
			if indicator == len(needle) {
				return i - indicator + 1
			}

			continue
		}
	}

	return res

}
