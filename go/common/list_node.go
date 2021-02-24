package common

type ListNode struct {
	Val  int
	Next *ListNode
}

var TestListNode = &ListNode{
	Val: 5,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	},
}
