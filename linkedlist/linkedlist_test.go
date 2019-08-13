package linkedlist_test

import (
	"fmt"
	"github.com/iesreza/gutil/linkedlist"
	"github.com/iesreza/gutil/log"
	"testing"
	"time"
)

type Session struct {
	Expire int64
	ID     string
}

func TestLinkedList(t *testing.T) {
	var list = linkedlist.List{}

	list.SetMatchFunc(func(needle interface{}, el interface{}) bool {
		return needle.(Session).ID == el.(Session).ID
	})

	list.PushOnce(Session{1558378380, "abcd1"})
	list.PushOnce(Session{1568378380, "abcd2"})
	list.PushOnce(Session{1538378380, "abcd3"})
	list.PushOnce(Session{1933378380, "abcd4"})
	fmt.Println(list.String())

	log.Error("Find abcd4: %+v", list.Find(Session{ID: "abcd4"}))
	log.Error("Find not in list: %+v", list.Find(Session{ID: "Not in list"}))

	var v int64
	v = time.Now().Unix()
	list.RemoveFunc(v, func(needle interface{}, el interface{}) bool {

		return needle.(int64) < el.(Session).Expire
	})

	list.Remove(Session{1933378380, "abcd4"})

	log.Info("Remove abcd4")
	fmt.Println(list.String())
}
