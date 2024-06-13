package main

// 206: https://leetcode.cn/problems/reverse-linked-list/
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

// 25: https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
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

// 141: https://leetcode.cn/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	first := head
	for first != nil && first.Next != nil {
		first = first.Next.Next
		head = head.Next
		if first == head {
			return true
		}
	}
	return false
}

// 21: https://leetcode.cn/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

// 23: https://leetcode.cn/problems/merge-k-sorted-lists/description/
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])

	return mergeTwoLists(left, right)
}

func main() {
	reverseKGroup(&ListNode{
		Val:  1,
		Next: nil,
	}, 1)
}
