package Bytes_test

import (
	"github.com/iesreza/gutil/byte"
	"github.com/iesreza/gutil/log"
	"testing"
)

var haystack = []byte{0x48, 0x69, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x6d, 0x65, 0x20, 0x74, 0x72, 0x69, 0x65, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2e}

var byteNeedle = []byte{0x48, 0x69, 0x20}
var byteLikeNeedle = []int16{0x20, 0x74, -1, 0x69, 0x73, 0x20, 0x69, -1, 0x20, -1}

func TestByte(t *testing.T) {
	log.Error("Index of needle in haystack =  %v", Bytes.IndexOf(byteNeedle, haystack))

	//match wildcard
	log.Error("Does haystack contains byte like needle? if yes index is =  %v", Bytes.IndexOfWildcard(byteLikeNeedle, haystack))

	log.Error(Bytes.StrToBytes("00 01 02 0e"))

	w, _ := Bytes.StrToByteWildcard("20 ?? 68 69")
	log.Error(w)

	log.Error("Does haystack contains byte like needle? if yes index is =  %v", Bytes.IndexOfWildcard(w, haystack))

	//Always return -1 if not found
}