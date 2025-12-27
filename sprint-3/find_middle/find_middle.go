package find_middle

type Node struct {
	Val  int
	Next *Node
}

func FindMiddle(head *Node) *Node {
	result := head
	if result != nil {
		i := 1
		next := head.Next
		for next != nil {
			i++
			if i%2 == 0 {
				result = result.Next
			}
			next = next.Next
		}
	}
	return result
}
