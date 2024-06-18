package ntest

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
是否有环 hash O(n)/O(n)
*/
func hasCycle(head *ListNode) bool {
	m := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}

/*
*
快慢指针 O(n)/O(1), 有环进入环的就等相遇
*/
func hasCyclePoint(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

/*
*
链表反转
*/
func reverseList(head *ListNode) *ListNode {
	// 前一个节点
	var prev *ListNode
	// 当前节点
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	// nextTemp 就是让curr一轮后，把自己指向下一位
	// 1轮: nil<-1(prev) (curr)2 -> 3 -> 4 -> 5
	// 2轮: nil<-1<-2(prev) (curr)3->4->5
	return prev
}
