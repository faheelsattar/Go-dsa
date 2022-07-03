package main

import "fmt"

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func printLinkedList(head *SinglyLinkedListNode) {
	for {
		if head == nil {
			break
		}
		fmt.Println(head.data)
		head = head.next
	}
}
