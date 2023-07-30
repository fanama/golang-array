package main

import (
	"fmt"

	"github.com/fanama/golang-array/domain"
)

func main() {

	arrayParam := []int{1, 2, 3, 4, 5}

	// init array object
	array := domain.BuildArray[int](arrayParam)
	fmt.Println(array.Get())

	// test Map
	array.Map(func(acc, index int) int {
		return acc * 10
	})

	fmt.Println(array.Get())

	// test reverse array
	array.ReverseArray()
	fmt.Println(array.Get())

	// test append
	array.Append(6)

	fmt.Println(array.Get())

	//test filter
	array.Filter(func(acc, index int) bool {
		return acc%2 == 0
	})
	fmt.Println(array.Get())

	//test find index
	index := array.FindIndex(func(acc, index int) bool {
		return acc < 10
	})
	fmt.Println(index)
	// test reduce
	reduce := array.Reduce(func(acc, val int) int { return acc + val }, 0)

	fmt.Println(reduce)

	// reset array
	arrayParam2 := []int{1, 2, 3, 4, 5}
	array.Set(arrayParam2)
	fmt.Println(array.Get())

}
