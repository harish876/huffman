package main

import (
	"container/heap"
)

var (
	INTERNAL_NODE = '*'
	NOT_FOUND     = '!'
)

type HuffmanTree struct {
	Root  *Node
	Codes map[byte]string
}

func NewHuffmanTree() *HuffmanTree {
	return &HuffmanTree{
		Root:  nil,
		Codes: make(map[byte]string, 0),
	}
}

func NewNode(freq int, char byte) *Node {
	return &Node{
		Freq:  freq,
		Char:  char,
		Left:  nil,
		Right: nil,
	}
}

type Node struct {
	Freq  int
	Char  byte
	Left  *Node
	Right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Freq < pq[j].Freq
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	*pq = old[:n-1]
	return item

}

func (hf *HuffmanTree) Encode(word string) {
	hf.Root = encode(hf.Root, word)
	hf.AssignCodes()
}

func encode(root *Node, word string) *Node {
	freq := make(map[byte]int)
	pq := make(PriorityQueue, 0)

	for i := 0; i < len(word); i++ {
		freq[word[i]]++
	}

	for char, f := range freq {
		heap.Push(&pq, NewNode(f, char))
	}

	for len(pq) > 0 {

		if len(pq) == 1 {
			return root
		}

		first := heap.Pop(&pq).(*Node)
		second := heap.Pop(&pq).(*Node)

		newNode := NewNode(first.Freq+second.Freq, byte(INTERNAL_NODE))

		newNode.Left = first
		newNode.Right = second
		root = newNode

		heap.Push(&pq, newNode)
	}

	return root
}

func (hf HuffmanTree) AssignCodes() {
	assignCodes(hf.Root, "", hf.Codes)
}

func assignCodes(root *Node, path string, codes map[byte]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		codes[root.Char] = path
	}

	assignCodes(root.Left, path+"0", codes)
	assignCodes(root.Right, path+"1", codes)
}

func (hf *HuffmanTree) Decode(code string) byte {
	node := decode(hf.Root, code, 0)
	if node != nil {
		return node.Char
	}

	return byte(NOT_FOUND)

}

func decode(root *Node, code string, curr int) *Node {
	if root == nil || curr > len(code) {
		return nil
	}

	if root.Left == nil && root.Right == nil && curr == len(code) {
		return root
	}

	if code[curr] == '0' {
		return decode(root.Left, code, curr+1)
	}

	if code[curr] == '1' {
		return decode(root.Right, code, curr+1)
	}

	return nil

}
