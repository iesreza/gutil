package hashmap

import (
	"strings"
)

type clink struct {
	link []*clink
	data interface{}
}

func CLink() *clink {
	head := clink{link: make([]*clink, 40), data: nil}
	return &head
}

func (head *clink) InsertNode(input string, data interface{}) *interface{} {
	input = strings.ToLower(input)

	for i := len(input) - 1; i >= 0; i-- {
		if head.link[charMap(input[i])] == nil {
			head.link[charMap(input[i])] = &clink{link: make([]*clink, 40), data: nil}
		}
		head = head.link[charMap(input[i])]
	}
	head.data = data
	return &head.data
}

func (head *clink) Find(input string) (*interface{}, bool) {

	input = strings.ToLower(input)
	for i := len(input) - 1; i >= 0; i-- {
		if head.link[charMap(input[i])] != nil {
			head = head.link[charMap(input[i])]
		} else {
			return nil, false
		}
	}
	return &head.data, true
}

func charMap(char uint8) uint8 {
	if char > 96 && char < 123 {
		return char - 84
	}
	if char < 58 && char > 44 {
		return char - 45
	}
	return 100
}
