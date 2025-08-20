package pokecache

import (
	"reflect"
	"testing"
)

type testData struct {
	key string
	val []byte
}

func TestCacheAddGet(t *testing.T) {
	cases := []struct {
		input    testData
		expected []byte
	}{
		{
			input:    testData{key: "foo", val: []byte{'0', '2', '4'}},
			expected: []byte{'0', '2', '4'},
		},
		{
			input:    testData{key: "bar", val: []byte{}},
			expected: []byte{},
		},
		{
			input:    testData{key: "", val: []byte{'u', 's', 'e', 'd'}},
			expected: []byte{'u', 's', 'e', 'd'},
		},
	}
	for _, c := range cases {
		cache := NewCache()
		cache.Add(c.input.key, c.input.val)
		actual, ok := cache.Get("Fantasy")
		if ok != false {
			t.Error("Get expected to return ok == false")
		}
		var nil_slice []byte
		if reflect.DeepEqual(actual, nil_slice) {
			t.Errorf("Get expected to return nil-slice []byte if ok == false, but got %v instead", actual)
		}
		actual, ok = cache.Get(c.input.key)
		if !ok {
			t.Error("Get expected to return ok")
			continue
		}

		if len(actual) != len(c.expected) {
			t.Errorf("actual and expected do not have same length for input '%v' which produced '%q': expected %d vs %d actual", c.input, actual, len(c.expected), len(actual))
			continue
		}
		for i, b := range actual {
			expectedByte := c.expected[i]
			if b != expectedByte {
				t.Errorf("actual and expected are not the same word for index %d: expected '%v' vs '%v' actual", i, expectedByte, b)
			}
		}
	}
}
