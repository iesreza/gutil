package linkedlist

import (
	"fmt"
	"strings"
)

type MatchFunc func(needle interface{}, el interface{}) bool
type Node struct {
	prev  *Node
	next  *Node
	Value interface{}
}

type List struct {
	head   *Node
	tail   *Node
	length int
	Match  MatchFunc
}

func (L *List) Prepend(key interface{}) {
	newHead := &Node{
		next:  L.head,
		Value: key,
	}

	// Keep track of oldHead
	oldHead := L.head

	// Always insert new head
	L.head = newHead

	if oldHead == nil {
		// If oldHead was originally null, also make it the tail
		L.tail = newHead
	} else {
		// If list oldHead is not empty, add a 'prev' node entry to the oldHead
		oldHead.prev = newHead
	}
	L.length++
}

func (L *List) Push(key interface{}) {
	L.Append(key)
}

func (L *List) PushOnce(key interface{}) {
	if !L.Contains(key) {
		L.Append(key)
	}
}

func (L *List) Remove(needle interface{}) {
	if L.Match != nil {
		L.RemoveFunc(needle, L.Match)
	} else {
		for e := L.First(); e != nil; e = e.Next() {

			if L.Match != nil && L.Match(needle, e.Value) {

				if e != nil {
					if e.prev != nil {
						e.prev.next = e.next
					} else {
						L.head = e.next
					}
					if e.next != nil {
						e.next.prev = e.prev
					} else {
						L.tail = e.prev
					}

					L.length--
				}

			}
		}
	}
}

func (L *List) RemoveFunc(needle interface{}, match MatchFunc) {

	for e := L.First(); e != nil; e = e.Next() {

		if match(needle, e.Value) {

			if e != nil {
				if e.prev != nil {
					e.prev.next = e.next
				} else {
					L.head = e.next
				}
				if e.next != nil {
					e.next.prev = e.prev
				} else {
					L.tail = e.prev
				}

				L.length--
			}

		}
	}

}

func (L *List) Append(key interface{}) {
	newTail := &Node{
		prev:  L.tail,
		next:  nil,
		Value: key,
	}

	oldTail := L.tail
	L.tail = newTail

	if oldTail == nil {

		L.head = newTail
	} else {
		oldTail.next = newTail
	}
	L.length++

}

func (L *List) Find(needle interface{}) []interface{} {

	if L.Match != nil {
		return L.FindFunc(needle, L.Match)
	} else {
		var ls []interface{}
		for e := L.First(); e != nil; e = e.Next() {
			if fmt.Sprint(e.Value) == fmt.Sprint(needle) {
				ls = append(ls, e.Value)
			}
		}
		return ls
	}

}

func (L *List) Contains(needle interface{}) bool {
	for e := L.First(); e != nil; e = e.Next() {
		if L.Match == nil {
			if fmt.Sprint(e.Value) == fmt.Sprint(needle) {
				return true
			}
		} else {
			if L.Match(needle, e.Value) {
				return true
			}
		}
	}
	return false
}

func (L *List) ContainsFunc(needle interface{}, match MatchFunc) bool {
	for e := L.First(); e != nil; e = e.Next() {
		if match(needle, e.Value) {
			return true
		}
	}
	return false
}

func (L *List) FindFunc(needle interface{}, match MatchFunc) []interface{} {
	var ls []interface{}
	for e := L.First(); e != nil; e = e.Next() {
		if match(needle, e.Value) {
			ls = append(ls, e.Value)
		}
	}
	return ls
}

func (L *List) First() *Node {
	return L.head
}

func (L *List) Last() *Node {

	return L.tail
}

func (node *Node) Prev() *Node {
	return node.prev
}

func (node *Node) Next() *Node {
	if node == nil {
		return nil
	}
	return node.next
}

func (L *List) SetMatchFunc(fn MatchFunc) {
	L.Match = fn
}

func (l *List) String() string {
	res := ""
	list := l.head
	for list != nil {
		res += fmt.Sprintf("%+v,", list.Value)
		list = list.next
	}
	return strings.TrimRight(res, ",")
}

func String(list *Node) string {
	res := ""
	for list != nil {
		res += fmt.Sprintf("%v,", list.Value)
		list = list.next
	}
	return strings.TrimRight(res, ",")
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *List) Length() int {
	return l.length
}
