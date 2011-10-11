package smooth_test

import (
  "testing"
  "smooth"
  "sort"
)

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

func BenchmarkQuicksortOnUnsorted(b *testing.B) {
  v := make([]int, 1000000)
  for i := 0; i < b.N; i++ {
    fillReverse(v)
    sort.Ints(v)
  }
}

func BenchmarkSmoothsortOnUnsorted(b *testing.B) {
  v := make([]int, 1000000)
  for i := 0; i < b.N; i++ {
    fillReverse(v)
    smooth.Sort(v)
  }
}

func BenchmarkQuicksortOnSorted(b *testing.B) {
  v := make([]int, 1000000)
  fill(v)
  for i := 0; i < b.N; i++ {
    sort.Ints(v)
  }
}

func BenchmarkSmoothsortOnSorted(b *testing.B) {
  v := make([]int, 1000000)
  fill(v)
  for i := 0; i < b.N; i++ {
    smooth.Sort(v)
  }
}

