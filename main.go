package main

import (
	"fmt"
	"go-datastructures/linkedlist"
	"go-datastructures/model"
)

func main() {
	single := linkedlist.SinglyLinkedList{}
	single.Add(model.Object{Value: "first"})
	single.Add(model.Object{Value: "second"}) // This will become head since we're adding on to the front of the list

	for single.HasNext() {
		fmt.Println(fmt.Sprintf("single linkedlist current value: %s", single.Current.Value))
	}

	_, sfound1 := single.Find(model.Object{Value: "second"})
	fmt.Println(fmt.Sprintf("singly linked list contains [second] %t", sfound1))
	single.Remove(model.Object{Value: "second"})
	_, sfound2 := single.Find(model.Object{Value: "second"})
	fmt.Println(fmt.Sprintf("singly linked list contains [second] after Remove() %t", sfound2))

	double := linkedlist.DoublyLinkedList{}
	double.AddHead(model.Object{Value: "first"})
	double.AddHead(model.Object{Value: "second"})
	double.AddTail(model.Object{Value: "third"})

	for double.HasNext() {
		fmt.Println(fmt.Sprintf("double linkedlist current value: %s", double.Current.Value))
	}

	_, dfound1 := double.Find(model.Object{Value: "second"})
	fmt.Println(fmt.Sprintf("doubly linked list contains [second] %t", dfound1))
	double.Remove(model.Object{Value: "second"})
	_, dfound2 := double.Find(model.Object{Value: "second"})
	fmt.Println(fmt.Sprintf("doubly linked list contains [second] after Remove() %t", dfound2))

	// tail does in fact get added to the tail
	double.AddTail(model.Object{Value: "fourth"})
	for double.HasNext() {
		fmt.Println(fmt.Sprintf("double linkedlist current value: %s", double.Current.Value))
	}
}
