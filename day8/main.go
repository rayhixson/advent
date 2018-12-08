package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	ID         int
	ChildCount int
	MetaData   []int
	Children   Nodes
}

// reprint the original input string
func (n Node) String() string {
	s := fmt.Sprintf("%d %d ", n.ChildCount, len(n.MetaData))
	for _, c := range n.Children {
		s += c.String()
	}

	// then print the meta data last
	for i := 0; i < len(n.MetaData); i++ {
		s += fmt.Sprintf("%d ", n.MetaData[i])
	}
	return s
}

/*
not working
func (n Node) Dump() {
	s := fmt.Sprintf("%d - ", n.ID)
	hyphenCount := 1 + len(n.MetaData) + n.Len()

	for i := 0; i < hyphenCount; i++ {
		s += "- "
	}

	fmt.Println(s)

	for _, c := range n.Children {
		c.Dump()
	}
}
*/

func (n Node) CountMetaData() (sum int) {
	for _, e := range n.MetaData {
		sum += e
	}
	for _, c := range n.Children {
		sum += c.CountMetaData()
	}
	return sum
}

func (n Node) Len() int {
	l := 2 + len(n.MetaData)
	for _, c := range n.Children {
		l += c.Len()
	}
	return l
}

type Nodes []*Node

func (ns Nodes) String() string {
	return "nodesss"
}

// consume what is needed to create a node and return the rest
func createNode(count, childCount, metaCount int, data []string) (*Node, []string) {
	n := &Node{
		ID:         count,
		ChildCount: childCount,
		MetaData:   make([]int, metaCount),
		Children:   Nodes{},
	}
	for i := 0; i < n.ChildCount; i++ {
		nextChild := data[0]
		nextMeta := data[1]
		count++
		var kid *Node
		kid, data = createNode(count, num(nextChild), num(nextMeta), data[2:])
		n.Children = append(n.Children, kid)
	}

	// after all the kids are created consume the metadata
	for i := 0; i < metaCount; i++ {
		n.MetaData[i] = num(data[i])
	}
	return n, data[metaCount:]
}

func parseTree(s string) *Node {
	a := strings.Split(s, " ")
	top, remainder := createNode(1, num(a[0]), num(a[1]), a[2:])
	if len(remainder) != 0 {
		panic(fmt.Sprintf("Didn't consume: %s\n", remainder))
	}
	return top
}

func main() {
	tree := parseTree(license)
	sum := tree.CountMetaData()
	fmt.Println("Sum of all the meta data entries: ", sum)
}

func num(s string) int {
	c, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return c
}

