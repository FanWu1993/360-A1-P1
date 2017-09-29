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

var allPath [][]string
var visit []int
var cost []int
var count []int
var s map[string]int

func readVertices() Graph {
	s = make(map[string]int)
	V := readInt("V")
	if 0 == V {
		fmt.Printf("No Star-system.\n")
		os.Exit(0)
	}
	g := make(Graph, V, V)
	count = make([]int, V, V)
	visit = make([]int, V, V)
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

func mapkey(m map[string]int, value int) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func DFS(g Graph, start string, end string) {
	stack := NewStack()
	var dive func(g Graph, start string, end string)
	dive = func(g Graph, start string, end string) {
		stack.Push(start)
		visit[s[start]] = 1
		for {
			if start == end {
				var path []string
				for i := 0; i < stack.Size(); i++ {
					path = append(path, stack.Get(i).(string))
				}
				allPath = append(allPath, path)
				stack.Pop()
				visit[s[start]] = 0
				break
			}
			for i := range g {
				if g[s[start]][i] {
					if 0 == visit[i] {
						star, _ := mapkey(s, i)
						dive(g, star, end)
					}
				}
			}
			stack.Pop()
			visit[s[start]] = 0
			break
		}
	}
	dive(g, start, end)
}

func main() {
	//	dat, _ := ioutil.ReadFile("./test2.in")
	//	fmt.Println(string(dat))
	g := readVertices()
	E := readInt("E")
	readEdge(g, E)
	DFS(g, "Scarif", "Yavin")

	for k, _ := range allPath {
		for _, v := range allPath[k] {
			if v != "Scarif" && v != "Yavin" {
				count[s[v]]++
			}
		}
	}
	var keypoint []string
	for k, v := range count {
		if v == len(allPath) {
			star, _ := mapkey(s, k)
			keypoint = append(keypoint, star)
		}
	}
	if nil == keypoint {
		fmt.Println("Leia escapes with the plans!")
	} else {

		temp := 0
		ctemp := 1<<31 - 1
		for k, v := range keypoint {
			if cost[s[v]] <= ctemp {
				ctemp = cost[s[v]]
				temp = k
			}
		}
		fmt.Printf("Darth blockades %s (%d).", keypoint[temp], cost[s[keypoint[temp]]])
	}
}
