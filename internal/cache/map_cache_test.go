package cache

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/gopher-translator-service/pkg/assert"
)

func deepEqual(t assert.ErrorAsserter, a, b []KeyValue) {
    // If one is nil, the other must also be nil.
    if (a == nil) != (b == nil) {
        t.Fatalf("DeepEqual failed for %v %v", a, b)
    }

    if len(a) != len(b) {
        t.Fatalf("DeepEqual failed for %v %v", a, b)
    }

    for i := range a {
        if a[i] != b[i] {
            t.Fatalf("DeepEqual failed for %v %v", a, b)
        }
    }
}

func TestCacheAddOne(t *testing.T) {
	m := NewMapCache()
	assert.AssertEqual(t, m.Len(), 0)

	m.Add("abc", "def")
	assert.AssertEqual(t, m.Len(), 1)
	assert.AssertEqual(t, m.Get("abc"), "def")
	deepEqual(t, m.GetAll(), []KeyValue{ {Key: "abc", Value: "def"} })
}

func BenchmarkCacheAddManyInParallel(t *testing.B) {
	const n = 100000
	m := NewMapCache()
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			v := strconv.Itoa(i)
			m.Add(v, v)
		}(i)
	}

	wg.Wait()
	assert.AssertEqual(t, m.Len(), n)
	errMsg := fmt.Sprintf("Not everything in the range [0..%d] was saved in the cache", n)
	for i := 0; i < n; i++ {
		v := strconv.Itoa(i)
		assert.AssertTrue(t, m.Get(v) != "", errMsg)
	}
}