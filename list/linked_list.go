package list

import "github.com/charlienet/go-datastructures/locker"

type linkedList[T any] struct {
	locker.Locker
}

func NewLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{}
}
