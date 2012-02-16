package smooth_test

import (
	"math"
	"math/rand"
	"testing"
)

import . "github.com/runningwild/sorts/smooth"

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestSortIntSlice(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	Sort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestSortFloat64Slice(t *testing.T) {
	data := float64s
	a := Float64Slice(data[0:])
	Sort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", float64s)
		t.Errorf("   got %v", data)
	}
}

func TestSortStringSlice(t *testing.T) {
	data := strings
	a := StringSlice(data[0:])
	Sort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", strings)
		t.Errorf("   got %v", data)
	}
}

func TestInts(t *testing.T) {
	data := ints
	Ints(data[0:])
	if !IntsAreSorted(data[0:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestFloat64s(t *testing.T) {
	data := float64s
	Float64s(data[0:])
	if !Float64sAreSorted(data[0:]) {
		t.Errorf("sorted %v", float64s)
		t.Errorf("   got %v", data)
	}
}

func TestStrings(t *testing.T) {
	data := strings
	Strings(data[0:])
	if !StringsAreSorted(data[0:]) {
		t.Errorf("sorted %v", strings)
		t.Errorf("   got %v", data)
	}
}

func TestSortLarge_Random(t *testing.T) {
	n := 1000000
	if testing.Short() {
		n /= 100
	}
	data := make([]int, n)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
	}
	if IntsAreSorted(data) {
		t.Fatalf("terrible rand.rand")
	}
	Ints(data)
	if !IntsAreSorted(data) {
		t.Errorf("sort didn't sort - 1M ints")
	}
}

func shuffle(v []int) {
	var n int
	for i := len(v) - 1; i > 0; i-- {
		n = rand.Intn(i)
		v[i], v[n] = v[n], v[i]
	}
}

func partialShuffle(v []int, n int) {
	for i := len(v) - 1; n > 0; i-- {
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

func BenchmarkSorted10(b *testing.B) {
	v := make([]int, 10)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		Ints(v)
	}
}

func BenchmarkMostlySorted10(b *testing.B) {
	b.StopTimer()
	v := make([]int, 10)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		partialShuffle(v, 5)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkShuffled10(b *testing.B) {
	b.StopTimer()
	v := make([]int, 10)
	for i := 0; i < b.N; i++ {
		shuffle(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkReversed10(b *testing.B) {
	b.StopTimer()
	v := make([]int, 10)
	for i := 0; i < b.N; i++ {
		reverse(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkSorted100(b *testing.B) {
	v := make([]int, 100)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		Ints(v)
	}
}

func BenchmarkMostlySorted100(b *testing.B) {
	b.StopTimer()
	v := make([]int, 100)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		partialShuffle(v, 5)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkShuffled100(b *testing.B) {
	b.StopTimer()
	v := make([]int, 100)
	for i := 0; i < b.N; i++ {
		shuffle(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkReversed100(b *testing.B) {
	b.StopTimer()
	v := make([]int, 100)
	for i := 0; i < b.N; i++ {
		reverse(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkSorted1k(b *testing.B) {
	v := make([]int, 1000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		Ints(v)
	}
}

func BenchmarkMostlySorted1k(b *testing.B) {
	b.StopTimer()
	v := make([]int, 1000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		partialShuffle(v, 5)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkShuffled1k(b *testing.B) {
	b.StopTimer()
	v := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		shuffle(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkReversed1k(b *testing.B) {
	b.StopTimer()
	v := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		reverse(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkSorted10k(b *testing.B) {
	v := make([]int, 10000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		Ints(v)
	}
}

func BenchmarkMostlySorted10k(b *testing.B) {
	b.StopTimer()
	v := make([]int, 10000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		partialShuffle(v, 5)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkShuffled10k(b *testing.B) {
	b.StopTimer()
	v := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		shuffle(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkReversed10k(b *testing.B) {
	b.StopTimer()
	v := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		reverse(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkSorted100k(b *testing.B) {
	v := make([]int, 100000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		Ints(v)
	}
}

func BenchmarkMostlySorted100k(b *testing.B) {
	b.StopTimer()
	v := make([]int, 100000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		partialShuffle(v, 5)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkShuffled100k(b *testing.B) {
	b.StopTimer()
	v := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		shuffle(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkReversed100k(b *testing.B) {
	b.StopTimer()
	v := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		reverse(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkSorted1M(b *testing.B) {
	v := make([]int, 1000000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		Ints(v)
	}
}

func BenchmarkMostlySorted1M(b *testing.B) {
	b.StopTimer()
	v := make([]int, 1000000)
	inOrder(v)
	for i := 0; i < b.N; i++ {
		partialShuffle(v, 5)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkShuffled1M(b *testing.B) {
	b.StopTimer()
	v := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		shuffle(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}

func BenchmarkReversed1M(b *testing.B) {
	b.StopTimer()
	v := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		reverse(v)
		b.StartTimer()
		Ints(v)
		b.StopTimer()
	}
}
