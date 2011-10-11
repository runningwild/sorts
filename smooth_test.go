package smooth_test

import (
  . "gospec"
  "gospec"
  "smooth"
  "rand"
  "sort"
)

func BasicSpec(c gospec.Context) {
  v := make([]int, 100)
  for i := range v {
    v[i] = len(v) - i - 1
  }
  smooth.Sort(v)
  for i := range v {
    c.Expect(v[i], Equals, i)
  }
}

func RepeatedNumbersSpec(c gospec.Context) {
  v := make([]int, 100)
  for i := range v {
    v[i] = i % 5
  }
  smooth.Sort(v)
  for i := 1; i < len(v); i++ {
    c.Expect(v[i-1] <= v[i], Equals, true)
  }
}

func ShuffleSpec(c gospec.Context) {
  v1 := make([]int, 10000)
  v2 := make([]int, 10000)
  for j := 0; j < 10; j++ {
    for i := range v1 {
      v1[i] = rand.Int()
      v2[i] = v1[i]
    }
    sort.Ints(v1)
    smooth.Sort(v2)
    for i := range v1 {
      c.Expect(v2[i], Equals, v1[i])
    }
  }
}
