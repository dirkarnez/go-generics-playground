package main

import (
	"fmt"
)

type DB struct {
	IntData []int
	StringData []string
}

// unexport
func where[T any](a []T, predicate func (t T) bool) {
	for _, v := range a {
		if predicate(v) == true{
			fmt.Print(v)
		}
	}
}

func main() {
	db := DB{
		IntData: []int{2, 3, 5, 7, 11, 13},
		StringData: []string{"Hello", "playground"},
	}

	where(db.StringData, func (str string) bool {
		if len(str) > 5 {
			return true
		} else {
			return false
		}
	})

	where(db.IntData, func (value int) bool {
		if value > 5 {
			return true
		} else {
			return false
		}
	})
}