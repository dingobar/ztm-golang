//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg, input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%s: %s", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	timeParts := strings.Split(input, ":")
	if len(timeParts) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time components", input}
	}
	hour, err := strconv.Atoi(timeParts[0])
	if err != nil {
		return Time{}, &TimeParseError{"Invalid hour", input}
	}
	minute, err := strconv.Atoi(timeParts[1])
	if err != nil {
		return Time{}, &TimeParseError{"Invalid minute", input}
	}
	second, err := strconv.Atoi(timeParts[2])
	if err != nil {
		return Time{}, &TimeParseError{"Invalid second", input}
	}

	return Time{hour, minute, second}, nil

}

func main() {
	fmt.Println(ParseTime("13:15:62"))
	fmt.Println(ParseTime("50:15:62"))
	fmt.Println(ParseTime("1a:15:62"))
	fmt.Println(ParseTime("13:1b:62"))
	fmt.Println(ParseTime("13:11:6c"))
	fmt.Println(ParseTime("1311:6c"))
}
