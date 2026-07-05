package main

import (
	"iter"
	"strconv"
	"strings"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

func (l *ListNode) Fill(vals []int) {
	for i, v := range vals {
		l.Val = v

		if i < len(vals)-1 {
			l.Next = new(ListNode)
			l = l.Next
		}
	}
}

func (l *ListNode) Iter() iter.Seq[int] {
	return func(yield func(int) bool) {
		for head := l; head != nil; head = head.Next {
			if !yield(head.Val) {
				return
			}
		}
	}
}

func (l *ListNode) String() string {
	var sb strings.Builder

	sb.WriteString("[ ")

	for v := range l.Iter() {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString(" ")
	}

	sb.WriteByte(']')

	return sb.String()
}

func isValid(l *ListNode) bool {
	return l != nil
}

func zeroIfNil(l *ListNode) int {
	if l == nil {
		return 0
	}

	return l.Val
}

type pairs struct {
	n1, n2 int
	isLast bool
}

func isLast(l *ListNode) bool {
	return l != nil && l.Next == nil
}

func nodePairs(l1, l2 *ListNode) iter.Seq[pairs] {
	return func(yield func(pairs) bool) {
		for isValid(l1) || isValid(l2) {
			n1 := zeroIfNil(l1)
			n2 := zeroIfNil(l2)
			p := pairs{
				n1:     n1,
				n2:     n2,
				isLast: isLast(l1) && isLast(l2),
			}

			if !yield(p) {
				return
			}

			if isValid(l1) {
				l1 = l1.Next
			}

			if isValid(l2) {
				l2 = l2.Next
			}
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	root.Next = new(ListNode)
	head := root.Next

	var car int
	for p := range nodePairs(l1, l2) {
		sum := (p.n1 + p.n2 + car)

		head.Val = sum % 10
		car = sum / 10

		head.Next = new(ListNode)
		head = head.Next
	}

	if car != 0 {
		head.Val = car
	}

	return root.Next
}
