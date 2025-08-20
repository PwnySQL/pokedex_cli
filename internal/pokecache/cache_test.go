package pokecache

import (
	"reflect"
	"testing"
	"time"
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
		cache := NewCache(5 * time.Second)
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

func TestCacheReapLoop(t *testing.T) {
	cases := []struct {
		testDataBefore []testData
		testDataAfter  []testData
	}{
		{
			testDataBefore: []testData{{key: "foo", val: []byte{'0', '2', '4'}}},
			testDataAfter:  []testData{{key: "bar", val: []byte{'1', '2', '3'}}},
		},
		{
			testDataBefore: []testData{},
			testDataAfter:  []testData{{key: "bar", val: []byte{'1', '2', '3'}}},
		},
		{
			testDataBefore: []testData{{key: "foo", val: []byte{'0', '2', '4'}}},
			testDataAfter:  []testData{},
		},
		{
			testDataBefore: []testData{},
			testDataAfter:  []testData{},
		},
	}
	for _, c := range cases {
		waitDuration := 4 * time.Millisecond
		reapInterval := waitDuration - time.Millisecond
		// NewCache would also reap the cache. Hence, make its reapInterval longer such that the explicit call
		// to reapLoop reaps the cache (after adding all data) to be able to test that reapLoop does only reap
		// "stale" data aka data older than the interval.
		cache := NewCache(waitDuration + time.Millisecond)
		for _, data := range c.testDataBefore {
			cache.Add(data.key, data.val)
		}
		time.Sleep(waitDuration)
		for _, data := range c.testDataAfter {
			cache.Add(data.key, data.val)
		}

		if len(cache.cache) != len(c.testDataBefore)+len(c.testDataAfter) {
			t.Error("Sanity check before reaping: All data is still in the cache")
		}
		cache.reapLoop(reapInterval)

		if len(cache.cache) != len(c.testDataAfter) {
			t.Error("The number of elements in the cache and in the data added after the sleep must be the same")
		}

		for _, data := range c.testDataBefore {
			_, isKeyInCache := cache.Get(data.key)
			if isKeyInCache {
				t.Error("Test data before sleep should be removed")
			}
		}
		for _, data := range c.testDataAfter {
			actual, isKeyInCache := cache.Get(data.key)
			if !isKeyInCache {
				t.Error("Test data after sleep must be still in cache")
			}
			if !reflect.DeepEqual(data.val, actual) {
				t.Error("Test data values should be unaltered by reapLoop")
			}
		}
	}
}

func TestCacheReapInNewCache(t *testing.T) {
	cases := []struct {
		testDataBefore []testData
		testDataAfter  []testData
	}{
		{
			testDataBefore: []testData{{key: "foo", val: []byte{'0', '2', '4'}}},
			testDataAfter:  []testData{{key: "bar", val: []byte{'1', '2', '3'}}},
		},
	}
	for _, c := range cases {
		waitDuration := 4 * time.Millisecond
		reapInterval := 2 * time.Millisecond
		// NewCache triggers reaping the returned channel.
		cache := NewCache(reapInterval)
		for _, data := range c.testDataBefore {
			cache.Add(data.key, data.val)
		}
		time.Sleep(waitDuration)
		for _, data := range c.testDataAfter {
			cache.Add(data.key, data.val)
		}

		if len(cache.cache) != len(c.testDataAfter) {
			t.Error("The number of elements in the cache and in the data added after the sleep must be the same")
		}

		for _, data := range c.testDataBefore {
			_, isKeyInCache := cache.Get(data.key)
			if isKeyInCache {
				t.Error("Test data before sleep should be removed")
			}
		}
		for _, data := range c.testDataAfter {
			actual, isKeyInCache := cache.Get(data.key)
			if !isKeyInCache {
				t.Error("Test data after sleep must be still in cache")
			}
			if !reflect.DeepEqual(data.val, actual) {
				t.Error("Test data values should be unaltered by reapLoop")
			}
		}
	}
}
