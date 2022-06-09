package timeparse

import "testing"

func TestParseTime(test *testing.T) {
	var wasError bool
	table := []struct {
		time string
		ok   bool
	}{
		{"20:10:10", true},
		{"00:00:00", true},
		{"15", false},
		{"a:10:10", false},
		{"10:a:10", false},
		{"10:10:a", false},
	}
	for _, testcase := range table {
		_, err := ParseTime(testcase.time)
		if err != nil {
			wasError = true
		} else {
			wasError = false
		}

		if wasError == testcase.ok {
			test.Errorf("Got wasError=%v when testcase expected ok=%v for %v", wasError, testcase.ok, testcase.time)
		}
	}
}
