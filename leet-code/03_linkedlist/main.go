package main

// https://leetcode.cn/problems/reverse-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}

// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
func getEnd(head *ListNode, k int) *ListNode {
	for head != nil {
		k--
		if k == 0 {
			return head
		}
		head = head.Next
	}
	return nil
}

func reverse(head *ListNode, stop *ListNode) {
	last := head
	for head != stop {
		nextHead := head.Next
		head.Next = last
		last = head
		head = nextHead
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	protect := &ListNode{Val: 0, Next: head}
	last := protect
	for head != nil {
		end := getEnd(head, k)
		if end == nil {
			break
		}

		nextGroupHead := end.Next

		reverse(head, nextGroupHead)

		last.Next = end
		head.Next = nextGroupHead
		last = head
		head = nextGroupHead
	}
	return protect.Next
}

func main() {
}
