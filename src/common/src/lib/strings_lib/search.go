package strings_lib

import (
	"regexp"
)

var cachedPatterns = make(map[interface{}]*regexp.Regexp)

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return 0
}

func Get(a []string, x string) string {
	for _, n := range a {
		if x == n {
			return x
		}
	}
	return ""
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ContainsRegexp(a []interface{}, x string) bool {
	for _, n := range a {
		if cachedPatterns[n] == nil {
			cachedPatterns[n] = regexp.MustCompile(n.(string))
		}
		if cachedPatterns[n].MatchString(x) == true {
			return true
		}
	}
	return false
}

func SliceContains(a []string, x []string) bool {
	for _, n := range a {
		for _, m := range x {
			if m == n {
				return true
			}
		}
	}
	return false
}
