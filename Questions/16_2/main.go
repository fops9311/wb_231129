package main

import (
	"errors"
	"fmt"
)

var testArr []int = []int{2, 4, 2, 5, 4, 3, 2, 1}

var ErrSmallArray = errors.New("small array")

func SliceToNodes(s []int) (start *Node, end *Node, err error) {
	if len(s) < 2 {
		return start, end, ErrSmallArray
	}
	start = &Node{
		V: s[0],
	}
	end = &Node{
		V: s[len(s)-1],
	}
	start.Next = end
	end.Prev = start
	for i := 1; i < len(s)-1; i++ {
		n := &Node{
			V: s[i],
		}
		end.InsertBefore(n)
	}
	return
}

type Node struct {
	V    int
	Next *Node
	Prev *Node
}

func (n *Node) InsertAfter(node *Node) {
	next := n.Next
	n.Next = node
	next.Prev = node
}
func (n *Node) InsertBefore(node *Node) {
	node.Next = n
	node.Prev = n.Prev
	node.FixPointers()
}
func (n *Node) FixPointers() {
	if n.Prev != nil {
		n.Prev.Next = n
	}
	if n.Next != nil {
		n.Next.Prev = n
	}
}
func (n *Node) Travel(f func(*Node) (*Node, error), dest *Node) {
	var err error
	if newn, err := f(n); newn != dest.Next && err == nil {
		newn.Travel(f, dest)
	}
	fmt.Println("dest:", err)
	return
}
func (n *Node) Swap(node *Node) {
	n.Prev, node.Prev = node.Prev, n.Prev
	n.Next, node.Next = node.Next, n.Next
	node.FixPointers()
	n.FixPointers()
}

func (nodes *PNodes) Swap(n1 *Node, n2 *Node) {
	n2.Prev, n1.Prev = n1.Prev, n2.Prev
	n2.Next, n1.Next = n1.Next, n2.Next
	n1.FixPointers()
	n2.FixPointers()
	defer func() {
		if nodes.End == n1 {
			nodes.End = n2
			return
		}
		if nodes.End == n2 {
			nodes.End = n1
			return
		}

	}()
	if nodes.Start == n1 {
		nodes.Start = n2
		return
	}
	if nodes.Start == n2 {
		nodes.Start = n1
		return
	}
}

type PNodes struct {
	Start *Node
	End   *Node
}

func main() {
	start, end, err := SliceToNodes(testArr)
	if err != nil {
		fmt.Println(err)
		return
	}
	test := func(n *Node) (*Node, error) {
		if n == nil {
			return n, fmt.Errorf("end of list")
		}
		fmt.Printf("%p\t%v\n", n, n)
		return n.Next, nil
	} /*
		end.Travel(test, start)*/

	fmt.Println("sort\t---------------------")
	quickSort(start, end)
	fmt.Println("trav\t---------------------")
	start.Travel(test, end)
}

type QuickSortable struct {
	start *Node
	end   *Node
}

func (q *QuickSortable) QuickSort(low *Node, high *Node) *Node {
	q.start = &Node{
		Next: low,
	}
	low.Prev = q.start
	q.end = &Node{
		Prev: high,
	}
	high.Next = q.end

	if low != high {
		pi := q.Partion(low, high)
		if pi == nil {
			return q.start.Next
		}
		//quickSort(low, pi.Prev)
		//quickSort(pi.Next, high)
	}
	return q.start.Next
}
func (q *QuickSortable) Partion(low *Node, high *Node) *Node {
	if high == nil {
		return nil
	}
	if low == nil {
		return nil
	}
	i := &Node{
		Next: low,
	}

	trav := func(n *Node) (*Node, error) {
		if n == nil {
			return n, fmt.Errorf("end of list")
		}
		next := n.Next
		if n.V < high.V {
			i = i.Next
			//i.V, n.V = n.V, i.V
			i.Swap(n)
		}
		if next == nil {
			fmt.Printf("end trav\n")
			return next, nil
		}
		fmt.Printf("%p\t%v\n", next.Prev, next.Prev)
		return next, nil
	}
	low.Travel(trav, high)

	i = i.Next

	i.V, high.V = high.V, i.V

	fmt.Printf("part: %p\t%v\n", i, i)
	return i
}

func quickSort(low *Node, high *Node) {
	if low != high {
		pi := partion(low, high)
		if pi == nil {
			return
		}
		//quickSort(low, pi.Prev)
		//quickSort(pi.Next, high)
	}
}

func partion(low *Node, high *Node) *Node {
	if high == nil {
		return nil
	}
	if low == nil {
		return nil
	}
	i := &Node{
		Next: low,
	}

	trav := func(n *Node) (*Node, error) {
		if n == nil {
			return n, fmt.Errorf("end of list")
		}
		next := n.Next
		if n.V < high.V {
			i = i.Next
			//i.V, n.V = n.V, i.V
			i.Swap(n)
		}
		if next == nil {
			fmt.Printf("end trav\n")
			return next, nil
		}
		fmt.Printf("%p\t%v\n", next.Prev, next.Prev)
		return next, nil
	}
	low.Travel(trav, high)

	i = i.Next

	i.V, high.V = high.V, i.V

	fmt.Printf("part: %p\t%v\n", i, i)
	return i
}
