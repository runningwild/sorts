package smooth_test

import (
  . "gospec"
  "gospec"
  "smooth"
  "rand"
  "sort"
  "fmt"
)

const N = 1000

func BasicSpec(c gospec.Context) {
  v := make([]int, N)
  for i := range v {
    v[i] = len(v) - i - 1
  }
  smooth.Ints(v)
  for i := range v {
    c.Expect(v[i], Equals, i)
  }
}

func RepeatedNumbersSpec(c gospec.Context) {
  v := make([]int, N)
  for i := range v {
    v[i] = i % 5
  }
  smooth.Ints(v)
  for i := 1; i < len(v); i++ {
    c.Expect(v[i-1] <= v[i], Equals, true)
  }
}

func ShuffleSpec(c gospec.Context) {
  v1 := make([]int, N)
  v2 := make([]int, N)
  for j := 0; j < 100; j++ {
    for i := range v1 {
      v1[i] = rand.Int()
      v2[i] = v1[i]
    }
    sort.Ints(v1)
    smooth.Ints(v2)
    for i := range v1 {
      c.Expect(v2[i], Equals, v1[i])
    }
  }
}

func ShuffleSpec2(c gospec.Context) {
  v1 := make([]int, 100)
  v2 := make([]int, 100)
  for j := 0; j < N; j++ {
    for i := range v1 {
      v1[i] = rand.Int()
      v2[i] = v1[i]
    }
    sort.Ints(v1)
    smooth.Ints(v2)
    for i := range v1 {
      c.Expect(v2[i], Equals, v1[i])
    }
  }
}

type IntCounter []int
var swap_count,less_count int
func (p IntCounter) Len() int { return len(p) }
func (p IntCounter) Less(i, j int) bool {
  less_count++
  return p[i] < p[j]
}
func (p IntCounter) Swap(i, j int) {
  swap_count++
  p[i], p[j] = p[j], p[i]
}

func CountSpec(c gospec.Context) {
  src := make([]int, N)
  v := make([]int, N)
  for i := range src {
    src[i] = i
  }

  swap_count = 0
  less_count = 0
  copy(v, src)
  sort.Sort(IntCounter(v))
  fmt.Printf("Quicksort on sorted: %d %d\n", swap_count, less_count)

  swap_count = 0
  less_count = 0
  copy(v, src)
  smooth.Sort(IntCounter(v))
  fmt.Printf("Smoothsort on sorted: %d %d\n", swap_count, less_count)

  for i := range src {
    src[i] = len(src) - i - 1
  }
  swap_count = 0
  less_count = 0
  copy(v, src)
  sort.Sort(IntCounter(v))
  fmt.Printf("Quicksort on reversed: %d %d\n", swap_count, less_count)

  swap_count = 0
  less_count = 0
  copy(v, src)
  smooth.Sort(IntCounter(v))
  fmt.Printf("Smoothsort on reversed: %d %d\n", swap_count, less_count)
}
