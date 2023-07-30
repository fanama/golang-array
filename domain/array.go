package domain

import "github.com/fanama/golang-array/types"

type Array[T any] struct {
	value []T
}

func BuildArray[T any](v []T) Array[T] {
	return Array[T]{value: v}
}

func (array *Array[T]) Get() []T {
	return array.value
}

func (array *Array[T]) Set(v []T) []T {
	array.value = v
	return array.value
}
func (array *Array[T]) Reduce(reducer types.Reducer[T], initialValue T) T {
	acc := initialValue
	for _, v := range array.value {
		acc = reducer(acc, v)
	}
	return acc
}

func (array *Array[T]) Append(element T) *Array[T] {
	array.value = append(array.value, element)
	return array
}

func Retype[T any, U any](array Array[T], mapper types.RetypeMapper[T, U]) *Array[U] {
	newArray := make([]U, len(array.value))
	for i, v := range array.value {
		newArray[i] = mapper(v, i)
	}
	result := BuildArray(newArray)
	return &result
}

func (array *Array[T]) Map(mapper types.Mapper[T]) *Array[T] {
	for i, v := range array.value {
		array.value[i] = mapper(v, i)
	}
	return array
}

func (array *Array[T]) Filter(condition types.Condition[T]) *Array[T] {
	j := 0
	for i, v := range array.value {
		if condition(v, i) {
			if i != j {
				array.value[j] = v
			}
			j++
		}
	}
	array.value = array.value[:j]
	return array
}

func (array *Array[T]) FindIndex(condition types.Condition[T]) int {
	for i, v := range array.value {
		if condition(v, i) {
			return i
		}
	}
	return -1
}

func (array *Array[T]) ReverseArray() *Array[T] {
	for i := len(array.value)/2 - 1; i >= 0; i-- {
		opp := len(array.value) - 1 - i
		array.value[i], array.value[opp] = array.value[opp], array.value[i]
	}
	return array
}
