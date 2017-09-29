package main

import (
	//  "io/ioutil"
	"container/list"
	"fmt"
	"log"
	"os"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func (stack *Stack) PopAll() {
	if !stack.Empty() {
		fmt.Println(stack.Pop())
		stack.PopAll()
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
