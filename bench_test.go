package smooth_test

import (
  "testing"
  "smooth"
  "sort"
  "rand"
)

const size = 10000
var shuffled []int
func init() {
  shuffled = make([]int, size)
  for i := range shuffled {
    shuffled[i] = rand.Int()
  }
}

func fillReverse(v []int) {
  for i := range v {
    v[i] = len(v) - i - 1
  }
}

func fill(v []int) {
  for i := range v {
    v[i] = i
  }
}

func BenchmarkCopy(b *testing.B) {
  v := make([]int, size)
  v2 := make([]int, size)
  for i := 0; i < b.N; i++ {
    for i := range v {
      v[i] = v2[i]
    }
  }
}

func BenchmarkQuicksortOnUnsorted(b *testing.B) {
  v := make([]int, size)
  for i := 0; i < b.N; i++ {
    fillReverse(v)
    sort.Ints(v)
  }
}

func BenchmarkSmoothsortOnUnsorted(b *testing.B) {
  v := make([]int, size)
  for i := 0; i < b.N; i++ {
    fillReverse(v)
    smooth.Ints(v)
  }
}

func BenchmarkQuicksortOnSorted(b *testing.B) {
  v := make([]int, size)
  fill(v)
  for i := 0; i < b.N; i++ {
    sort.Ints(v)
  }
}

func BenchmarkSmoothsortOnSorted(b *testing.B) {
  v := make([]int, size)
  fill(v)
  for i := 0; i < b.N; i++ {
    smooth.Ints(v)
  }
}

func BenchmarkQuicksortOnShuffled(b *testing.B) {
  v := make([]int, size)
  for i := 0; i < b.N; i++ {
    for i := range v {
      v[i] = shuffled[i]
    }
    sort.Ints(v)
  }
}

func BenchmarkSmoothsortOnShuffled(b *testing.B) {
  v := make([]int, size)
  for i := 0; i < b.N; i++ {
    for i := range v {
      v[i] = shuffled[i]
    }
    smooth.Ints(v)
  }
}

func BenchmarkSmoothsortOnSorted1000(b *testing.B) {
  v := make([]int, 1000)
  fill(v)
  for i := 0; i < b.N; i++ {
    smooth.Ints(v)
  }
}

func BenchmarkSmoothsortOnSorted10000(b *testing.B) {
  v := make([]int, 10000)
  fill(v)
  for i := 0; i < b.N; i++ {
    smooth.Ints(v)
  }
}

func BenchmarkSmoothsortOnSorted100000(b *testing.B) {
  v := make([]int, 100000)
  fill(v)
  for i := 0; i < b.N; i++ {
    smooth.Ints(v)
  }
}

func BenchmarkSmoothsortOnSorted1000000(b *testing.B) {
  v := make([]int, 1000000)
  fill(v)
  for i := 0; i < b.N; i++ {
    smooth.Ints(v)
  }
}

