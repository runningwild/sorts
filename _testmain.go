package main

import "./_xtest_"
import "testing"
import __os__ "os"
import __regexp__ "regexp"

var tests = []testing.InternalTest{
	{"smooth_test.TestAllSpecs", smooth_test.TestAllSpecs},
}

var benchmarks = []testing.InternalBenchmark{	{"smooth_test.BenchmarkQuicksortOnUnsorted", smooth_test.BenchmarkQuicksortOnUnsorted},
	{"smooth_test.BenchmarkSmoothsortOnUnsorted", smooth_test.BenchmarkSmoothsortOnUnsorted},
	{"smooth_test.BenchmarkQuicksortOnSorted", smooth_test.BenchmarkQuicksortOnSorted},
	{"smooth_test.BenchmarkSmoothsortOnSorted", smooth_test.BenchmarkSmoothsortOnSorted},
}

var matchPat string
var matchRe *__regexp__.Regexp

func matchString(pat, str string) (result bool, err __os__.Error) {
	if matchRe == nil || matchPat != pat {
		matchPat = pat
		matchRe, err = __regexp__.Compile(matchPat)
		if err != nil {
			return
		}
	}
	return matchRe.MatchString(str), nil
}

func main() {
	testing.Main(matchString, tests, benchmarks)
}
