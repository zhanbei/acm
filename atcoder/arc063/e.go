package main

import "fmt"

type Node struct {
	next    map[int]bool
	checked bool
	values  map[int]bool
}

var nodes []*Node

// 2018-08-04T00:24:54+0800
// 2018-08-04T00:41:59+0800
// 2018-08-04T01:37:57+0800 # Paused for sleeping.
func main() {
	var N, K int
	fmt.Scanf("%d", &N)
	nodes = make([]*Node, N)
	for i := 0; i < N; i++ {
		nodes[i] = &Node{
			make(map[int]bool),
			false,
			make(map[int]bool),
		}
	}
	a, b := 0, 0
	for i := 0; i < N-1; i++ {
		fmt.Scanf("%d%d", &a, &b)
		nodes[a-1].next[b-1] = true
		nodes[b-1].next[a-1] = true
	}

	// Building the tree.
	buildTree(0)

	fmt.Scanf("%d", &K)
	for i := 0; i < K; i++ {
		fmt.Scanf("%d%d", &a, &b)
		nodes[a-1].checked = true
		nodes[a-1].values[b] = true
	}
	bfs(nodes[a-1])
	fmt.Println("Yes")
}

func buildTree(parent int) {
	for i := range nodes[parent].next {
		delete(nodes[i].next, parent)
		buildTree(i)
	}
}

func bfs(node *Node) {
	for i := range node.next {
		pv := potentialValues(node.values)
		if nodes[i].checked {
			pv := merge(pv, nodes[i].values)
			if len(pv) == 0 {
				fmt.Println("No")
				return
			}
			nodes[i].values = pv
		} else {
			nodes[i].values = pv
		}
		nodes[i].checked = true
	}
	for i := range node.next {
		bfs(nodes[i])
	}
}

func potentialValues(pv map[int]bool) map[int]bool {
	m := make(map[int]bool, len(pv)*2)
	for i := range pv {
		m[i-1] = true
		m[i+1] = true
	}
	return m
}

func merge(pv, v2 map[int]bool) map[int]bool {
	m := make(map[int]bool)
	for i := range pv {
		if v2[i] {
			m[i] = true
		}
	}
	return m
}
