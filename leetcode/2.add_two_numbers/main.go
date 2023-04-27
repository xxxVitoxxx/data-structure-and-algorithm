package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	carry, dummy := 0, &ListNode{}
	for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.next {
		if l1 != nil {
			carry += l1.value
			l1 = l1.next
		}

		if l2 != nil {
			carry += l2.value
			l2 = l2.next
		}

		node.next = &ListNode{value: carry % 10}
		carry /= 10
	}
	return dummy.next
}

func main() {
	tests := map[string]struct {
		l1     *ListNode
		l2     *ListNode
		answer *ListNode
	}{
		"example1": {
			l1: &ListNode{
				value: 2,
				next: &ListNode{
					value: 4,
					next: &ListNode{
						value: 3,
					},
				},
			},
			l2: &ListNode{
				value: 5,
				next: &ListNode{
					value: 6,
					next: &ListNode{
						value: 4,
					},
				},
			},
			answer: &ListNode{
				value: 7,
				next: &ListNode{
					value: 0,
					next: &ListNode{
						value: 8,
					},
				},
			},
		},
		"example2": {
			l1: &ListNode{
				value: 0,
			},
			l2: &ListNode{
				value: 0,
			},
			answer: &ListNode{
				value: 0,
			},
		},
		"example3": {
			l1: &ListNode{
				value: 9,
				next: &ListNode{
					value: 9,
					next: &ListNode{
						value: 9,
						next: &ListNode{
							value: 9,
							next: &ListNode{
								value: 9,
								next: &ListNode{
									value: 9,
									next: &ListNode{
										value: 9,
									},
								},
							},
						},
					},
				},
			},
			l2: &ListNode{
				value: 9,
				next: &ListNode{
					value: 9,
					next: &ListNode{
						value: 9,
					},
				},
			},
			answer: &ListNode{
				value: 8,
				next: &ListNode{
					value: 9,
					next: &ListNode{
						value: 9,
						next: &ListNode{
							value: 9,
							next: &ListNode{
								value: 0,
								next: &ListNode{
									value: 0,
									next: &ListNode{
										value: 0,
										next: &ListNode{
											value: 1,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for k, v := range tests {
		fmt.Printf("%s -> ", k)
		result := addTwoNumbers(v.l1, v.l2)
		fmt.Printf("%v \n", equal(result, v.answer))
	}
}

func equal(l1, l2 *ListNode) bool {
	node1, node2 := l1, l2
	for node1 != nil && node2 != nil {
		if node1.value != node2.value {
			return false
		}
		node1, node2 = node1.next, node2.next
	}

	if node1 == nil && node2 == nil {
		return true
	}
	return false
}
