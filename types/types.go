package types

type Reducer[T any] func(acc, val T) T
type Mapper[T any] func(val T, index int) T
type RetypeMapper[T any, U any] func(val T, index int) U
type Condition[T any] func(val T, index int) bool
