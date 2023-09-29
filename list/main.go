package main

import (
	"container/list"
	"fmt"
)

func main() {

	var intList = list.New()

	intList.PushBack(11)
	intList.PushBack(12)
	intList.PushBack(23)
	intList.PushBack(55)
	intList.PushBack(1)
	fmt.Println("length is", intList.Len())
	intList.PushBackList(intList)
	fmt.Println("length is", intList.Len())
	intList.PushFront(1234)
	fmt.Println(intList.Back().Next())
	intList.Remove(intList.Back())

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

}
