package smooth_test

import (
	"testing"
	"github.com/runningwild/sorts/smooth"
	"sort"
	"rand"
)

const size = 1000

var shuffled []int

func init() {
	shuffled = make([]int, size)
	for i := range shuffled {
		shuffled[i] = rand.Int()
	}
}

func shuffle(v []int) {
	var n int
	for i := len(v)-1; i > 0; i-- {
		n = rand.Intn(i)
		v[i], v[n] = v[n], v[i]
	}
}

func partialShuffle(v []int, n int) {
	for i := len(v)-1; n > 0; i-- {
		r := rand.Intn(i)
		v[i], v[r] = v[r], v[i]
		n--
	}
}

func reverse(v []int) {
	for i := range v {
		v[i] = len(v) - i - 1
	}
}

func inOrder(v []int) {
	for i := range v {
		v[i] = i
	}
}

func BenchmarkQuicksortReversed(b *testing.B) {
	b.StopTimer()
	v := make([]int, size)
	for i := 0; i < b.N; i++ {
		reverse(v)
		b.StartTimer()
		sort.Ints(v)
		b.StopTimer()
	}
}

func BenchmarkSmoothsortReversed(b *testing.B) {
	b.StopTimer()
	v := make([]int, size)
	for i := 0; i < b.N; i++ {
		reverse(v)
		b.StartTimer()
		smooth.Ints(v)
		b.StopTimer()
	}
}


func BenchmarkQuicksortInOrder(b *testing.B) {
	v := make([]int, size)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		sort.Ints(v)
	}
}

func BenchmarkSmoothsortInOrder(b *testing.B) {
	v := make([]int, size)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		smooth.Ints(v)
	}
}


func BenchmarkQuicksortShuffled(b *testing.B) {
	b.StopTimer()
	v := make([]int, size)
	for i := 0; i < b.N; i++ {
		shuffle(v)
		b.StartTimer()
		sort.Ints(v)
		b.StopTimer()
	}
}

func BenchmarkSmoothsortShuffled(b *testing.B) {
	b.StopTimer()
	v := make([]int, size)
	for i := 0; i < b.N; i++ {
		shuffle(v)
		b.StartTimer()
		smooth.Ints(v)
		b.StopTimer()
	}
}

func BenchmarkQuicksortMostlySorted(b *testing.B) {
	b.StopTimer()
	v := make([]int, size)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		partialShuffle(v, 5)
		b.StartTimer()
		sort.Ints(v)
		b.StopTimer()
	}
}

func BenchmarkSmoothsortMostlySorted(b *testing.B) {
	b.StopTimer()
	v := make([]int, size)
	for i := 0; i < b.N; i++ {
		partialShuffle(v, 5)
		b.StartTimer()
		smooth.Ints(v)
		b.StopTimer()
	}
}

func BenchmarkQuicksortInOrder1k(b *testing.B) {
	v := make([]int, 1000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		sort.Ints(v)
	}
}


func BenchmarkSmoothsortInOrder1k(b *testing.B) {
	v := make([]int, 1000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		smooth.Ints(v)
	}
}

func BenchmarkQuicksortOnInOrder10k(b *testing.B) {
	v := make([]int, 10000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		sort.Ints(v)
	}
}

func BenchmarkSmoothsortOnInOrder10k(b *testing.B) {
	v := make([]int, 10000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		smooth.Ints(v)
	}
}

func BenchmarkQuicksortOnInOrder100k(b *testing.B) {
	v := make([]int, 100000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		sort.Ints(v)
	}
}

func BenchmarkSmoothsortOnInOrder100k(b *testing.B) {
	v := make([]int, 100000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		smooth.Ints(v)
	}
}

func BenchmarkQuicksortOnSorted1M(b *testing.B) {
	v := make([]int, 1000000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		sort.Ints(v)
	}
}

func BenchmarkSmoothsortOnSorted1M(b *testing.B) {
	v := make([]int, 1000000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		smooth.Ints(v)
	}
}
