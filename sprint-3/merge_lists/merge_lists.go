package mergelists

type Node struct {
	Val  int
	Next *Node
}

func MergeLists(list1 *Node, list2 *Node) *Node {

	result := &Node{}
	ptr := result

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			ptr.Next = list1
			list1 = list1.Next
		} else {
			ptr.Next = list2
			list2 = list2.Next
		}
		ptr = ptr.Next
	}

	if list1 == nil {
		ptr.Next = list2
	}

	if list2 == nil {
		ptr.Next = list1
	}

	return result.Next
}
