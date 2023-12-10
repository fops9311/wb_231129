// бинарный поиск
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var input = createHugeRandomSliceInt(50)

func createHugeRandomSliceInt(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(10000)
	}
	return res
}

// узел дерева
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// состояния конечного автомата
const (
	TravelLeft = iota
	TravelRight
	Found
	NotFound
	ReplaceLeft
	ReplaceRight
	ReplaceVal
	DoNothing
)

var ErrNotFound error = errors.New("not found")
var ErrBadTravelFunction error = errors.New("bad travel function. func must return const TravelLeft TravelRight	Found")

// f - функция определяющее следующее состояние
func (t *TreeNode) Travel(node *TreeNode, f func(*TreeNode, *TreeNode) int) (*TreeNode, error) {
	switch f(t, node) {
	case TravelLeft:
		fmt.Println(t, "\t\t<<", node)
		return t.Left.Travel(node, f)
	case TravelRight:
		fmt.Println(t, "\t\t>>", node)
		return t.Right.Travel(node, f)
	case ReplaceLeft:
		fmt.Println(t, "\t\t<<=", node)
		t.Left = node
		return node, nil
	case ReplaceRight:
		fmt.Println(t, "\t\t=>>", node)
		t.Right = node
		return node, nil
	case DoNothing:
		fmt.Println(t, "\t\tidle", node)
		return node, nil
	case Found:
		fmt.Println(t, "\t\tfound", node)
		node = t
		return t, nil
	case NotFound:
		fmt.Println(t, "\t\tnot found", node)
		return nil, ErrNotFound
	default:
		fmt.Println(t, "\t\tdefault", node)
		return nil, ErrBadTravelFunction
	}
}

func main() {
	start := &TreeNode{
		Val: input[0],
	}
	fmt.Println(input)
	//алгоритм добавления в дерево
	put := func(p *TreeNode, n *TreeNode) int {
		if n.Val == p.Val {
			return DoNothing
		}
		if n.Val < p.Val {
			if p.Left != nil {
				return TravelLeft
			}
			return ReplaceLeft
		}
		if p.Right != nil {
			return TravelRight
		}
		return ReplaceRight
	}
	for i := 1; i < len(input); i++ {
		val := input[i]
		start.Travel(&TreeNode{Val: val}, put)
	}

	found := 0
	notfound := 0
	//алгоритм поиска по дереву
	find := func(p *TreeNode, n *TreeNode) int {
		if n.Val == p.Val {
			found++
			return Found
		}
		if n.Val < p.Val {
			if p.Left != nil {
				return TravelLeft
			}
			notfound++
			return NotFound
		}
		if p.Right != nil {
			return TravelRight
		}
		notfound++
		return NotFound
	}

	FindNode := &TreeNode{Val: 150}
	fmt.Println(FindNode.Val, "finding-------------")
	fmt.Println(start.Travel(FindNode, find))
	for _, v := range input {
		FindNode = &TreeNode{Val: v}
		fmt.Println(FindNode.Val, "finding-------------")
		fmt.Println(start.Travel(FindNode, find))
	}
	fmt.Println("total:", len(input), "\tfound:", found, "\tnot found", notfound)
}
