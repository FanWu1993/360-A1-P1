package main

import (
	//  "io/ioutil"
	"errors"
	"fmt"
	"log"
	"os"
)

type Stack struct {
	Element []interface{} //Element
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(value ...interface{}) {
	stack.Element = append(stack.Element, value...)
}

func (stack *Stack) Top() (value interface{}) {
	if stack.Size() > 0 {
		return stack.Element[stack.Size()-1]
	}
	return nil //read empty stack
}

func (stack *Stack) Pop() (err error) {
	if stack.Size() > 0 {
		stack.Element = stack.Element[:stack.Size()-1]
		return nil
	}
	return errors.New("Stack为空.") //read empty stack
}

func (stack *Stack) Get(idx int) (value interface{}) {
	if idx >= 0 && stack.Size() > 0 && stack.Size() > idx {
		return stack.Element[idx]
	}
	return nil //read empty stack
}

func (stack *Stack) Size() int {
	return len(stack.Element)
}

func (stack *Stack) Empty() bool {
	if stack.Element == nil || stack.Size() == 0 {
		return true
	}
	return false
}

func (stack *Stack) Print() {
	for i := len(stack.Element) - 1; i >= 0; i-- {
		fmt.Println(i, "=>", stack.Element[i])
	}
}

func readInt(s string) int {
	var i int
	_, err := fmt.Fscanf(os.Stdin, "%d", &i)
	if err != nil {
		log.Fatal(err)
	}
	if i < 0 {
		log.Fatalf(s + " must be greater than 0")
	}
	return i
}

func readString(s string) string {
	var star string
	_, err := fmt.Fscanf(os.Stdin, "%s", &star)
	if err != nil {
		log.Fatal(err)
	}
	return star
}

type Graph [][]bool

var visit []int
var cost []int
var path []string

var s map[string]int

func readVertices() Graph {
	s = make(map[string]int)
	V := readInt("V")
	if 0 == V {
		fmt.Printf("No Star-system.\n")
		os.Exit(0)
	}
	g := make(Graph, V, V)
	//visit := make([]int, V, V)
	for v := 0; v < V; v++ {
		c := 0
		name := readString("star-system")
		if name != string("Scarif") && name != string("Yavin") {
			c = readInt("cost")
		}
		s[name] = v
		cost = append(cost, c)
		g[v] = make([]bool, V, V)
	}
	return g
}

func readEdge(g Graph, E int) {
	for e := 0; e < E; e++ {
		s1 := readString("s1")
		s2 := ""
		_, err := fmt.Fscanf(os.Stdin, "%s", &s2)
		if err != nil {
			log.Fatal(err)
		}
		if s1 == s2 {
			log.Fatal("Loop at %s star-system not allow", s1)
		}
		g[s[s1]][s[s2]] = true
		g[s[s2]][s[s1]] = true
	}
}

func main() {
	//	dat, _ := ioutil.ReadFile("./test2.in")
	//	fmt.Println(string(dat))
	g := readVertices()
	E := readInt("E")
	readEdge(g, E)
	fmt.Println(g, E, cost)
}