const license = `9 11 6 3 4 4 3 5 1 9 0 10 1 7 2 6 3 2 8 4 6 3 3 1 1 2 2 3 1 3 1 1 8 0 9 1 1 8 5 2 5 6 6 4 3 1 3 2 2 1 1 1 1 8 0 6 1 3 5 8 4 3 1 1 2 3 1 3 2 3 3 4 1 4 5 3 4 1 8 0 7 9 1 6 4 2 7 9 1 1 1 1 2 2 1 1 1 5 0 6 3 5 1 6 1 3 2 2 1 3 2 1 9 0 11 1 2 8 9 6 7 7 1 1 3 4 1 2 1 1 2 3 3 1 3 2 2 5 2 3 5 1 5 0 7 6 3 1 3 7 7 3 2 1 2 2 1 1 7 0 9 7 6 4 9 1 7 9 1 5 2 2 1 2 3 1 3 1 9 0 6 1 6 3 2 8 1 3 3 1 3 1 3 1 3 1 1 2 5 3 2 3 4 1 7 0 7 7 4 1 7 3 9 8 3 1 1 1 3 2 2 1 6 0 8 1 4 5 4 4 1 1 3 3 1 2 1 1 2 1 5 0 8 8 7 3 8 4 3 9 1 1 3 1 3 1 4 1 3 1 4 5 4 6 5 3 3 5 1 8 0 11 8 4 3 5 3 1 4 1 8 8 7 3 3 1 1 1 2 3 3 1 7 0 11 3 4 2 6 6 4 4 6 2 5 1 3 1 3 2 1 1 3 1 5 0 9 7 2 4 1 1 2 3 5 9 3 2 1 2 1 3 5 1 1 3 3 6 1 9 0 7 3 7 4 1 3 4 9 1 1 3 1 1 3 2 2 1 1 5 0 6 6 1 8 9 9 7 2 3 2 1 3 1 8 0 6 6 3 4 9 1 4 1 1 3 3 2 3 1 2 4 5 1 4 5 3 3 4 1 9 0 7 3 4 4 5 9 1 1 1 1 2 1 1 2 2 1 1 1 6 0 10 7 1 9 8 8 5 7 8 7 5 1 3 2 3 1 3 1 9 0 9 1 3 2 1 1 6 8 1 1 1 1 2 1 2 3 3 2 1 2 3 1 2 3 6 1 6 0 7 4 1 9 1 9 2 9 3 1 1 1 1 2 1 6 0 9 1 9 7 1 6 1 2 3 5 1 3 3 2 1 1 1 5 0 9 2 9 5 1 1 7 6 6 1 3 3 3 1 2 3 3 2 4 1 1 3 6 1 9 0 9 7 6 1 1 8 3 2 2 1 1 3 2 3 2 2 3 3 3 1 6 0 7 1 7 6 2 2 4 7 2 1 3 1 3 1 1 5 0 11 7 3 5 1 4 3 2 3 8 8 7 3 1 1 2 1 5 4 5 3 1 5 2 5 6 4 3 3 7 1 8 0 8 1 6 8 6 7 9 7 1 1 1 1 2 1 2 2 2 1 8 0 7 1 8 4 1 9 7 1 1 3 1 2 2 1 3 1 1 9 0 8 9 3 7 8 6 1 1 4 2 3 1 3 1 2 2 2 1 3 1 5 4 4 1 2 3 7 1 5 0 11 2 3 4 5 7 4 3 9 1 2 1 3 2 1 3 1 1 7 0 11 9 1 1 7 6 6 6 2 6 1 1 2 2 1 3 1 3 1 1 5 0 7 4 9 2 1 7 7 3 1 2 1 2 1 5 2 5 5 2 5 3 3 5 1 8 0 8 4 8 3 9 4 5 3 1 1 2 3 3 1 3 3 2 1 5 0 7 2 1 8 6 2 4 9 1 2 1 2 1 1 6 0 6 9 1 8 5 8 1 1 2 3 2 1 1 5 1 1 4 1 3 6 1 5 0 11 6 2 1 9 8 9 7 6 2 6 5 3 2 1 1 1 1 8 0 9 1 5 6 2 9 2 8 4 8 2 2 2 2 2 1 1 3 1 9 0 9 6 2 6 8 1 4 4 9 4 1 2 1 2 3 3 3 1 3 5 2 2 2 1 3 4 1 4 5 3 3 5 1 8 0 8 2 1 2 3 1 2 5 8 2 2 2 1 1 1 3 2 1 8 0 6 4 3 9 8 1 9 3 2 2 1 3 2 3 1 1 7 0 7 1 3 1 7 4 7 3 1 1 2 2 2 1 1 4 3 1 2 2 3 4 1 9 0 9 3 1 1 9 5 1 4 6 7 3 2 1 2 1 2 2 1 1 1 6 0 11 2 1 5 4 4 1 3 7 3 5 1 2 1 3 3 1 3 1 9 0 7 6 1 8 8 7 3 4 3 2 2 3 2 1 2 2 3 5 1 4 3 3 6 1 7 0 11 1 9 6 3 4 1 2 7 2 1 3 1 2 1 3 3 1 1 1 9 0 10 7 5 7 6 1 3 6 8 8 1 3 1 3 1 3 3 3 2 2 1 9 0 6 1 6 2 5 4 1 2 1 2 3 2 3 2 3 1 5 2 2 5 1 5 3 4 1 8 0 9 4 5 1 6 2 8 3 7 1 3 3 1 1 3 3 2 1 1 8 0 11 2 1 1 7 3 3 7 9 7 1 2 2 3 2 3 1 1 2 1 1 7 0 10 7 3 2 3 3 3 1 9 2 6 2 2 1 2 1 1 2 5 2 2 5 3 7 1 9 0 6 8 5 1 9 6 5 2 2 1 1 2 1 3 2 1 1 8 0 11 6 3 7 8 2 2 9 1 7 1 5 3 1 3 2 1 3 1 2 1 7 0 7 8 9 1 2 4 8 1 3 2 1 1 1 1 1 3 1 2 3 3 5 3 4 2 2 4 4 3 6 1 5 0 9 9 7 5 1 8 1 2 5 8 1 3 2 2 1 1 7 0 6 1 3 5 3 3 4 2 1 2 1 1 1 1 1 6 0 11 1 7 2 8 3 1 3 2 9 7 6 1 3 1 3 1 2 3 2 5 4 3 2 3 5 1 8 0 6 3 3 4 7 6 1 2 1 2 1 2 1 3 2 1 6 0 6 6 1 6 5 2 6 2 2 1 3 2 1 1 6 0 11 5 3 7 9 7 2 8 1 2 3 3 1 1 2 3 2 3 3 5 2 2 4 3 6 1 5 0 9 2 1 4 1 6 3 4 4 1 1 1 1 3 2 1 9 0 9 4 2 6 1 3 8 1 8 5 2 2 1 2 2 1 3 3 1 1 8 0 11 1 6 8 9 4 4 1 1 6 7 1 2 2 2 3 2 1 1 1 5 1 1 2 4 3 3 6 1 6 0 8 1 9 6 3 7 9 2 1 3 1 3 2 2 2 1 9 0 6 1 8 4 1 1 1 1 3 2 2 1 1 3 2 2 1 6 0 11 1 1 2 7 7 1 4 6 6 8 9 1 3 2 3 1 3 5 2 2 2 3 2 1 6 4 2 5 4 3 4 1 9 0 8 8 9 4 1 1 9 9 3 3 1 3 2 3 3 3 2 2 1 5 0 11 3 7 5 6 1 1 7 3 1 2 6 2 1 1 3 1 1 6 0 6 1 1 4 2 5 7 1 1 1 2 1 1 5 2 3 2 3 5 1 9 0 6 5 1 1 3 3 1 3 1 1 2 2 2 2 1 2 1 5 0 9 4 3 1 8 1 3 5 1 6 2 3 2 1 1 1 5 0 10 8 6 4 9 4 3 9 1 8 9 1 2 1 1 2 1 1 2 4 2 3 4 1 9 0 9 1 9 2 1 1 3 5 6 5 3 3 3 3 2 1 3 2 2 1 6 0 11 1 4 3 2 1 3 6 7 3 9 9 1 1 1 2 3 2 1 5 0 11 8 7 5 3 2 4 6 7 8 5 1 1 2 2 1 3 1 2 2 1 3 6 1 6 0 9 1 8 8 9 1 4 2 2 4 1 1 3 1 2 2 1 7 0 9 1 1 3 1 1 9 7 1 3 1 1 1 2 2 3 2 1 7 0 6 4 9 3 1 7 7 1 1 2 2 2 3 2 4 2 2 4 5 4 3 7 1 6 0 11 9 7 2 4 4 5 1 8 6 7 7 2 3 3 2 1 1 1 6 0 6 7 8 1 6 1 9 3 1 1 3 1 1 1 6 0 8 8 5 2 5 5 9 4 1 2 1 1 2 3 3 4 5 1 3 4 2 1 4 5 1 7 2 6 2 6 3 5 4 3 6 1 5 0 9 3 4 9 5 9 1 7 5 4 1 1 1 3 1 1 8 0 10 3 6 7 4 6 5 1 1 4 1 1 2 1 1 1 3 3 2 1 7 0 7 1 4 9 7 3 5 3 1 3 3 1 1 1 3 1 2 4 1 3 3 3 6 1 8 0 9 9 1 7 1 8 1 2 8 4 3 2 2 2 1 3 3 1 1 6 0 7 7 8 9 1 7 8 6 2 1 1 1 2 2 1 5 0 8 2 2 4 2 4 5 2 1 1 1 2 3 3 1 1 3 3 1 1 3 4 1 8 0 6 6 1 4 2 4 9 1 3 1 2 1 2 1 2 1 5 0 6 1 4 4 1 7 7 2 1 1 3 2 1 5 0 6 8 1 4 5 7 9 2 1 1 2 1 5 2 5 1 3 4 1 7 0 8 1 1 1 1 5 4 7 9 3 1 1 3 1 3 3 1 9 0 10 1 1 2 6 7 2 1 1 4 6 3 2 1 2 1 1 1 3 2 1 7 0 8 3 1 8 7 5 8 4 8 1 2 1 1 3 1 2 4 5 1 2 3 7 1 6 0 6 9 1 5 4 2 4 1 1 1 1 2 2 1 5 0 7 5 1 2 2 5 6 2 1 1 3 2 2 1 6 0 11 7 8 1 4 2 9 1 1 3 3 3 3 1 1 2 1 1 3 4 5 1 1 5 1 3 1 5 4 5 4 3 4 1 6 0 6 3 4 2 1 9 6 1 2 3 2 1 1 1 8 0 6 2 5 1 1 1 4 1 2 2 1 1 2 1 2 1 9 0 6 1 5 2 4 1 1 3 1 1 2 3 2 1 3 2 3 3 1 2 3 4 1 7 0 11 7 7 7 4 7 8 8 1 3 5 1 1 3 1 2 3 2 1 1 8 0 6 9 7 8 6 1 6 1 3 1 3 2 2 1 1 1 8 0 9 9 1 2 5 3 3 9 5 8 2 3 1 1 1 1 3 1 2 5 5 3 3 6 1 9 0 7 1 1 3 5 2 4 8 3 2 1 3 3 1 2 1 1 1 9 0 10 1 9 8 5 5 8 9 4 5 5 1 1 1 3 3 3 1 3 1 1 8 0 9 1 2 5 7 8 5 6 1 6 1 1 3 2 1 2 1 3 4 3 3 1 1 2 3 5 1 8 0 7 9 1 3 8 1 2 4 2 2 2 2 2 1 1 2 1 9 0 8 4 4 9 8 9 2 7 1 3 3 1 1 2 3 1 1 3 1 5 0 6 1 6 5 3 9 5 3 2 3 3 1 5 1 1 4 5 3 5 1 8 0 7 5 3 9 3 3 2 1 3 1 1 3 1 2 2 2 1 5 0 9 4 2 3 7 1 4 8 5 3 1 1 3 2 2 1 6 0 6 9 6 4 7 9 1 1 1 2 3 2 1 3 1 2 4 4 5 1 1 1 5 4 3 4 1 5 0 10 1 4 8 6 4 7 5 3 8 5 2 1 3 2 2 1 7 0 6 3 1 8 1 9 3 2 2 1 2 1 3 2 1 9 0 9 4 7 2 3 1 1 7 7 7 1 1 1 1 1 1 2 2 3 2 3 5 1 3 6 1 8 0 7 4 6 5 1 8 9 2 2 3 2 1 2 1 1 2 1 5 0 11 6 5 9 5 7 8 5 3 1 6 6 1 3 1 1 3 1 5 0 9 4 3 3 4 4 4 1 1 4 2 3 2 1 3 4 1 2 3 5 1 3 7 1 6 0 7 7 2 6 3 1 5 6 1 1 1 1 3 2 1 5 0 8 7 6 2 5 2 7 1 4 3 1 2 1 3 1 9 0 8 7 9 9 7 6 7 1 4 1 3 1 2 3 3 2 1 3 1 3 5 5 2 5 2 3 4 1 5 0 9 7 7 2 3 3 1 5 1 9 2 1 3 1 2 1 7 0 7 2 7 3 7 3 1 2 1 3 3 2 2 3 3 1 7 0 10 8 6 9 9 3 5 4 3 1 5 3 3 1 3 1 3 1 5 4 2 2 3 6 1 7 0 8 5 1 8 6 8 5 2 4 1 1 1 1 2 2 1 1 9 0 9 5 1 1 7 1 8 5 7 6 1 1 3 3 1 1 3 1 1 1 7 0 7 5 6 1 9 1 5 6 2 1 2 1 1 1 1 5 2 3 5 3 3 6 5 4 1 4 4 3 6 1 7 0 6 3 2 7 9 1 5 1 1 2 3 3 1 3 1 8 0 6 1 8 8 1 4 4 2 2 3 3 1 3 2 1 1 5 0 10 8 2 9 1 7 4 1 4 8 5 1 2 3 1 1 1 3 4 3 5 1 3 7 1 6 0 9 4 4 9 6 6 1 5 5 1 1 2 2 1 1 3 1 7 0 9 9 1 2 2 7 3 7 7 8 1 2 3 1 1 1 3 1 5 0 6 2 3 1 4 7 1 1 1 3 1 3 1 5 1 4 5 1 3 3 6 1 7 0 6 5 3 1 2 6 3 1 2 3 2 3 2 1 1 6 0 6 9 1 6 1 2 8 1 3 1 1 2 1 1 6 0 10 1 9 9 9 7 5 3 4 5 4 1 1 2 2 3 2 2 2 5 2 2 1 3 5 1 5 0 7 6 3 4 6 1 3 9 2 1 2 3 1 1 7 0 11 1 1 1 1 7 9 5 3 5 1 9 1 3 1 3 1 3 3 1 6 0 7 1 4 1 3 2 7 5 1 1 2 2 3 1 2 1 3 5 5 3 1 1 4 5 4 3 4 1 8 0 9 8 6 8 2 7 9 1 2 9 3 3 2 3 1 1 3 3 1 8 0 7 7 6 9 2 1 2 3 2 1 2 1 3 3 1 2 1 5 0 8 9 2 8 1 9 5 8 3 2 1 2 2 1 3 5 3 5 3 5 1 6 0 6 3 6 8 5 1 1 1 2 2 2 2 2 1 6 0 8 9 1 3 7 3 3 1 9 1 2 2 3 3 1 1 7 0 9 7 8 7 5 1 4 7 4 8 3 1 1 2 1 1 1 4 5 1 2 4 3 7 1 6 0 10 7 8 6 1 3 3 9 5 6 1 2 1 1 3 3 1 1 8 0 9 8 1 9 1 1 8 3 3 6 3 1 1 3 3 2 1 1 1 7 0 7 1 2 1 1 9 7 2 2 3 1 3 3 3 3 3 2 1 3 3 5 3 3 7 1 7 0 11 4 5 1 6 3 5 6 8 6 8 6 2 2 1 2 1 2 2 1 9 0 10 5 3 1 8 6 6 8 6 3 7 1 3 2 1 2 2 3 3 3 1 6 0 7 9 6 8 9 1 6 3 3 2 1 2 3 1 1 5 1 5 5 4 1 3 4 1 8 0 6 3 3 7 2 7 1 2 1 3 2 3 3 2 3 1 5 0 7 8 9 9 1 1 8 1 2 1 2 1 2 1 8 0 6 5 2 9 1 3 9 2 1 1 3 3 2 2 3 2 5 3 5 3 3 6 1 5 5 3 4 1 8 0 6 1 9 9 1 3 9 2 3 1 3 1 3 3 2 1 6 0 8 9 3 1 2 1 6 5 5 1 3 1 3 1 2 1 9 0 6 1 1 4 6 9 9 2 1 3 1 2 1 1 2 1 1 2 1 2 3 7 1 9 0 10 7 9 1 7 5 1 6 7 9 3 1 2 1 1 3 3 2 2 1 1 5 0 8 3 1 6 1 3 5 1 1 1 3 3 1 3 1 5 0 7 1 2 6 3 9 2 2 3 1 2 1 3 1 4 3 2 1 4 1 3 6 1 9 0 6 1 4 1 5 7 7 3 1 1 2 3 3 2 2 1 1 7 0 10 8 1 2 7 7 5 8 8 1 6 3 2 3 3 1 1 3 1 7 0 7 4 1 3 7 8 3 6 1 2 3 1 1 2 1 2 3 5 2 2 5 3 6 1 6 0 10 7 1 9 1 4 1 1 2 3 8 1 3 3 1 2 3 1 8 0 7 9 4 5 6 5 4 1 1 1 3 3 2 3 3 3 1 8 0 6 9 1 1 1 3 1 2 2 1 3 3 1 3 2 1 2 5 3 1 2 3 4 1 9 0 10 1 5 5 2 1 2 3 9 2 6 3 3 1 2 1 1 3 2 1 1 9 0 11 9 6 4 6 5 8 3 5 1 1 9 3 3 2 1 1 3 2 3 2 1 6 0 11 6 6 2 7 9 1 8 4 5 4 2 1 1 3 2 2 1 3 1 5 1 6 2 2 3 3 1 3 4 7 3 5 4 3 6 1 5 0 6 3 5 7 8 1 3 1 3 1 1 2 1 5 0 11 9 9 3 7 1 6 1 3 1 9 8 1 2 3 2 1 1 8 0 10 5 4 1 1 4 6 9 3 9 6 3 1 3 1 1 1 3 3 1 3 2 3 3 1 3 7 1 8 0 8 4 1 6 6 9 1 1 1 1 3 3 2 3 2 1 2 1 6 0 8 3 8 3 1 4 9 9 1 2 1 3 1 1 1 1 8 0 11 7 7 1 7 4 8 6 6 8 3 6 2 3 1 1 1 2 3 3 4 3 4 4 2 3 2 3 5 1 9 0 11 3 2 1 2 7 4 2 2 1 7 4 3 3 3 1 3 2 1 2 3 1 5 0 11 4 8 3 4 3 6 7 9 4 1 6 2 3 1 2 2 1 6 0 7 6 8 7 7 9 1 8 3 1 1 3 2 3 1 3 1 1 3 3 6 1 8 0 7 1 6 3 7 1 5 3 1 1 2 2 2 3 3 1 1 6 0 11 3 7 9 1 9 1 4 3 2 1 5 1 2 2 2 1 3 1 7 0 6 4 3 3 7 9 1 3 3 2 1 3 1 1 1 1 4 2 2 4 3 6 1 9 0 8 5 6 9 8 4 9 7 1 1 2 1 1 2 1 1 3 3 1 6 0 8 3 3 8 1 6 1 1 8 3 3 1 1 2 1 1 7 0 8 8 4 9 7 7 8 1 8 2 2 1 1 1 3 1 1 1 2 3 1 5 7 1 6 1 5 3 3 6 1 9 0 7 8 1 2 8 1 8 8 1 2 2 2 1 2 1 1 1 1 6 0 11 1 8 8 1 5 5 7 5 2 2 2 2 1 1 3 1 1 1 6 0 10 6 6 4 4 9 1 7 6 2 2 3 1 3 1 2 1 3 3 1 1 5 4 3 7 1 5 0 11 7 6 3 2 1 5 6 1 4 7 4 1 3 2 1 2 1 6 0 7 6 1 5 1 5 7 8 2 1 1 2 3 2 1 9 0 9 2 3 8 1 4 4 9 8 3 2 1 2 2 1 2 3 2 1 3 1 1 1 4 3 2 3 6 1 9 0 7 6 6 1 1 5 3 5 1 1 2 2 3 1 2 3 3 1 5 0 10 4 2 6 5 7 5 7 1 9 4 1 3 1 1 1 1 8 0 7 1 7 1 2 6 6 8 1 1 3 3 1 1 1 1 2 4 1 1 1 1 3 7 1 5 0 8 8 1 2 3 9 3 8 6 2 1 1 1 3 1 9 0 6 9 6 7 2 9 1 3 3 2 1 1 1 2 1 2 1 9 0 8 4 8 1 2 9 6 9 8 2 3 3 2 2 2 1 1 1 3 3 1 1 5 3 4 3 4 1 6 0 7 9 1 9 7 9 4 3 1 3 2 1 1 3 1 6 0 6 9 1 5 9 5 7 1 1 2 1 3 3 1 9 0 9 5 9 3 4 4 5 1 6 3 2 1 1 3 3 1 1 3 2 1 5 5 3 4 4 6 4 4 3 4 1 5 0 10 6 5 1 2 8 6 8 1 5 5 2 3 3 3 1 1 9 0 11 8 9 5 8 6 9 3 6 1 9 1 1 1 1 2 1 3 3 3 2 1 6 0 10 5 8 9 2 9 2 1 4 1 1 3 2 1 1 1 2 4 1 2 5 3 7 1 6 0 10 4 5 7 8 4 1 8 7 4 3 1 3 1 3 2 2 1 5 0 10 1 7 2 3 1 4 2 8 3 7 1 1 1 2 2 1 7 0 7 1 2 2 2 9 4 9 2 2 2 2 1 1 1 5 5 2 2 2 1 5 3 5 1 5 0 10 5 1 4 5 4 9 3 6 7 5 2 3 1 1 2 1 9 0 9 8 6 8 1 9 5 7 8 3 3 2 1 2 3 2 1 3 2 1 6 0 6 5 4 9 1 8 1 2 1 2 1 2 2 4 2 5 2 4 3 6 1 8 0 11 8 9 7 2 6 9 8 4 3 1 5 1 2 1 1 2 1 1 2 1 7 0 10 1 3 4 5 1 7 9 9 5 3 2 1 1 3 2 3 2 1 5 0 10 7 3 9 9 4 3 7 1 6 5 1 1 3 1 1 2 3 5 2 4 2 6 5 2 4 5 3 3 5 1 8 0 10 9 5 9 2 8 1 2 1 2 7 2 1 2 1 1 1 3 2 1 9 0 8 5 8 7 9 3 1 1 4 2 2 1 1 1 3 2 1 3 1 8 0 11 9 2 1 8 9 5 8 9 4 1 3 1 2 1 2 1 1 1 3 2 1 3 1 3 3 7 1 9 0 7 6 3 5 1 5 4 1 1 1 3 1 2 1 1 3 2 1 9 0 10 4 8 9 2 1 3 8 1 5 6 2 2 3 2 1 3 1 3 3 1 8 0 8 3 4 7 6 9 8 1 4 2 2 1 1 3 3 3 3 2 5 5 1 1 4 5 3 7 1 5 0 9 9 3 6 7 4 1 9 5 8 2 1 3 1 3 1 8 0 8 6 1 5 2 1 1 9 8 1 2 1 1 2 2 1 3 1 9 0 10 6 8 4 7 1 7 4 7 6 7 1 3 2 2 1 1 3 1 2 4 5 2 2 1 1 2 3 5 1 6 0 6 7 1 6 1 5 2 3 2 3 3 1 1 1 5 0 7 6 8 2 7 1 2 6 2 2 3 3 1 1 9 0 10 1 7 1 8 1 1 4 6 6 2 2 1 3 3 3 1 1 1 3 1 5 1 4 4 3 5 1 6 0 11 1 4 1 7 3 5 4 9 9 3 6 3 1 2 2 1 1 1 7 0 9 6 9 1 6 8 6 3 2 9 2 1 2 1 2 1 3 1 8 0 11 5 8 1 9 4 5 2 8 7 8 1 3 1 1 1 3 1 1 1 1 2 2 3 4 2 4 3 5 3 3 4 1 5 0 8 4 4 4 7 2 1 2 6 2 2 2 1 1 1 9 0 10 7 6 1 3 1 6 1 1 9 2 3 1 3 2 3 1 3 1 2 1 8 0 9 4 6 1 1 9 3 2 4 6 2 2 2 2 2 1 2 3 2 4 2 5 3 6 1 9 0 6 5 1 8 2 8 6 2 1 1 3 1 3 3 1 2 1 9 0 10 9 8 7 8 6 7 3 4 1 7 2 2 3 1 2 1 3 2 2 1 8 0 11 6 7 4 4 9 1 7 5 1 7 2 1 1 1 1 2 2 1 1 3 3 1 3 5 4 3 4 1 7 0 10 4 3 5 3 7 7 5 4 7 1 2 1 3 1 2 3 3 1 8 0 10 1 3 4 4 5 1 3 9 8 4 2 2 1 1 1 1 1 3 1 6 0 11 8 2 5 1 7 1 8 4 5 6 2 1 1 1 1 1 2 1 2 4 4 3 5 1 5 0 8 1 1 8 6 2 3 7 1 1 1 2 3 1 1 6 0 7 1 1 2 1 9 3 4 3 1 2 1 1 3 1 9 0 11 1 1 8 9 4 8 7 2 6 6 7 1 2 2 3 3 1 2 1 2 3 4 4 1 3 3 5 1 5 0 9 8 6 8 7 1 7 8 8 1 2 1 1 2 2 1 7 0 11 1 6 8 1 1 7 8 1 1 8 9 3 1 3 1 1 1 3 1 8 0 10 9 9 3 1 7 8 2 9 9 5 2 3 2 3 1 1 1 2 2 3 3 5 1 7 4 4 4 5 3 4 1 6 0 6 4 7 4 4 1 8 1 3 1 1 2 3 1 9 0 7 7 5 8 1 1 2 4 3 1 3 1 2 2 2 1 2 1 9 0 9 9 9 2 5 7 1 5 1 9 1 2 1 2 1 1 2 2 2 2 1 4 5 3 5 1 6 0 11 6 3 9 1 6 1 7 1 9 5 8 2 1 3 2 1 1 1 8 0 10 9 7 8 6 9 9 1 8 1 4 3 1 2 3 1 1 1 2 1 6 0 7 1 3 3 1 8 5 4 1 2 1 1 1 2 2 3 2 3 1 3 4 1 8 0 9 5 2 1 4 5 3 2 2 4 2 2 2 3 3 3 1 1 1 6 0 7 1 6 6 4 7 7 7 1 1 1 1 1 1 1 8 0 7 2 8 1 3 8 9 6 1 1 2 3 3 1 3 1 2 3 4 5 3 5 1 8 0 10 1 5 2 7 3 5 8 4 1 7 1 3 2 3 2 1 1 1 1 7 0 8 9 1 2 6 1 6 6 4 1 2 1 1 1 3 3 1 5 0 6 1 5 8 1 6 9 2 1 3 1 2 2 1 3 1 3 5 4 4 5 1 5 5 3 5 1 7 0 6 8 9 3 5 1 2 2 3 2 2 3 1 2 1 9 0 6 9 4 1 7 1 2 2 2 2 1 2 3 1 3 2 1 6 0 7 7 5 3 9 1 3 7 1 1 3 1 3 1 3 4 1 2 4 3 6 1 5 0 10 1 9 2 5 4 8 1 7 6 5 3 1 2 1 3 1 9 0 11 8 2 6 7 7 2 4 1 4 2 5 3 2 1 2 1 3 2 2 1 1 7 0 8 1 8 2 9 6 4 6 9 1 2 3 3 2 1 1 5 4 3 1 1 3 3 5 1 9 0 9 2 9 4 1 1 9 1 9 6 1 2 1 3 1 2 3 2 1 1 9 0 10 1 7 6 8 8 8 8 7 9 9 1 2 1 2 1 3 3 3 2 1 8 0 6 9 8 8 5 7 1 2 1 3 1 3 3 1 1 1 2 5 4 3 3 7 1 8 0 7 3 3 5 8 1 6 1 1 3 3 3 1 2 2 1 1 7 0 6 8 1 1 4 4 7 2 1 3 1 2 2 1 1 5 0 7 9 3 8 1 8 9 1 1 3 2 1 2 3 3 4 1 4 1 2 3 7 1 7 0 10 9 8 5 5 8 4 1 1 2 3 3 3 3 3 2 3 1 1 6 0 6 1 2 3 9 9 3 3 3 2 1 3 1 1 6 0 9 3 2 1 1 2 6 2 1 2 1 2 1 1 1 1 5 2 4 3 4 2 2 2 3 4 6 7 3 4 3 6 2 5 5 3 4 1 5 0 9 8 7 2 6 2 8 1 9 7 2 1 3 3 3 1 8 0 6 9 4 1 1 8 7 1 2 1 2 1 2 2 1 1 7 0 9 3 8 2 3 5 7 7 7 1 3 2 1 3 1 3 2 5 3 4 1 3 5 1 9 0 11 3 9 1 2 1 4 2 2 1 9 5 1 2 3 3 1 2 2 2 1 1 5 0 11 9 2 7 1 5 4 3 7 8 4 5 2 1 3 3 1 1 7 0 9 5 8 5 7 1 4 9 2 6 3 3 3 3 3 1 2 2 2 2 4 5 3 4 1 7 0 6 5 1 7 1 6 9 3 2 1 1 3 2 1 1 5 0 9 5 6 1 5 6 2 5 8 7 3 3 3 2 1 1 5 0 7 7 3 5 5 3 1 9 2 3 1 1 2 2 5 2 4 3 4 1 9 0 8 7 6 8 2 9 9 7 1 3 1 1 2 2 1 1 2 1 1 7 0 9 6 5 9 6 7 1 9 6 2 3 1 1 3 1 1 3 1 5 0 7 1 4 7 7 5 5 5 3 3 2 1 1 1 5 1 1 3 7 1 8 0 10 7 5 2 1 9 4 5 3 6 1 3 2 3 3 1 2 1 3 1 8 0 10 7 1 7 1 2 5 8 3 2 7 2 1 3 1 1 3 3 2 1 8 0 11 6 2 7 7 1 6 8 1 1 1 6 2 1 2 3 1 2 1 1 3 2 2 3 4 2 1 2 6 1 4 2 5 5 3 5 1 7 0 7 7 1 3 1 9 2 3 1 1 2 2 3 2 1 1 6 0 8 1 8 3 5 2 3 3 6 3 3 3 2 1 1 1 5 0 9 6 8 5 8 1 1 7 7 1 1 3 2 2 1 1 4 2 4 3 3 7 1 8 0 7 9 3 7 1 8 1 7 1 1 2 1 2 1 1 2 1 9 0 9 4 1 6 2 9 1 8 5 1 1 3 3 2 1 3 3 2 1 1 9 0 11 8 7 1 7 9 4 9 7 4 5 9 3 3 3 2 2 1 2 3 2 5 3 1 2 1 4 4 3 5 1 7 0 8 1 5 5 9 7 3 1 2 1 3 1 1 2 2 1 1 8 0 6 2 7 7 9 1 2 2 1 2 2 1 1 2 1 1 5 0 6 4 1 1 4 9 4 1 1 3 3 2 1 3 3 1 2 3 6 1 8 0 7 3 6 6 4 2 1 6 3 3 2 1 2 2 1 2 1 6 0 10 4 6 8 5 9 9 4 5 1 7 3 3 2 3 1 3 1 9 0 7 8 7 7 5 1 9 4 1 3 1 2 1 2 1 1 3 4 4 3 3 4 1 3 6 1 8 0 7 4 1 6 5 3 5 5 1 1 1 3 1 2 2 1 1 7 0 10 6 5 5 6 5 7 1 4 3 6 3 3 3 1 1 3 1 1 9 0 11 1 2 3 5 8 2 2 6 1 6 5 1 3 2 1 1 1 1 1 1 2 1 1 1 3 2 6 3 5 2 7 5 5 3 5 1 9 0 6 6 8 5 1 2 1 2 1 1 1 2 3 2 2 1 1 8 0 6 6 3 7 4 7 1 2 1 1 1 1 1 2 2 1 6 0 8 7 8 2 1 4 7 2 7 3 1 1 2 2 3 5 1 3 5 5 3 7 1 6 0 8 3 8 5 2 7 5 1 6 1 3 3 2 3 1 1 9 0 8 1 1 3 2 1 8 8 5 3 1 1 1 2 3 1 2 2 1 7 0 9 3 1 8 5 4 3 4 6 3 3 3 2 2 3 1 1 1 2 1 3 1 3 3 3 6 1 6 0 6 1 4 3 3 9 9 2 1 2 1 1 1 1 6 0 6 1 2 9 5 6 4 1 1 3 1 2 3 1 6 0 9 2 1 8 1 8 5 2 7 7 2 3 2 1 3 1 1 4 4 1 5 1 3 6 1 8 0 7 4 4 5 1 2 3 7 3 1 1 1 2 2 2 2 1 5 0 9 1 1 9 6 3 9 5 7 5 2 1 3 2 1 1 5 0 11 3 3 7 7 1 5 1 2 8 7 1 1 1 2 1 1 1 2 2 5 5 2 3 7 1 7 0 11 6 1 3 4 8 3 8 1 1 8 9 3 3 1 3 2 2 3 1 8 0 11 9 8 5 7 1 8 7 2 6 5 5 1 2 3 1 1 2 1 1 1 8 0 9 3 5 5 7 1 9 5 5 6 2 3 2 1 2 1 2 2 1 2 5 2 3 3 1 1 5 4 3 7 4 5 3 5 1 5 0 11 1 9 9 2 9 2 7 2 2 9 7 3 3 1 3 2 1 9 0 8 4 2 6 1 8 8 2 9 1 1 1 1 2 1 1 2 3 1 8 0 11 3 9 6 3 6 4 9 2 1 7 1 2 3 2 3 1 2 3 1 3 3 2 3 2 3 6 1 9 0 7 9 8 2 9 7 1 3 3 3 2 1 1 1 1 3 1 1 5 0 6 6 2 1 6 1 6 1 1 1 1 1 1 5 0 10 7 5 5 5 7 6 1 1 2 1 2 2 2 1 3 5 1 3 4 3 3 3 6 1 7 0 7 1 4 2 9 9 1 5 1 2 3 2 2 1 3 1 6 0 8 8 6 7 8 4 1 9 3 2 1 2 2 1 1 1 9 0 10 8 1 5 5 6 9 7 5 5 8 2 1 3 2 1 3 1 2 2 3 5 3 1 1 3 3 4 1 8 0 7 4 6 1 7 2 4 6 2 2 1 1 3 3 2 2 1 7 0 11 2 8 1 5 2 5 8 7 9 1 7 3 1 3 3 3 2 1 1 5 0 9 6 5 6 9 1 5 6 9 2 2 1 3 3 1 1 3 4 3 3 6 6 6 3 5 4 3 5 1 7 0 6 7 1 6 1 2 2 1 1 1 1 3 2 2 1 7 0 10 6 2 2 3 8 4 1 4 3 7 1 1 1 2 1 3 3 1 7 0 9 1 5 9 2 7 8 2 9 7 2 3 2 1 3 2 1 2 3 2 1 4 3 5 1 8 0 11 7 9 7 3 9 8 6 5 3 1 3 3 1 1 2 2 1 3 3 1 6 0 6 1 1 5 4 1 1 3 1 1 2 2 1 1 7 0 9 8 3 5 4 6 9 1 6 3 2 3 1 1 2 2 2 5 1 1 2 1 3 6 1 7 0 11 3 1 6 3 3 9 8 7 9 7 3 1 3 3 2 2 3 3 1 9 0 8 1 5 1 4 2 3 3 3 1 2 2 1 1 3 1 1 1 1 6 0 8 3 8 3 1 5 6 3 2 1 1 2 2 3 3 1 3 2 4 4 5 3 5 1 8 0 6 9 6 5 9 9 1 3 2 3 1 3 2 2 3 1 6 0 7 5 5 2 8 1 2 6 1 1 1 3 1 2 1 5 0 6 1 4 1 2 1 4 2 1 1 3 2 4 5 3 1 4 3 6 1 6 0 7 1 9 7 7 5 1 3 2 3 1 1 3 3 1 7 0 10 2 3 2 1 3 9 2 6 4 1 1 1 3 2 1 3 2 1 7 0 7 8 1 1 6 1 5 2 1 2 2 3 2 1 3 2 1 5 2 2 5 2 4 5 4 5 5 3 4 1 5 0 9 5 7 8 1 8 9 1 6 1 3 3 2 2 1 1 9 0 8 3 5 1 5 1 8 9 7 3 1 2 3 1 2 1 3 1 1 8 0 8 1 2 5 5 7 4 5 9 2 1 1 1 2 1 3 1 4 5 1 1 3 5 1 6 0 7 9 5 1 1 6 2 3 1 1 3 1 1 3 1 8 0 9 6 7 4 2 8 1 6 6 7 1 2 2 2 1 2 1 1 1 7 0 8 2 2 5 6 7 5 3 1 2 3 1 2 1 3 2 2 2 2 3 3 3 5 1 5 0 10 3 7 8 3 1 3 7 8 8 7 2 1 2 3 1 1 9 0 8 4 3 1 4 7 3 1 3 3 3 3 1 3 1 3 1 1 1 9 0 11 1 6 6 5 1 8 2 9 1 5 9 3 2 3 2 1 3 2 3 1 2 1 3 5 1 3 5 1 8 0 7 9 4 5 3 1 3 7 1 1 1 2 1 1 3 3 1 9 0 11 6 1 4 8 8 4 9 8 1 9 5 1 3 1 1 1 1 2 3 2 1 7 0 7 9 1 9 1 2 8 4 1 2 1 2 2 3 1 3 4 5 2 1 3 4 1 5 0 10 8 3 8 9 6 5 5 1 2 8 3 1 3 2 1 1 9 0 10 9 9 1 1 4 8 1 6 5 4 2 3 3 2 1 1 3 2 1 1 7 0 9 3 2 2 1 8 1 5 4 2 1 2 3 1 3 1 1 3 1 3 1 5 7 5 5 6 1 3 7 2 4 3 3 6 1 9 0 7 6 8 9 2 1 3 3 2 3 1 3 1 2 3 1 2 1 6 0 8 8 4 3 1 9 9 7 7 3 2 3 1 1 1 1 6 0 9 6 5 1 6 1 5 8 3 3 1 1 1 1 3 1 2 3 2 5 2 2 3 6 1 9 0 9 1 6 7 2 4 1 6 1 9 2 1 2 3 3 2 3 2 2 1 6 0 7 8 9 1 3 8 1 8 3 1 2 2 2 1 1 7 0 10 8 9 3 3 9 2 3 1 8 4 1 3 1 3 3 1 2 3 5 3 1 1 1 3 5 1 9 0 7 2 7 1 8 5 7 9 3 3 3 2 3 1 2 2 1 1 9 0 10 8 4 7 6 1 8 5 9 8 1 2 1 3 3 2 2 3 3 1 1 6 0 11 5 4 3 4 6 1 3 9 6 7 1 3 1 1 1 2 1 5 1 2 2 2 3 5 1 8 0 8 1 5 8 8 2 1 1 8 2 2 1 3 1 1 1 1 1 6 0 7 1 9 8 1 5 4 7 2 3 3 1 3 2 1 9 0 6 6 7 9 7 2 1 3 3 1 2 2 2 2 2 3 3 1 1 3 1 1 5 2 4 3 3 4 1 5 0 11 1 3 6 2 1 8 1 5 2 3 6 3 1 2 1 1 1 8 0 9 1 1 7 9 1 3 8 1 7 2 3 2 2 1 1 3 3 1 7 0 10 4 8 3 9 5 8 1 2 4 5 1 3 1 2 3 2 3 2 5 5 2 3 4 1 5 0 8 7 1 9 8 8 8 3 9 2 2 1 1 3 1 5 0 9 9 4 7 1 8 2 1 1 5 1 2 3 2 3 1 6 0 11 5 8 1 9 8 1 1 5 3 5 2 1 3 1 2 2 2 4 3 4 1 3 6 1 8 0 10 7 1 3 1 5 9 2 5 9 1 2 1 2 1 1 2 3 3 1 9 0 9 5 1 8 5 6 8 9 1 4 1 1 3 3 3 1 3 3 2 1 6 0 9 3 1 8 6 7 1 1 9 9 2 2 2 1 2 1 5 5 5 2 3 5 3 7 1 8 0 6 4 9 4 4 1 3 3 2 2 3 2 3 1 1 1 5 0 6 4 5 1 7 4 7 1 2 3 3 2 1 7 0 8 2 5 7 5 8 6 2 1 1 3 2 1 3 1 3 2 4 4 4 4 3 5 3 2 4 5 5 3 7 1 7 0 10 1 5 3 2 3 6 1 1 1 4 1 1 3 1 1 2 1 1 8 0 6 2 3 3 1 1 9 1 1 2 3 3 3 1 2 1 8 0 6 2 1 8 5 3 2 1 1 3 2 1 2 1 1 5 5 1 2 5 5 4 3 4 1 6 0 6 1 9 7 5 1 4 2 3 3 3 1 1 1 8 0 11 3 3 3 2 2 1 4 9 3 2 4 3 1 3 3 1 1 2 2 1 8 0 7 1 8 3 9 3 4 3 1 3 2 2 1 1 3 2 1 5 3 2 3 6 1 9 0 6 2 1 6 3 3 5 2 1 1 3 2 2 2 2 1 1 8 0 9 2 6 1 5 6 7 4 5 3 3 1 1 1 3 3 3 3 1 5 0 11 7 4 4 4 1 9 9 8 1 2 9 1 1 1 1 2 2 1 3 2 4 3 3 7 1 8 0 9 4 6 7 6 9 1 3 2 9 1 1 1 1 1 3 2 3 1 5 0 6 1 4 7 8 6 2 2 2 1 2 1 1 6 0 11 2 7 2 2 1 5 9 6 9 1 1 1 3 2 3 1 1 3 1 2 3 3 4 4 3 7 1 8 0 6 8 4 8 6 2 1 3 3 1 3 2 3 2 1 1 5 0 9 4 7 9 2 1 2 5 6 1 1 2 1 3 1 1 6 0 6 4 1 9 9 8 2 1 3 3 3 1 1 1 1 5 3 3 1 1 1 1 3 1 1 4 4 3 5 1 7 0 7 4 2 3 1 2 9 8 2 3 3 1 1 3 1 1 7 0 6 1 2 1 5 2 3 1 1 1 3 1 2 2 1 7 0 8 1 9 4 5 1 5 9 4 1 2 2 1 1 3 1 1 1 5 2 1 3 7 1 7 0 8 9 1 7 1 7 3 3 5 1 1 2 3 1 1 1 1 6 0 11 6 1 8 5 3 6 9 7 1 8 5 1 1 1 1 3 3 1 6 0 9 2 3 7 9 7 1 5 3 1 2 2 1 1 3 3 5 4 5 4 1 1 4 3 4 1 7 0 6 3 7 7 5 4 1 1 1 2 3 3 1 2 1 6 0 10 7 1 5 2 3 6 6 5 8 4 1 2 1 2 1 1 1 9 0 10 7 3 8 3 5 5 3 1 1 9 1 1 3 3 3 2 1 1 1 3 1 2 3 3 4 1 8 0 10 7 6 2 8 7 1 6 6 6 4 3 1 2 2 3 3 2 3 1 6 0 7 1 7 2 9 4 4 1 3 1 3 2 1 3 1 7 0 11 2 1 3 5 7 9 1 7 6 7 9 2 3 3 3 1 1 1 3 4 1 2 3 2 3 5 5 3 3 4 1 9 0 7 9 6 8 9 1 4 7 3 3 1 3 1 2 3 1 1 1 9 0 7 8 9 1 1 5 1 8 3 2 2 1 3 1 3 1 2 1 5 0 8 5 8 1 6 5 1 8 1 1 3 1 3 1 1 4 3 4 3 6 1 5 0 6 8 8 5 7 1 1 2 1 2 1 1 1 7 0 9 1 2 2 3 9 2 2 1 3 3 1 1 2 3 1 2 1 5 0 9 6 5 1 1 9 4 4 3 2 1 1 2 3 1 4 5 5 3 3 3 3 4 1 9 0 8 2 5 2 1 5 1 6 6 2 3 1 2 3 2 3 1 2 1 9 0 6 1 5 6 1 1 1 2 1 1 1 1 1 1 1 1 1 9 0 9 2 2 8 5 4 5 7 4 1 3 1 3 1 1 3 2 2 3 2 1 4 1 3 5 1 9 0 11 6 9 5 4 1 8 3 8 3 9 6 1 1 3 3 2 1 1 1 3 1 7 0 10 1 1 2 6 8 1 3 1 6 8 2 1 1 3 1 1 3 1 8 0 11 5 1 9 1 1 1 9 2 5 3 9 2 2 1 1 2 2 1 3 1 1 1 4 5 3 5 1 8 0 8 1 3 2 3 7 1 2 9 3 2 2 2 1 2 1 1 1 6 0 10 2 6 1 1 6 7 1 6 7 2 3 1 1 3 2 1 1 8 0 11 8 7 1 9 2 4 4 6 8 8 2 2 3 1 3 2 3 1 2 1 3 1 5 1 2 5 2 4 4 3 7 1 7 0 8 2 1 8 2 9 5 2 1 1 2 2 2 2 3 1 1 7 0 8 9 7 3 6 5 6 1 4 1 1 1 3 1 2 3 1 7 0 7 3 6 1 9 1 8 7 2 1 1 2 1 2 1 4 4 3 3 1 3 2 3 7 1 8 0 10 5 6 3 1 3 9 1 3 5 4 3 2 2 3 1 1 3 1 1 5 0 6 1 1 5 2 7 2 1 1 2 1 1 1 7 0 8 3 6 1 4 2 8 2 4 1 1 2 1 3 1 1 3 5 2 2 5 1 2 3 4 1 8 0 11 1 6 1 5 7 5 6 1 8 4 5 1 2 2 1 1 1 1 3 1 8 0 8 1 7 9 7 4 3 5 9 2 1 2 3 1 3 1 1 1 8 0 7 4 5 1 8 1 9 4 1 2 3 2 2 1 1 1 4 3 3 1 3 7 1 5 0 7 9 1 1 9 5 1 8 3 3 2 1 2 1 8 0 10 9 8 8 5 7 3 3 1 5 7 2 2 3 3 1 1 1 2 1 8 0 10 4 1 1 4 9 6 9 3 3 9 3 1 3 2 1 1 3 3 1 1 1 1 2 5 1 3 2 5 6 4 3 3 5 1 6 0 6 3 1 6 2 2 1 1 1 3 3 3 3 1 7 0 9 3 8 6 8 8 8 1 6 2 2 3 3 1 2 2 1 1 6 0 9 5 8 4 2 3 4 1 2 3 2 2 2 3 1 1 3 2 2 5 1 3 7 1 7 0 7 9 9 9 5 4 7 1 2 1 2 1 2 2 1 1 5 0 10 3 6 1 1 2 1 2 1 7 6 3 3 1 3 2 1 5 0 8 1 9 7 1 3 3 1 3 1 3 3 1 3 3 1 2 2 3 1 3 3 6 1 6 0 8 4 1 5 4 3 1 8 4 2 3 1 1 3 2 1 7 0 7 3 9 1 9 1 3 2 1 2 3 2 1 1 1 1 8 0 11 8 7 6 1 1 4 6 9 5 9 5 2 2 3 2 3 3 2 1 5 5 1 1 1 5 3 4 1 8 0 7 3 4 1 9 7 3 4 2 3 1 3 2 1 1 2 1 7 0 11 1 1 1 8 8 5 3 4 8 9 5 1 2 1 3 3 2 3 1 9 0 10 7 4 1 1 7 7 9 1 2 2 3 2 1 1 3 3 1 3 1 5 1 5 3 3 2 5 3 4 7 2 5 5 3 7 1 5 0 9 1 1 2 4 3 9 4 1 8 3 1 3 2 3 1 9 0 8 1 4 3 5 3 1 3 7 2 1 1 3 2 3 1 2 2 1 9 0 10 9 6 5 9 7 1 1 3 9 5 3 1 1 2 3 1 2 2 2 1 2 3 2 5 3 4 3 5 1 7 0 9 5 3 9 1 7 7 9 7 1 1 1 2 3 1 2 2 1 6 0 10 8 5 3 4 1 1 1 7 8 6 1 2 2 2 2 3 1 6 0 6 2 1 6 4 8 7 2 2 1 3 1 2 3 5 2 4 3 3 6 1 9 0 11 8 9 1 8 1 4 2 5 3 3 4 2 2 2 3 2 3 2 2 1 1 8 0 8 4 5 6 1 1 9 1 9 3 1 1 1 1 3 2 1 1 7 0 6 1 7 6 5 2 4 2 3 1 3 1 2 1 4 1 4 2 1 4 3 4 1 7 0 6 3 6 6 1 2 1 1 3 1 1 2 2 3 1 7 0 8 3 7 6 3 1 6 1 7 3 1 1 3 1 1 3 1 9 0 9 5 3 3 6 3 3 9 6 1 1 1 1 1 3 2 1 3 1 2 3 1 1 3 6 1 9 0 6 2 7 2 1 9 4 1 1 1 3 1 2 1 1 2 1 9 0 6 1 1 2 8 5 8 1 3 1 1 1 2 1 1 3 1 5 0 9 4 7 7 9 4 7 1 9 2 3 1 2 2 2 1 4 3 5 4 3 5 3 5 1 4 4 4 3 5 1 7 0 8 2 3 1 1 8 5 9 5 1 3 2 1 2 3 1 1 7 0 9 8 5 1 1 7 1 1 5 3 1 1 1 3 3 1 3 1 6 0 7 1 1 1 5 6 5 4 3 3 1 3 3 1 5 2 4 5 1 3 6 1 5 0 11 5 9 1 1 2 2 9 1 7 4 6 1 2 1 1 2 1 6 0 6 9 1 6 5 2 8 2 2 2 3 1 1 1 8 0 7 4 6 4 2 4 3 1 1 1 3 2 3 2 3 1 3 4 2 4 2 1 3 4 1 7 0 10 8 6 6 6 1 1 8 8 5 5 3 1 3 1 3 1 3 1 8 0 8 7 1 7 4 9 2 4 8 2 1 1 1 1 3 1 3 1 7 0 11 1 1 1 8 6 2 3 8 1 8 3 1 1 1 2 3 1 3 2 4 3 3 3 6 1 6 0 6 3 9 5 5 1 5 3 3 1 1 1 1 1 6 0 11 2 1 6 8 5 8 9 8 7 1 8 1 1 1 3 3 3 1 8 0 6 6 9 7 7 1 4 3 1 3 3 2 1 1 1 3 5 3 1 3 2 3 2 1 3 4 5 3 6 1 9 0 6 4 3 7 7 1 2 3 3 1 1 3 3 2 3 3 1 8 0 11 1 4 2 5 9 3 2 1 3 3 5 2 3 1 1 2 2 1 1 1 5 0 9 1 4 9 8 1 7 7 8 5 3 1 2 3 1 4 5 5 2 4 4 3 4 1 8 0 6 6 2 1 9 5 9 3 1 1 1 1 1 3 3 1 5 0 9 2 4 6 8 2 8 1 8 7 1 3 1 3 1 1 9 0 6 7 1 8 7 9 4 1 1 1 3 1 2 1 3 2 1 2 2 1 3 6 1 6 0 11 6 3 8 3 6 9 2 4 6 8 1 1 3 2 1 3 2 1 6 0 7 7 3 4 9 3 1 1 3 2 2 1 3 1 1 7 0 6 2 9 6 1 4 8 2 1 3 1 1 2 2 4 1 1 3 5 3 3 7 1 7 0 6 2 1 1 6 6 4 2 1 1 3 3 2 1 1 9 0 7 5 7 9 1 7 5 1 1 1 2 2 2 1 3 1 2 1 8 0 11 1 3 7 9 2 1 6 4 2 1 4 3 2 3 2 1 2 1 2 4 1 1 4 4 3 4 1 4 2 5 3 4 5 3 4 1 5 0 6 2 8 5 1 9 7 3 1 1 3 2 1 8 0 11 7 8 4 4 3 5 8 8 1 2 7 2 1 1 1 3 1 3 3 1 7 0 6 3 4 2 1 3 5 1 1 3 2 1 3 3 3 3 2 5 3 6 1 5 0 6 1 9 6 3 9 1 1 3 2 1 1 1 7 0 9 9 8 1 8 1 9 8 8 9 2 3 1 3 1 3 3 1 5 0 9 1 8 1 9 1 6 7 6 9 1 1 1 2 2 3 2 2 4 1 5 3 6 1 5 0 7 2 3 7 1 3 1 3 3 3 3 1 3 1 7 0 8 2 9 1 5 6 2 7 6 1 1 2 1 1 3 2 1 6 0 9 1 5 8 4 6 4 6 1 5 3 1 3 2 1 1 3 2 4 3 5 5 3 5 1 7 0 8 6 8 9 4 4 2 1 6 1 1 3 2 3 1 1 1 6 0 7 4 5 5 7 8 1 4 3 2 3 1 3 1 1 9 0 8 3 1 5 6 2 4 9 4 2 3 1 1 2 2 2 1 2 2 1 2 1 3 3 3 4 1 5 4 5 3 6 1 5 0 9 9 2 5 6 1 5 5 7 4 3 2 1 1 2 1 8 0 10 9 9 1 4 2 2 4 2 6 1 1 1 3 3 2 2 3 1 1 9 0 8 5 9 8 6 1 6 6 1 1 3 3 1 1 1 1 2 1 5 4 4 1 3 3 3 4 1 7 0 8 6 8 5 8 3 2 9 1 1 3 1 2 2 1 1 1 6 0 8 1 5 3 4 5 2 1 1 1 1 2 2 1 3 1 5 0 9 1 1 2 9 1 1 4 7 6 1 2 1 1 1 4 1 2 1 3 4 1 9 0 11 4 9 8 4 1 2 3 8 3 8 3 1 1 1 1 2 1 3 1 2 1 6 0 10 3 9 1 7 8 5 1 8 7 6 3 2 1 1 3 1 1 7 0 7 7 1 4 4 7 4 7 2 2 2 3 3 1 2 2 5 1 3 3 4 1 9 0 6 5 1 3 1 9 9 1 3 2 1 3 1 2 3 2 1 6 0 9 6 4 1 5 4 7 5 9 7 3 2 1 3 1 1 1 7 0 10 7 3 3 6 4 1 1 1 7 6 2 1 3 2 1 2 3 5 3 1 3 2 6 2 5 4 5 4 3 4 1 5 0 11 9 9 7 1 5 4 6 5 8 3 4 2 2 1 1 1 1 5 0 11 6 1 9 9 1 3 7 5 4 4 1 2 2 3 1 1 1 9 0 11 9 1 4 6 5 2 1 4 9 2 5 1 2 3 1 1 2 3 2 3 1 2 4 3 3 6 1 9 0 7 6 5 8 5 1 7 4 1 3 3 1 2 3 1 2 1 1 7 0 10 2 8 9 7 3 2 1 2 4 1 3 1 2 3 3 2 3 1 8 0 8 6 2 4 6 1 1 8 5 2 1 3 3 3 1 1 3 3 1 1 5 1 4 3 5 1 6 0 7 2 9 7 1 8 6 8 1 1 1 3 3 1 1 5 0 10 9 4 8 4 4 1 7 8 4 2 1 1 3 2 2 1 8 0 6 5 7 7 3 1 3 2 3 1 3 2 1 1 1 1 3 2 3 5 3 6 1 5 0 9 5 7 2 5 8 2 4 2 1 2 1 1 2 1 1 8 0 8 7 9 9 9 2 1 4 1 1 1 2 3 1 1 1 1 1 5 0 8 1 5 5 2 3 6 1 1 3 2 3 1 1 5 1 5 4 2 4 3 7 1 5 0 11 2 8 6 1 1 8 4 1 3 8 9 3 2 1 2 1 1 8 0 10 4 8 2 7 2 7 8 4 1 9 3 3 3 1 2 2 1 2 1 6 0 8 5 9 9 7 1 1 4 5 1 3 3 1 1 1 1 2 1 4 5 5 3 1 2 2 1 5 5 3 6 1 8 0 11 3 2 7 8 7 6 1 8 8 5 6 1 2 3 1 1 2 3 1 1 9 0 10 3 3 3 2 2 9 2 4 5 1 1 1 3 1 1 1 2 2 3 1 9 0 7 2 3 1 6 8 2 6 1 3 1 3 1 1 2 1 2 1 3 2 3 2 3 3 7 1 8 0 6 3 9 3 9 1 9 1 1 1 3 1 3 2 3 1 9 0 10 4 8 8 8 6 6 9 1 6 3 1 2 3 3 2 3 1 3 3 1 7 0 10 1 2 2 6 3 3 6 5 1 1 2 1 2 3 1 3 1 1 2 5 1 4 3 4 3 4 1 5 0 9 2 8 6 9 2 1 6 6 2 2 2 3 1 2 1 8 0 10 4 1 7 4 6 6 8 7 2 2 1 2 1 3 1 2 2 1 1 9 0 10 1 5 4 5 2 8 6 6 6 4 2 2 1 3 3 3 1 1 1 3 4 3 5 3 6 1 6 0 9 6 4 6 4 1 1 1 5 9 2 1 2 2 3 2 1 6 0 8 2 7 8 1 8 9 2 7 3 1 3 1 3 1 1 6 0 7 3 2 5 3 2 1 1 1 2 1 3 3 3 5 5 3 5 2 2 3 5 1 9 0 6 2 7 7 1 7 4 2 3 2 1 3 1 1 1 1 1 6 0 6 6 5 7 2 5 1 2 3 1 1 1 1 1 5 0 11 2 4 2 5 3 6 1 4 1 2 5 1 3 3 2 1 2 2 5 1 1 5 1 4 5 4 3 5 6 3 5 4 3 4 1 8 0 6 1 7 9 9 6 9 2 3 1 1 3 1 1 3 1 7 0 8 5 7 1 7 3 6 4 7 3 2 1 3 2 3 1 1 5 0 8 1 7 1 4 7 3 9 4 2 1 2 2 2 1 5 4 3 3 6 1 9 0 8 4 8 6 7 2 1 4 8 1 3 3 2 3 3 3 2 1 1 8 0 9 9 1 3 1 1 7 8 8 3 1 1 3 3 2 1 1 3 1 6 0 11 2 5 8 9 4 7 1 4 8 7 1 3 1 3 2 1 1 5 4 5 2 5 3 3 7 1 7 0 6 1 7 8 3 7 7 2 3 1 1 3 1 1 1 7 0 6 6 7 8 1 8 5 2 1 3 2 1 1 3 1 9 0 8 1 5 3 6 1 1 8 1 3 3 3 1 2 2 3 1 1 4 1 3 5 3 1 5 3 6 1 6 0 6 3 4 8 8 1 5 1 1 3 2 3 1 1 8 0 11 6 5 6 1 4 3 5 4 5 4 4 2 1 1 2 3 1 1 3 1 6 0 7 2 9 3 2 3 1 3 1 2 1 1 1 3 4 5 1 1 3 4 3 6 1 9 0 10 8 5 1 1 5 9 7 3 3 7 1 3 2 1 2 1 2 1 3 1 7 0 11 1 9 2 2 8 9 7 2 2 3 4 2 1 1 1 2 2 3 1 8 0 9 1 4 1 4 8 2 8 1 6 3 3 2 1 2 1 3 3 2 3 5 4 1 4 5 1 5 4 4 3 3 4 1 6 0 8 1 9 1 4 4 9 5 3 1 1 1 2 1 3 1 7 0 8 5 1 1 8 5 2 9 4 1 1 3 3 1 3 1 1 5 0 9 1 9 7 3 8 7 8 8 2 1 1 2 2 1 2 5 4 1 3 7 1 8 0 8 8 8 4 5 5 8 5 1 2 3 3 3 2 1 2 3 1 8 0 6 9 1 1 4 8 3 1 2 1 2 3 3 2 3 1 5 0 7 8 1 6 5 8 1 2 2 1 2 1 2 3 1 5 3 2 5 5 3 6 1 9 0 7 3 5 9 4 9 1 1 1 2 3 1 2 1 1 1 3 1 7 0 9 8 1 7 3 3 1 2 5 8 2 3 3 1 3 1 3 1 9 0 6 8 6 1 7 3 4 1 1 2 3 3 2 3 1 3 4 1 4 3 5 3 3 5 1 5 0 11 1 4 8 4 7 2 5 7 1 4 5 1 1 1 3 3 1 7 0 10 5 4 6 5 6 9 1 2 4 1 3 1 1 2 3 3 1 1 7 0 11 7 7 8 1 1 8 3 6 3 2 2 1 3 1 3 3 2 1 3 1 3 1 3 2 4 4 4 3 3 7 1 5 0 8 5 3 8 8 2 1 1 4 3 2 1 1 1 1 9 0 7 6 7 6 1 2 6 4 1 3 2 1 2 1 3 1 3 1 6 0 7 7 6 1 7 4 2 9 2 3 1 2 3 3 5 4 3 3 3 1 1 3 7 1 5 0 10 9 8 9 6 2 2 1 3 1 1 1 1 2 2 1 1 7 0 11 4 7 4 5 4 1 4 1 2 8 7 3 3 2 1 1 2 3 1 7 0 10 1 8 1 1 1 5 1 9 5 4 2 1 1 2 3 1 3 5 1 4 3 5 5 5 3 7 1 9 0 10 7 3 2 1 7 2 3 7 7 1 1 1 2 3 1 2 3 2 1 1 8 0 7 1 1 1 2 2 3 9 2 2 1 1 2 1 3 2 1 6 0 10 2 2 8 1 7 9 6 6 2 8 1 3 1 3 2 1 1 3 1 1 4 1 5 3 5 1 8 0 11 7 9 1 3 9 2 6 1 1 1 3 1 2 1 1 2 1 2 2 1 8 0 9 8 9 2 8 2 5 4 1 7 1 1 2 3 2 2 1 3 1 9 0 7 1 1 6 4 1 4 8 3 2 1 3 1 3 3 3 2 3 2 1 1 4 4 4 4 4 3 3 4 1 9 0 11 9 1 9 3 5 1 5 5 7 5 5 1 2 2 1 3 3 1 2 1 1 9 0 8 6 9 1 7 8 1 9 2 1 2 3 2 1 2 3 3 3 1 9 0 10 8 9 1 4 5 7 1 7 1 7 1 1 3 1 2 1 3 1 2 5 4 1 2 3 5 1 9 0 9 9 7 3 1 1 8 8 6 6 1 2 2 3 2 3 3 1 2 1 6 0 7 7 1 6 5 5 8 6 3 1 1 3 1 2 1 6 0 7 2 2 1 3 6 4 6 1 3 3 3 2 2 1 1 4 1 1 3 5 1 5 0 9 2 9 2 2 3 1 2 9 8 1 1 3 1 1 1 9 0 9 2 5 2 4 7 4 1 5 1 2 1 1 1 1 3 3 1 1 1 5 0 9 9 8 4 4 5 2 6 6 1 1 2 3 3 2 3 1 1 5 3 3 7 1 7 0 11 5 6 7 5 1 4 1 9 2 1 4 1 2 2 2 3 1 1 1 9 0 11 8 6 3 5 4 1 9 8 4 4 6 3 2 2 1 1 2 3 3 3 1 7 0 11 8 5 3 9 8 2 1 4 4 8 2 1 3 3 2 3 1 1 5 2 3 1 1 2 5 4 1 3 5 5 3 5 1 9 0 9 8 5 1 1 2 2 8 1 2 2 1 1 1 2 3 1 2 2 1 8 0 11 3 1 5 8 1 4 6 7 1 2 6 2 1 1 3 1 1 2 1 1 9 0 9 1 5 2 5 5 1 2 8 7 1 3 1 1 3 3 2 3 3 5 4 5 1 1 3 7 1 6 0 7 5 6 1 9 7 1 8 1 1 3 1 3 3 1 5 0 6 4 6 1 8 2 2 1 3 1 3 1 1 5 0 10 9 9 7 7 8 1 1 8 9 4 1 2 3 1 1 5 1 1 4 2 1 3 3 6 1 6 0 8 9 4 2 7 4 5 1 7 1 3 2 3 3 1 1 7 0 8 6 6 1 9 8 7 2 2 3 3 2 3 1 3 1 1 6 0 6 4 6 8 1 6 2 2 2 1 1 1 2 3 4 4 2 2 1 3 4 1 7 0 8 3 2 8 9 1 5 1 5 1 2 3 2 3 3 1 1 8 0 11 4 2 1 3 5 4 5 2 4 1 6 1 2 3 3 3 1 1 1 1 7 0 8 6 8 4 1 9 2 9 5 2 1 3 1 1 1 3 3 2 2 1 3 5 1 5 0 6 8 1 7 9 8 2 1 2 1 1 1 1 9 0 9 2 1 7 8 1 9 2 2 9 3 1 2 1 3 1 2 1 1 1 5 0 9 3 7 5 3 7 1 3 4 5 1 3 3 3 1 5 2 1 1 1 1 7 4 2 4 4 4 3 4 1 5 0 8 1 5 1 2 6 8 2 3 2 1 2 2 1 1 6 0 8 1 4 7 3 7 1 4 7 3 1 1 1 2 1 1 7 0 10 6 4 3 6 5 1 5 9 3 8 1 1 3 2 3 3 1 5 4 2 3 3 4 1 5 0 7 9 1 8 9 2 5 2 1 3 3 3 1 1 9 0 11 9 1 7 2 6 7 9 9 2 4 9 3 1 3 1 3 1 3 2 1 1 5 0 7 1 5 1 5 1 7 3 2 1 1 2 1 1 2 2 5 3 5 1 7 0 10 8 1 7 7 1 7 3 3 7 3 3 3 1 3 1 3 2 1 7 0 7 9 3 4 1 6 4 1 3 1 2 1 3 1 2 1 7 0 8 8 6 1 3 3 4 8 3 3 3 2 2 1 1 1 3 1 3 4 4 3 7 1 9 0 11 1 5 9 5 7 9 2 8 6 1 4 1 3 2 2 3 1 2 1 1 1 5 0 9 5 5 6 8 5 2 3 1 2 1 3 3 1 1 1 7 0 8 1 6 2 5 4 1 2 1 1 2 1 2 1 3 2 5 4 3 3 1 2 4 3 3 1 3 5 4 3 6 2 5 5 3 7 1 6 0 6 7 1 1 5 3 9 1 2 1 2 1 1 1 7 0 7 3 1 1 3 7 3 6 1 3 3 1 3 1 2 1 7 0 11 2 6 2 4 6 8 2 6 6 7 1 1 1 3 3 2 3 1 3 2 3 1 1 3 2 3 4 1 7 0 11 6 5 4 7 7 6 1 8 6 6 9 1 3 2 1 1 1 2 1 5 0 11 4 6 6 3 8 7 9 1 1 2 1 2 1 2 1 2 1 7 0 8 8 1 6 8 1 5 2 1 1 1 1 3 1 2 1 3 5 1 2 3 6 1 8 0 7 6 5 3 1 5 4 3 1 2 1 3 2 3 1 1 1 6 0 7 1 1 5 5 3 2 8 1 2 2 1 3 1 1 8 0 8 1 1 4 8 2 1 9 8 2 3 3 1 1 3 1 3 1 2 5 2 2 2 3 4 1 5 0 9 7 4 5 8 1 2 4 7 4 3 1 3 2 2 1 9 0 6 6 9 1 9 4 6 2 1 2 3 1 2 1 2 3 1 5 0 10 1 6 2 5 6 3 6 5 9 9 1 2 1 3 3 1 1 1 2 3 7 1 5 0 8 9 1 2 5 8 1 1 6 1 1 1 3 3 1 6 0 10 4 6 7 6 9 4 8 1 3 9 3 2 1 3 1 1 1 6 0 6 5 2 6 1 9 8 1 1 1 1 1 3 1 2 1 1 2 3 2 3 6 5 2 4 5 5 3 6 1 7 0 6 1 3 7 5 3 1 3 3 3 2 2 3 1 1 5 0 10 6 1 7 6 9 6 9 4 5 8 2 1 1 2 3 1 7 0 10 2 9 5 9 2 4 7 8 1 5 1 3 1 3 1 3 1 2 2 3 5 3 3 3 4 1 9 0 7 9 9 1 4 9 5 1 1 1 1 1 1 1 2 3 2 1 5 0 8 6 8 4 6 3 7 1 2 1 1 2 3 2 1 7 0 11 5 5 3 7 1 9 5 7 6 4 4 2 3 2 2 2 3 1 5 3 3 2 3 5 1 5 0 10 1 4 1 8 2 2 5 6 4 7 1 3 2 2 3 1 7 0 10 3 4 7 9 1 2 8 2 3 4 1 3 1 2 3 1 3 1 6 0 9 1 3 3 6 7 2 3 3 3 1 1 1 2 2 1 2 3 2 5 3 3 6 1 7 0 6 9 1 6 2 3 1 1 3 1 2 2 3 1 1 5 0 7 5 6 5 5 2 4 1 1 1 1 2 1 1 8 0 8 5 4 3 3 1 4 5 2 1 3 3 1 1 1 1 3 5 5 3 3 2 4 3 6 1 6 0 9 3 9 5 5 1 4 4 8 8 3 3 2 2 3 1 1 9 0 9 2 9 9 1 5 1 1 4 7 1 3 3 3 3 1 2 3 3 1 8 0 6 8 1 1 3 5 7 2 1 1 2 1 1 2 2 1 4 2 2 2 5 6 4 4 5 3 4 5 3 4 1 5 0 9 8 9 7 1 1 9 5 3 8 3 1 1 2 2 1 6 0 11 5 6 1 1 2 1 5 9 8 7 6 3 3 3 3 1 2 1 8 0 9 9 5 4 5 7 4 4 1 8 1 2 3 3 2 2 3 2 2 2 5 2 3 7 1 5 0 6 6 2 3 9 4 1 2 3 3 3 1 1 6 0 8 5 5 7 7 5 1 9 1 2 3 3 1 2 2 1 5 0 10 3 2 8 4 1 7 7 1 6 5 1 2 1 2 1 1 5 2 2 3 1 1 3 4 1 7 0 10 9 5 1 2 5 3 2 9 6 4 1 2 1 1 1 3 1 1 5 0 9 1 1 2 1 1 4 8 8 7 1 2 3 3 1 1 9 0 11 5 7 1 9 2 4 5 3 4 3 8 1 2 2 1 2 1 3 1 1 1 3 3 1 3 7 1 9 0 7 1 6 1 8 8 5 3 3 2 1 1 2 1 2 2 3 1 6 0 8 1 7 4 2 9 1 9 5 1 2 2 3 3 2 1 6 0 11 1 3 3 4 7 4 6 3 6 9 9 3 1 2 3 3 2 3 3 1 2 2 4 1 6 5 1 5 2 4 3 3 7 1 8 0 9 1 1 2 3 3 9 9 7 8 3 1 2 2 3 2 3 1 1 9 0 9 8 9 3 7 2 1 1 5 5 3 3 2 2 1 1 2 3 1 1 7 0 11 8 1 3 6 2 8 1 4 4 2 5 1 3 3 1 2 1 1 3 1 4 2 2 2 2 3 5 1 9 0 9 6 3 7 9 6 1 6 2 6 3 3 2 1 2 1 2 3 1 1 7 0 8 4 1 9 1 9 3 7 1 2 3 2 1 2 1 1 1 7 0 6 8 6 1 8 1 5 3 1 3 2 3 2 1 2 5 1 3 3 3 5 1 7 0 10 6 2 8 6 1 1 8 5 4 9 2 3 1 3 1 1 3 1 9 0 7 3 3 1 8 6 5 1 3 1 2 3 3 1 1 2 2 1 8 0 10 9 7 9 8 4 2 4 1 4 8 1 2 1 1 1 1 2 3 5 1 2 2 1 3 4 1 8 0 10 8 3 1 6 4 5 8 9 8 8 3 1 2 1 3 2 1 2 1 9 0 6 9 1 1 1 1 5 2 3 1 1 3 3 3 2 2 1 8 0 8 1 3 3 9 2 3 3 7 2 1 3 1 3 3 2 3 3 2 1 3 3 4 3 5 5 3 4 1 5 0 11 7 7 5 8 8 1 7 8 4 7 9 3 1 3 3 2 1 6 0 9 8 7 4 9 1 1 1 4 2 3 1 1 1 2 1 1 9 0 6 8 8 2 5 2 1 2 1 1 3 1 2 1 2 1 3 1 4 1 3 4 1 7 0 11 1 1 4 7 3 8 9 7 5 7 5 2 2 3 1 2 3 1 1 7 0 7 5 5 1 6 1 2 1 2 1 3 2 2 1 2 1 5 0 11 1 9 8 6 5 9 4 5 7 5 7 1 1 2 3 1 1 3 4 2 3 5 1 9 0 11 4 5 2 4 1 6 9 7 6 9 1 1 2 1 1 1 1 3 1 1 1 5 0 9 5 2 4 4 9 7 1 9 9 3 3 1 1 2 1 5 0 6 1 9 8 4 1 7 1 3 1 1 2 3 3 3 3 1 3 6 1 8 0 6 5 1 7 4 2 9 3 1 1 3 3 3 3 1 1 9 0 10 1 5 4 2 6 6 7 4 2 1 2 1 1 3 3 3 2 2 1 1 5 0 6 2 9 1 5 5 8 1 2 3 1 1 2 4 1 1 3 4 3 5 1 8 0 7 9 5 1 1 2 1 7 3 2 2 1 2 3 2 3 1 5 0 11 3 1 1 9 5 2 6 5 1 2 7 3 3 3 3 1 1 8 0 6 4 7 7 2 7 1 1 3 2 3 2 2 2 3 1 1 3 4 2 4 1 5 4 3 5 3 3 6 1 6 0 10 1 7 1 6 4 7 9 9 1 4 2 2 1 3 2 2 1 8 0 9 1 3 1 4 5 6 8 6 1 2 1 2 3 2 1 3 2 1 9 0 8 4 2 6 1 8 8 1 2 2 2 2 1 1 2 2 1 1 5 3 3 1 5 2 3 6 1 9 0 10 4 6 2 2 1 7 6 4 8 7 3 1 2 2 3 1 2 1 1 1 5 0 9 1 8 3 4 7 3 8 1 4 1 3 2 1 2 1 8 0 7 9 6 8 1 5 5 1 3 2 3 2 1 2 1 1 4 3 1 1 3 1 3 5 1 6 0 8 9 4 1 9 9 5 4 6 2 2 1 1 2 3 1 6 0 7 4 7 6 6 4 1 1 1 1 2 2 1 2 1 9 0 8 5 6 9 1 8 7 9 2 3 2 2 3 2 2 1 1 1 3 3 5 4 2 3 6 1 7 0 9 1 1 5 6 3 3 4 6 7 3 2 3 2 2 3 1 1 5 0 6 7 7 3 1 5 1 3 2 2 2 1 1 7 0 8 2 6 1 5 4 3 6 4 1 1 1 2 1 3 1 4 4 1 2 2 3 3 7 1 9 0 11 7 5 2 1 1 7 3 6 4 9 6 1 1 1 1 2 1 2 2 2 1 8 0 9 7 2 9 4 1 2 2 4 5 2 3 2 3 3 3 2 1 1 6 0 10 4 1 4 7 7 5 3 3 6 1 1 3 1 1 2 2 1 5 3 3 4 1 3 7 6 2 2 3 6 2 5 5 3 4 1 6 0 9 2 8 5 9 1 6 2 6 5 1 1 3 1 3 1 1 8 0 8 5 7 3 2 1 7 4 5 3 3 1 2 1 3 3 2 1 9 0 6 1 5 3 3 5 6 2 3 1 2 1 2 3 1 1 2 1 2 1 3 7 1 7 0 9 2 1 4 9 1 4 3 1 1 1 3 2 3 1 3 1 1 6 0 7 5 9 1 2 1 2 4 1 1 2 1 1 1 1 8 0 7 1 4 6 9 3 5 9 1 2 2 1 1 3 3 2 3 3 5 5 3 2 2 3 5 1 7 0 10 2 9 4 6 2 4 1 2 5 8 1 1 1 3 1 3 2 1 7 0 7 9 9 1 1 2 5 3 2 1 3 1 2 1 1 1 9 0 6 5 3 1 3 3 2 2 2 2 1 3 1 2 1 3 2 5 2 2 1 3 6 1 6 0 11 6 6 6 4 2 1 5 2 1 4 3 1 1 2 2 1 3 1 5 0 9 7 4 1 1 5 4 5 5 1 1 3 3 1 3 1 5 0 10 3 9 3 3 9 1 5 4 3 8 2 1 2 3 2 1 4 4 1 2 3 3 6 1 8 0 11 7 5 8 2 5 2 1 8 6 2 8 1 3 3 2 3 3 3 2 1 6 0 8 1 2 1 7 8 3 7 3 1 2 1 1 3 1 1 5 0 11 2 1 8 1 4 2 1 5 8 9 6 1 3 1 1 2 2 5 2 5 4 3 2 3 7 7 1 5 4 3 5 1 6 0 10 4 8 2 5 9 6 5 1 4 6 3 1 3 1 3 2 1 5 0 7 3 2 1 6 9 4 1 2 2 1 2 1 1 7 0 8 8 6 8 5 4 1 2 2 1 3 3 1 3 3 2 4 5 2 3 5 3 4 1 5 0 8 8 7 3 1 3 4 4 6 3 3 2 2 1 1 8 0 6 9 1 4 1 8 8 2 1 1 1 3 2 2 1 1 9 0 9 4 4 2 3 1 8 3 8 9 3 1 3 2 2 1 2 1 3 2 2 3 3 3 4 1 9 0 10 8 1 2 9 2 2 5 9 2 1 1 1 1 2 3 1 1 2 2 1 9 0 6 7 1 6 1 3 7 2 1 2 2 1 1 1 2 1 1 5 0 10 6 9 6 9 8 5 1 8 1 4 1 3 1 1 2 2 3 5 3 3 5 1 7 0 8 6 1 7 8 6 8 3 2 2 3 1 1 1 2 1 1 7 0 9 1 7 9 7 4 2 7 8 2 3 3 2 1 3 3 3 1 8 0 7 7 3 1 3 6 7 2 3 1 1 3 1 1 2 1 2 3 4 1 3 3 6 1 9 0 10 4 1 2 9 7 6 9 3 5 1 1 1 3 1 2 1 1 3 3 1 6 0 7 9 9 9 1 6 7 1 1 3 3 3 2 1 1 5 0 11 7 2 9 2 9 2 7 1 9 5 9 1 2 2 1 1 3 4 3 3 5 2 3 4 6 1 5 3 3 4 1 5 0 10 1 2 9 2 8 6 4 4 5 8 2 2 1 1 1 1 5 0 7 1 9 4 2 6 8 5 1 3 3 1 1 1 8 0 10 1 5 5 3 7 7 8 3 1 1 3 2 1 3 2 1 1 1 2 4 5 5 3 6 1 7 0 6 1 4 3 1 2 8 2 3 3 1 1 1 3 1 5 0 8 1 6 9 4 9 3 7 2 2 2 1 3 1 1 9 0 7 5 6 1 4 4 9 3 3 2 1 1 2 1 2 2 1 5 5 2 1 4 4 3 4 1 9 0 7 3 5 1 3 9 8 7 2 1 2 3 2 1 1 3 2 1 7 0 6 1 6 6 5 5 7 2 1 3 3 2 3 1 1 5 0 9 1 9 5 9 7 7 8 9 4 2 3 1 2 3 3 2 1 3 3 5 1 5 0 9 9 4 4 1 9 8 9 5 3 2 2 1 1 1 1 7 0 10 8 8 4 4 3 1 6 3 2 1 1 2 2 1 3 2 2 1 7 0 10 7 5 1 5 3 4 5 3 2 8 1 3 3 3 3 1 1 5 1 1 5 3 3 6 1 7 0 6 1 8 8 1 2 2 3 2 1 2 1 3 2 1 9 0 8 1 2 4 4 5 8 2 4 1 1 1 1 1 3 2 1 2 1 7 0 9 6 7 2 3 4 2 1 5 5 3 1 2 1 3 2 1 5 1 1 1 1 2 2 4 1 5 3 3 7 1 7 0 6 9 1 8 8 3 9 1 1 1 2 1 2 1 1 9 0 10 8 9 3 2 2 1 3 5 1 7 3 1 3 2 2 1 1 1 1 1 6 0 8 2 3 4 4 9 8 1 7 1 1 2 2 1 1 1 5 1 2 2 2 5 3 4 1 7 0 6 3 2 1 7 3 8 3 1 2 1 2 1 2 1 5 0 7 1 3 8 5 1 8 1 3 2 2 1 2 1 9 0 7 1 2 9 9 1 3 1 2 3 1 1 2 1 3 2 2 1 3 2 4 3 4 1 8 0 9 7 9 1 8 7 5 2 7 9 3 1 2 2 1 2 2 1 1 6 0 8 5 7 5 1 1 3 2 8 1 3 1 1 1 3 1 6 0 8 5 1 3 2 1 8 2 9 2 1 3 2 3 3 1 2 5 2 3 7 1 8 0 10 6 6 5 6 1 6 5 1 7 6 1 3 1 1 1 1 3 3 1 7 0 9 4 1 9 1 1 9 4 8 6 2 2 1 2 2 3 1 1 9 0 11 1 9 5 6 5 4 5 2 2 8 6 1 2 2 1 1 3 1 2 1 5 4 2 3 1 4 3 3 4 1 6 0 6 2 4 4 1 3 8 1 3 1 3 3 1 1 9 0 10 1 7 7 1 2 8 7 3 4 4 3 2 2 1 1 2 1 2 1 1 7 0 10 1 3 7 6 6 9 7 5 1 3 3 1 3 3 1 2 3 2 1 1 5 2 7 2 5 5 3 6 1 8 0 9 5 4 5 7 1 7 4 4 8 1 3 2 2 1 1 1 1 1 5 0 10 3 3 6 1 3 8 8 6 8 9 1 3 2 3 1 1 9 0 8 2 6 5 8 7 1 1 1 3 1 2 3 3 1 3 1 2 4 2 1 2 2 3 3 5 1 7 0 8 7 1 2 6 4 9 6 6 1 1 3 1 3 1 1 1 9 0 9 1 9 8 2 8 2 1 9 1 2 2 1 3 3 2 3 2 1 1 6 0 7 2 5 5 1 8 1 1 1 1 1 2 2 2 1 1 2 5 2 3 6 1 8 0 11 2 1 2 8 8 3 6 7 9 4 4 3 3 1 3 2 2 1 1 1 6 0 7 9 5 1 1 3 6 2 1 2 2 2 1 3 1 5 0 8 1 6 6 9 1 7 1 5 1 1 3 2 3 3 2 2 4 3 5 3 7 1 9 0 10 9 9 6 4 3 5 1 9 8 2 3 1 1 2 2 3 1 3 1 1 9 0 8 2 9 1 1 5 7 7 7 1 1 2 1 1 1 1 1 2 1 9 0 9 3 1 3 2 9 8 5 3 8 1 3 3 1 1 1 2 1 2 2 2 3 5 2 5 2 3 5 1 5 0 9 7 7 9 6 1 8 8 4 6 3 2 3 3 1 1 8 0 11 1 7 8 8 4 7 8 8 1 1 1 3 2 1 3 3 2 1 2 1 8 0 9 4 2 1 8 4 7 4 5 1 2 1 3 2 1 1 1 3 1 1 4 5 2 7 5 3 6 2 4 5 3 7 1 5 0 10 1 5 5 8 9 3 9 6 1 5 1 2 1 3 1 1 9 0 10 8 7 6 6 9 5 3 7 4 1 3 1 1 2 3 2 2 1 3 1 6 0 9 9 7 1 7 2 5 6 1 6 1 2 3 3 1 3 3 4 2 5 4 2 5 3 5 1 8 0 8 8 7 2 2 1 7 9 8 3 2 1 2 3 3 3 3 1 7 0 10 1 7 1 7 6 7 8 1 8 5 1 2 1 1 3 3 1 1 6 0 6 7 1 7 1 3 5 2 2 3 1 1 2 1 4 5 2 2 3 6 1 7 0 8 6 1 8 2 3 6 9 6 1 3 3 1 1 1 2 1 9 0 11 4 4 3 3 3 1 6 6 4 1 7 1 3 2 3 2 3 1 1 1 1 6 0 9 4 6 2 9 8 1 1 1 6 1 2 2 3 2 1 3 3 3 1 2 1 3 6 1 9 0 7 9 2 3 5 9 8 1 3 1 1 2 1 1 2 2 1 1 6 0 7 1 2 3 5 4 2 4 1 3 1 2 1 2 1 7 0 6 5 1 5 2 3 9 3 2 3 1 3 3 1 4 5 2 2 3 1 1 2 1 3 5 7 4 8 7 10 1 9 2 7 9 9 11 5`
