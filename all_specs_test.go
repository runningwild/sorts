package smooth_test

import (
  "gospec"
  "testing"
)


func TestAllSpecs(t *testing.T) {
  r := gospec.NewRunner()
  r.AddSpec(BasicSpec)
  r.AddSpec(RepeatedNumbersSpec)
  r.AddSpec(ShuffleSpec)
  r.AddSpec(ShuffleSpec2)
  r.AddSpec(CountSpec)
  gospec.MainGoTest(r, t)
}

