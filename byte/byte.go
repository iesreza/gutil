package Bytes

import (
	"encoding/hex"
	"fmt"
	"strings"
)

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

var hexmap = map[uint8]uint8{}

func StrToByteWildcard(str string) ([]int16, error) {
	str = strings.Replace(str, " ", "", len(str))

	var res []int16
	if len(str)%2 != 0 {
		return res, fmt.Errorf("input should be divide of 2")
	}
	p := 0
	for {
		if p < len(str)/2 {
			b, err := hex.DecodeString(str[p*2 : p*2+2])
			if err == nil {
				res = append(res, int16(b[0]))
			} else {
				res = append(res, -1)
			}
		} else {
			break
		}
		p++
	}

	return res, nil

}

func StrToBytes(str string) ([]byte, error) {

	return hex.DecodeString(strings.Replace(str, " ", "", len(str)))

}
