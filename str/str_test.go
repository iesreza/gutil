package str

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	test := []byte(`   test \  string   `)
	fmt.Println(S(test).TrimSpace().ReplaceAll("t", "f").Quote())

	test2 := map[string]int{
		"Test": 20,
	}

	fmt.Println(S(test2).ReplaceAll("20", "19"))

}
