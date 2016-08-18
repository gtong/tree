package main

import (
	"fmt"
	"math/rand"
)

/*
   A     3
 B   C   1 3
D E F G  0 1

       A          7
   A       B      3 7
 A   B   C   D    1 3
A B C D E F G H   0 1 width * 2 + gap

              123               14
      123             123       6 13
  123     456     123     456   2 5
123 456 123 456 123 456 123 456 0 1
*/

const (
	width = 3
	format = "%3d"
	gap = 1
	tol = 2
)

type node struct {
	v int
	dl int
	dr int
	l *node
	r *node
}

func (n *node) add(v int) (*node, int) {
	if n == nil {
		return &node{v: v}, 1
	}
	if v < n.v {
		n.l, n.dl = n.l.add(v)
	} else if v > n.v {
		n.r, n.dr = n.r.add(v)
	}
	dl := n.l.depth()
	dr := n.r.depth()
	if dl >= dr + tol {
		n.rr()
	} else if dl <= dr - tol {
		n.rl()
	}
	return n, n.depth()
}

func (n *node) rr() {
	if n == nil || n.l == nil {
		return
	}
	l := n.l
	r := n.r
	v := n.v
	n.v = l.v
	n.l = l.l
	n.r = &node{
		v: v,
		l: l.r,
		dl: l.r.depth(),
		r: r,
		dr: r.depth(),
	}
	n.dl = n.l.depth()
	n.dr = n.r.depth()
}

func (n *node) rl() {
	if n == nil || n.r == nil {
		return
	}
	l := n.l
	r := n.r
	v := n.v
	n.v = r.v
	n.r = r.r
	n.l = &node{
		v: v,
		l: l,
		dl: l.depth(),
		r: r.l,
		dr: r.l.depth(),
	}
	n.dl = n.l.depth()
	n.dr = n.r.depth()
}

func (n *node) addAll(values... int) *node {
	ret := n
	for _,v := range values {
		ret,_ = ret.add(v)
	}
	return ret
}

func (n *node) depth() int {
	if n == nil {
		return 0
	}
	return max(n.dl, n.dr) + 1
}

func (n *node) print() {
	maxDepth := n.depth()
	for printDepth := 1; printDepth <= maxDepth; printDepth++ {
		p, _ := calc(printDepth, maxDepth)
		pad(p)
		n.printDepth(1, printDepth, maxDepth)
		fmt.Println()
	}
	fmt.Println("---")
}

func (n *node) printDepth(depth int, printDepth int, maxDepth int) {
	_, g := calc(depth, maxDepth)
	if n == nil {
		pad(width)
		pad(g)
	} else {
		if depth == printDepth {
			fmt.Printf("%d%d%d", n.dl, n.v, n.dr)
			//fmt.Printf(format, n.v)
			pad(g)
		} else {
			n.l.printDepth(depth + 1, printDepth, maxDepth)
			n.r.printDepth(depth + 1, printDepth, maxDepth)
		}
	}
}

func calc(depth int, maxDepth int) (int, int) {
	if depth == maxDepth {
		return 0, gap
	}
	_, g := calc(depth + 1, maxDepth)
	return g + (width / 2), (g + (width / 2)) * 2 + gap
}

func pad(num int) {
	for i := 0; i < num; i++ {
		fmt.Print(" ")
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	rand.Intn(100)
	var tree *node
	//tree = tree.addAll(5, 3, 7, 2, 4, 6, 8)
	//tree = tree.addAll(1, 2, 3, 4, 5, 6, 7)
	//tree = tree.addAll(1, 2, 3, 4, 7, 6, 5)
	for i := 0; i < 100; i++ {
		tree,_ = tree.add(rand.Intn(10))
	}
	//tree = tree.addAll(5, 4, 6)
	tree.print()
	fmt.Println(tree.depth())

	//tree.rl()
	//tree.print()
}
