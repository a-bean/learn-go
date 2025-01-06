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

// 142: https://leetcode.cn/problems/linked-list-cycle-ii/description/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	isCycle, slow := hasCycle142(head)
	if !isCycle {
		return nil
	}
	fast := head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func hasCycle142(head *ListNode) (bool, *ListNode) {
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true, slow
		}
	}
	return false, nil
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

// 86: 分隔链表 https://leetcode.cn/problems/partition-list/description/
func partition(head *ListNode, x int) *ListNode {
	// 构造 2 个链表，一个链表专门存储比 x 小的结点，另一个专门存储比 x 大的结点
	beforeHead := &ListNode{Val: 0, Next: nil}
	before := beforeHead
	afterHead := &ListNode{Val: 0, Next: nil}
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}

// 92: 反转链表 II https://leetcode.cn/problems/reverse-linked-list-ii/description/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}

func main() {
	reverseKGroup(&ListNode{
		Val:  1,
		Next: nil,
	}, 1)
}
