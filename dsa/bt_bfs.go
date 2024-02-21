package main

func BFS(head *BinaryNode, needle int) bool {
	q := NewQueue()
	q.Enqueue(head)

	for q.Length > 0 {
		var curr *BinaryNode = q.Deque().(*BinaryNode)
		
		if curr == nil {
			continue
		}

		if curr.Value == needle {
			return true
		}

		q.Enqueue(curr.left)
		q.Enqueue(curr.right)
	}

	return false
}

func BFS2(head *BinaryNode, needle int) bool {
	var q []*BinaryNode
	q = append(q, head)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == nil {
			continue
		}

		if curr.Value == needle {
			return true
		}

		q = append(q, curr.left)
		q = append(q, curr.right)

	}

	return false
}
