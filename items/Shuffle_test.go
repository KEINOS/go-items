package items_test

import (
	"hash/maphash"
	"math/rand"
	"testing"
	"time"

	"github.com/KEINOS/go-items/items"
	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	// Items to be shuffled
	list := &items.TItems{
		Items:  genSampleSlice(t),
		Sorted: false,
	}

	// Items to compare
	orig := &items.TItems{
		Items:  genSampleSlice(t),
		Sorted: false,
	}

	assert.Equal(t, orig, list, "it should be equal before shuffle")

	list.Shuffle()

	t.Logf("Shuffled: %#v\n", list)
	t.Logf("Original: %#v\n", orig)

	assert.True(t, list.Sorted, "it should be true after shuffle")

	for _, v := range orig.Items {
		assert.Contains(t, list.Items, v, "it should contain equal elements")
	}

	assert.NotEqual(t, orig.Items, list.Items, "it should be shuffled")
}

// ----------------------------------------------------------------------------
//  Benchmarks
// ----------------------------------------------------------------------------

// Benches for randomizer.
// Even though maphash method was faster than unixnano, there was not a big
// difference in the results. So we just use unixnano for now.
//
//   $ go test -bench=. ./items
//   goos: linux
//   goarch: amd64
//   pkg: github.com/qithub/qiitask/items
//   cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
//   BenchmarkRandomizer_unixnano1-2           114583             10577 ns/op
//   BenchmarkRandomizer_unixnano2-2           101758             10484 ns/op
//   BenchmarkRandomizer_maphash1-2            100492             10450 ns/op
//   BenchmarkRandomizer_maphash2-2            101604             10483 ns/op
//   PASS
//   ok      github.com/qithub/qiitask/items 5.901s

func BenchmarkRandomizer_unixnano1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//nolint:gosec // using of weak random number generator here is ok
		_ = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
}

func BenchmarkRandomizer_unixnano2(b *testing.B) {
	rand64 := func() int64 { return time.Now().UnixNano() }

	for i := 0; i < b.N; i++ {
		//nolint:gosec // using of weak random number generator here is ok
		_ = rand.New(rand.NewSource(rand64()))
	}
}

func BenchmarkRandomizer_maphash1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//nolint:gosec // using of weak random number generator here is ok
		_ = rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	}
}

func BenchmarkRandomizer_maphash2(b *testing.B) {
	rand64 := func() int64 { return int64(new(maphash.Hash).Sum64()) }

	for i := 0; i < b.N; i++ {
		//nolint:gosec // using of weak random number generator here is ok
		_ = rand.New(rand.NewSource(rand64()))
	}
}

// ----------------------------------------------------------------------------
//  Helper functions
// ----------------------------------------------------------------------------

// The genSampleSlice function is used to generate a sample slice for testing.
func genSampleSlice(t *testing.T) []interface{} {
	t.Helper()

	return []interface{}{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
	}
}
