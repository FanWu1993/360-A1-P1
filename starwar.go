package main

import (
	"fmt"
	//	"io/ioutil"
	"log"
	"os"
)

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
	_, err := fmt.Fscanf(os.Stdin, "\n%s", &star)
	if err != nil {
		log.Fatal(err)
	}
	return star
}

type StarSystem string
type Neighbourhood []StarSystem
type Value int
type Vertex struct {
	name       StarSystem
	neighbours Neighbourhood
	cost       Value
}

func degree(v Vertex) int {
	return len(v.neighbours)
}

type Graph []Vertex

func readVertices() Graph {
	V := readInt("V")
	if 0 == V {
		fmt.Printf("No Star-system.\n")
		os.Exit(0)
	}
	g := make(Graph, 0, V)
	for v := 0; v < V; v++ {
		cost := 0
		name := readString("star-system")
		if name != string("Scarif") && name != string("Yavin") {
			cost = readInt("cost")
		}
		g = append(g, Vertex{StarSystem(name), make(Neighbourhood, 0, V), Value(cost)})
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
		for i := range g {
			if g[i].name == StarSystem(s1) {
				g[i].neighbours = append(g[i].neighbours, StarSystem(s2))
			}
			if g[i].name == StarSystem(s2) {
				g[i].neighbours = append(g[i].neighbours, StarSystem(s1))
			}
		}
	}
}
func main() {
	//	dat, _ := ioutil.ReadFile("./test2.in")
	//	fmt.Println(string(dat))
	g := readVertices()
	fmt.Fscanf(os.Stdin, "\n")
	E := readInt("E")
	readEdge(g, E)
	fmt.Println(g)
}
