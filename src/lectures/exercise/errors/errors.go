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

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// - Hour, minute, and second integer components
type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg, input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func (t *Time) Validate() error {
	if t.hour > 23 || t.hour < 0 {
		return &TimeParseError{"hour out of range: 0 <= hour <= 24", fmt.Sprintln(t.hour)}
	} else if t.minute > 59 || t.minute < 0 {
		return &TimeParseError{"minute out of range: 0 <= minute <= 60", fmt.Sprintln(t.minute)}
	} else if t.second > 59 || t.second < 0 {
		return &TimeParseError{"second out of range: 0 <= second <= 60", fmt.Sprintln(t.second)}
	} else {
		return nil
	}
}

// * Example time string: 14:07:33
func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")

	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time components", input}
	}

	var tempTime []int

	for _, c := range components {
		t, err := strconv.Atoi(c)
		if err != nil {
			return Time{}, &TimeParseError{"Failed to convert input to int", input}
		}
		tempTime = append(tempTime, t)
	}

	time := Time{hour: tempTime[0], minute: tempTime[1], second: tempTime[2]}

	if err := time.Validate(); err != nil {
		return Time{}, err
	}

	return time, nil
}

func main() {
	myTime := "24:55:41"

	t, err := ParseTime(myTime)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("time: %v", t)
}
