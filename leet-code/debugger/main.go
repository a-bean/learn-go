package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != b {
		if a.Next == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b.Next == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}
func main() {
	getIntersectionNode(nil, nil)
}
