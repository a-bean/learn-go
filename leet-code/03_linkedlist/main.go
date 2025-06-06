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
// getEnd 获取从头节点开始的第k个节点
// head: 链表头节点
// k: 要查找的位置
// 返回: 第k个节点，如果链表长度小于k则返回nil
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

// reverse 反转从头节点到stop节点之前的链表
// head: 要反转的链表头节点
// stop: 反转的终止节点（不包含在反转范围内）
func reverse(head *ListNode, stop *ListNode) {
	last := head // last指向当前反转部分的最后一个节点
	for head != stop {
		nextHead := head.Next // 保存下一个要处理的节点
		head.Next = last      // 当前节点指向前一个节点
		last = head           // last向后移动
		head = nextHead       // head指向下一个要处理的节点
	}
}

// reverseKGroup K个一组反转链表
// head: 链表头节点
// k: 每组的大小
// 返回: 反转后的链表头节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	protect := &ListNode{Val: 0, Next: head} // 哨兵节点，用于处理头节点的反转
	last := protect                          // last指向已处理部分的最后一个节点

	for head != nil {
		// 查找当前组的结束节点
		end := getEnd(head, k)
		if end == nil {
			break // 如果剩余节点不足k个，保持原有顺序
		}

		nextGroupHead := end.Next // 保存下一组的起始节点

		// 反转当前组的节点
		reverse(head, nextGroupHead)

		// 将反转后的部分连接到链表中
		last.Next = end           // 前一组的尾节点指向反转后的头节点
		head.Next = nextGroupHead // 反转后的尾节点指向下一组的头节点

		// 更新指针，准备处理下一组
		last = head          // 更新last为当前组的尾节点
		head = nextGroupHead // 移动到下一组的开始位置
	}

	return protect.Next
}

// 141: 环形链表 https://leetcode.cn/problems/linked-list-cycle/
// 解法1: 快慢指针
func hasCycle(head *ListNode) bool {
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head = head.Next
		if fast == head {
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
		pre = newHead // 归位，重头开始 TODO:
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

// 2. 两数相加 https://leetcode.cn/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return dummy.Next
}

// 24. 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/description/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := cur.Next.Next

		first.Next = second.Next
		second.Next = first
		cur.Next = second

		cur = first
	}

	return dummy.Next
}

// 138 . 复制带随机指针的链表 https://leetcode.cn/problems/copy-list-with-random-pointer/description/
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList 复制一个包含随机指针的链表
// 算法分三步：
// 1. 在每个节点后创建其复制节点
// 2. 处理随机指针
// 3. 分离原链表和复制的链表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 第一步：在每个节点后创建其复制节点
	// 例如：1->2->3 变成 1->1'->2->2'->3->3'
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}

	// 第二步：处理随机指针
	// 利用 N 和 N' 的关系：N'.random = N.random.next
	// node是原始节点，node.Next是其复制节点
	// node.Random是原始节点的随机指针指向的节点
	// node.Random.Next就是原始节点随机指针指向节点的复制节点
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}

	// 第三步：分离原链表和复制的链表
	headNew := head.Next // 保存新链表的头节点
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next       // 当前节点的复制节点
		node.Next = node.Next.Next // 恢复原链表的 next 指针
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next // 建立新链表的 next 指针
		}
	}

	return headNew
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
