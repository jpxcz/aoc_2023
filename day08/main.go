package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(path + "/input.txt")
	if err != nil {
		panic(err)
	}

	c := strings.Split(string(content), "\n\n")
	fmt.Println(c[0])

	// partOne(c[0], c[1])
	partTwo(c[0], c[1])
}

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func NewNode(value string) *Node {
	return &Node{Value: value}
}

func (n *Node) AddLeft(node *Node) {
	n.Left = node
}

func (n *Node) AddRight(node *Node) {
	n.Right = node
}

var nodes = make(map[string]*Node, 0)

func partOne(directions string, lines string) {
	text := strings.Split(lines, "\n")
	var initialNode *Node
	for _, t := range text {
		nodeName := t[:3]
		node := NewNode(nodeName)
		if nodeName == "AAA" {
			initialNode = node
		}
		nodes[nodeName] = node
	}

	for _, t := range text {
		parent := t[:3]
		childs := strings.Split(t[7:len(t)-1], ", ")

		parentNode := nodes[parent]
		childLeftNode := nodes[childs[0]]
		childRightNode := nodes[childs[1]]

		parentNode.AddLeft(childLeftNode)
		parentNode.AddRight(childRightNode)
	}

	steps := 0
	for true {
		n, step := walkLoop(initialNode, directions, steps)
		steps = step
		if n == nil {
			break
		}

		initialNode = n
	}

	fmt.Printf("totalGiven steps=%d\n", steps)
}

func partTwo(directions string, lines string) {
	text := strings.Split(lines, "\n")
	var initialNodes []*Node
	var cycles [][]int

	for _, t := range text {
		nodeName := t[:3]
		node := NewNode(nodeName)
		if nodeName[2] == 'A' {
			initialNodes = append(initialNodes, node)
		}
		nodes[nodeName] = node
	}

	for _, t := range text {
		parent := t[:3]
		childs := strings.Split(t[7:len(t)-1], ", ")
		parentNode := nodes[parent]
		childLeftNode := nodes[childs[0]]
		childRightNode := nodes[childs[1]]
		parentNode.AddLeft(childLeftNode)
		parentNode.AddRight(childRightNode)
	}

	for _, node := range initialNodes {
		steps := 0
		initialZFoundSteps := 0
		var cycle []int
		for true {
			zFound := false
			for steps == 0 || !zFound {
				node, steps, zFound = walkLoopNodes(node, directions, steps)
			}
			cycle = append(cycle, steps)
			if initialZFoundSteps == 0 {
				initialZFoundSteps = steps
				steps = 0
			} else if steps == initialZFoundSteps {
				break
			}
		}

		cycles = append(cycles, cycle)
	}

	var nums []int
	for _, c := range cycles {
		nums = append(nums, c[0])
	}

	lcm := 1

	for _, n := range nums {
		lcm = LCM(lcm, n)
	}

	fmt.Printf("part2 steps=%+v\n", lcm)
}

func walkLoopNodes(node *Node, rute string, step int) (*Node, int, bool) {
	direction := getNewDirection(rute, step)
	if direction == 'L' {
		if node.Left.Value[2] == 'Z' {
			return node.Left, step + 1, true
		}
		return node.Left, step + 1, false
	} else {
		if node.Right.Value[2] == 'Z' {
			return node.Left, step + 1, true
		}
		return node.Right, step + 1, false
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func walkLoop(node *Node, rute string, step int) (*Node, int) {
	direction := getNewDirection(rute, step)
	// fmt.Printf("node=%v, step=%d direction=%v, left=%v, right=%v\n", node.Value, step, string(direction), node.Left.Value, node.Right.Value)
	if direction == 'L' {
		if node.Left.Value == "ZZZ" {
			return nil, step + 1
		}
		return node.Left, step + 1
	} else {
		if node.Right.Value == "ZZZ" {
			return nil, step + 1
		}
		return node.Right, step + 1
	}
}

func getNewDirection(route string, step int) byte {
	if len(route) > step {
		return route[step]
	} else {
		routeStep := (step) - (len(route) * ((step) / len(route)))
		return route[routeStep]
	}
}
