package main

import (
	"fmt"
	"sort"
)

// 206: 反转链表 https://leetcode.cn/problems/reverse-linked-list/
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

// 25: K 个一组翻转链表 https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
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

// 141: 环形链表 https://leetcode.cn/problems/linked-list-cycle/
// 解法1: 快慢指针
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

// 解法2: 哈希表
func hasCycle2(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 142 环形链表 II : https://leetcode.cn/problems/linked-list-cycle-ii/description/
// 解法1: 快慢指针
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

// 解法2: 哈希表
// 一个非常直观的思路是：我们遍历链表中的每个节点，并将它记录下来；一旦遇到了此前遍历过的节点，就可以判定链表中存在环。借助哈希表可以很方便地实现。
func detectCycle1(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 21: 合并两个有序链表 https://leetcode.cn/problems/merge-two-sorted-lists/description/
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

// 23: 合并 K 个升序链表 https://leetcode.cn/problems/merge-k-sorted-lists/description/
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

// 147. 对链表进行插入排序 https://leetcode.cn/problems/insertion-sort-list/description/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: nil} // 这里初始化不要直接指向 head，为了下面循环可以统一处理
	cur, pre := head, newHead
	for cur != nil {
		next := cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		pre = newHead // 归位，重头开始
		cur = next
	}
	return newHead.Next
}

// 148. 排序链表 https://leetcode.cn/problems/sort-list/description/

func sortList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	vec := make([]int, 0)
	cur := head
	for cur != nil {
		vec = append(vec, cur.Val)
		cur = cur.Next
	}
	sort.Ints(vec)
	cur = head
	cnt := 0
	for cur != nil {
		cur.Val = vec[cnt]
		cur = cur.Next
		cnt++
	}
	return head
}

func sortList(head *ListNode) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length <= 1 {
		return head
	}

	middleNode := middleNode(head)
	cur = middleNode.Next
	middleNode.Next = nil
	middleNode = cur

	left := sortList(head)
	right := sortList(middleNode)
	return mergeTwoLists1(left, right)
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

// 160. 相交链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/description/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	for a != b {

		// 当 a 到达链表 A 的末尾时，将 a 指向链表 B 的头部
		// 当 b 到达链表 B 的末尾时，将 b 指向链表 A 的头部
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}

		fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}

// 876. 链表的中间结点 https://leetcode.cn/problems/middle-of-the-linked-list/description/
func middleNode1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	reverseKGroup(&ListNode{
		Val:  1,
		Next: nil,
	}, 1)

	sortList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	})

	detectCycle(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	})

}
