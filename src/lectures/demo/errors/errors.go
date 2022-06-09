package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element at %d", index))
	}
	return s.values[index], nil
}

func main() {
	stuff := Stuff{values: []int{1, 2, 3, 4}}
	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value was", value)
	}
}
