package ujson_test

import (
	"bytes"
	"fmt"

	"github.com/ng-vu/ujson"
)

func ExampleWalk() {
	input := []byte(`{"order_id": 12345678901234, "number": 12, "item_id": 12345678905678, "counting": [1,2,3]}`)

	err := ujson.Walk(input, func(st int, key, value []byte) bool {
		fmt.Println(st, string(key), string(value))
		return true
	})
	if err != nil {
		panic(err)
	}
	// Output:
	// 0  {
	// 1 "order_id" 12345678901234
	// 1 "number" 12
	// 1 "item_id" 12345678905678
	// 1 "counting" [
	// 2  1
	// 2  2
	// 2  3
	// 1  ]
	// 0  }
}

func ExampleWalk_reconstruct() {
	input := []byte(`{"order_id": 12345678901234, "number": 12, "item_id": 12345678905678, "counting": [1,2,3]}`)

	b := make([]byte, 0, 256)
	err := ujson.Walk(input, func(st int, key, value []byte) bool {
		if len(b) != 0 && ujson.ShouldAddComma(value, b[len(b)-1]) {
			b = append(b, ',')
		}
		if len(key) > 0 {
			b = append(b, key...)
			b = append(b, ':')
		}
		b = append(b, value...)
		return true
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
	// Output: {"order_id":12345678901234,"number":12,"item_id":12345678905678,"counting":[1,2,3]}
}

func ExampleWalk_wrapInt64InString() {
	input := []byte(`{"order_id": 12345678901234, "number": 12, "item_id": 12345678905678, "counting": [1,2,3]}`)

	suffix := []byte(`_id`)
	b := make([]byte, 0, 256)
	err := ujson.Walk(input, func(_ int, key, value []byte) bool {
		// unquote key
		if len(key) != 0 {
			key = key[1 : len(key)-1]
		}

		// test for field with suffix _id
		wrap := bytes.HasSuffix(key, suffix)

		// transform the input, wrap values in double quote
		if len(b) != 0 && ujson.ShouldAddComma(value, b[len(b)-1]) {
			b = append(b, ',')
		}
		if len(key) > 0 {
			b = append(b, '"')
			b = append(b, key...)
			b = append(b, '"')
			b = append(b, ':')
		}
		if wrap {
			b = append(b, '"')
		}
		b = append(b, value...)
		if wrap {
			b = append(b, '"')
		}
		return true
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
	// Output: {"order_id":"12345678901234","number":12,"item_id":"12345678905678","counting":[1,2,3]}
}

func ExampleWalk_removeBlacklistFields() {
	input := []byte(`{"order_id": 12345678901234, "number": 12, "item_id": 12345678905678, "counting": [1,2,3]}`)

	blacklistFields := bytes.Split([]byte(`number,counting`), []byte(`,`))
	b := make([]byte, 0, 256)
	err := ujson.Walk(input, func(_ int, key, value []byte) bool {
		// unquote key and compare with blacklist
		if len(key) != 0 {
			key = key[1 : len(key)-1]
			for _, blacklist := range blacklistFields {
				if bytes.Equal(key, blacklist) {
					return false // remove the field from the output
				}
			}
		}

		// transform the input
		if len(b) != 0 && ujson.ShouldAddComma(value, b[len(b)-1]) {
			b = append(b, ',')
		}
		if len(key) > 0 {
			b = append(b, '"')
			b = append(b, key...)
			b = append(b, '"')
			b = append(b, ':')
		}
		b = append(b, value...)
		return true
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
	// Output: {"order_id":12345678901234,"item_id":12345678905678}
}
