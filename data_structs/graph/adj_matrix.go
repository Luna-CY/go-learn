package graph

import (
	"fmt"
	"github.com/Luna-CY/go-learn/data_structs"
)

// 无向图：邻接矩阵

type AdjMatGraph struct {
	vertexes []Vertex
	arcs     [][]int

	vertexNumbers, arcNumbers int

	traverseMarks []bool
}

func (amg *AdjMatGraph) Create(vertexes []Vertex, arcs []Arc) {
	amg.vertexes = vertexes
	amg.vertexNumbers = len(vertexes)
	amg.arcs = make([][]int, amg.vertexNumbers)

	for i := 0; i < amg.vertexNumbers; i++ {
		amg.arcs[i] = make([]int, amg.vertexNumbers)
		for j := 0; j < amg.vertexNumbers; j++ {
			amg.arcs[i][j] = IntMax
		}
	}

	arcNumbers := 0
	for _, arc := range arcs {
		row := arc.Vi - 1
		col := arc.Vj - 1

		if row >= len(amg.arcs) {
			continue
		}

		if col >= len(amg.arcs[row]) {
			continue
		}

		if amg.arcs[row][col] == IntMax {
			arcNumbers++
		}

		amg.arcs[row][col] = arc.Weight
		amg.arcs[col][row] = arc.Weight
	}

	amg.arcNumbers = arcNumbers
}

func (amg *AdjMatGraph) GetValue(v int) string {
	if v > amg.vertexNumbers || v <= 0 {
		return ""
	}

	return amg.vertexes[v-1].Value
}

func (amg *AdjMatGraph) PutValue(v int, value string) {
	if v > amg.vertexNumbers || v <= 0 {
		return
	}

	amg.vertexes[v].Value = value
}

func (amg *AdjMatGraph) dfs(index int) {
	vertex := amg.vertexes[index]
	fmt.Printf("Name: %s Value: %s\n", vertex.Name, vertex.Value)

	amg.traverseMarks[index] = true

	for i := 0; i < amg.vertexNumbers; i++ {
		if IntMax != amg.arcs[index][i] && !amg.traverseMarks[i] {
			amg.dfs(i)
		}
	}
}

// depth first search
func (amg *AdjMatGraph) DFSTraverse() {
	amg.traverseMarks = make([]bool, amg.vertexNumbers) // init marks

	for index, isAccessed := range amg.traverseMarks {
		if !isAccessed {
			amg.dfs(index)
		}
	}
}

// breadth first search
func (amg *AdjMatGraph) BFSTraverse() {
	amg.traverseMarks = make([]bool, amg.vertexNumbers) // init marks

	queue := data_structs.Queue{}
	for index, isAccessed := range amg.traverseMarks {
		amg.traverseMarks[index] = true

		if !isAccessed {
			queue.Put(index)

			for !queue.IsEmpty() {
				index = queue.Get().(int)

				vertex := amg.vertexes[index]
				fmt.Printf("Name: %s Value: %s\n", vertex.Name, vertex.Value)

				for i := 0; i < amg.vertexNumbers; i++ {
					if IntMax != amg.arcs[index][i] && !amg.traverseMarks[i] {
						amg.traverseMarks[i] = true
						queue.Put(i)
					}
				}
			}
		}
	}
}

func (amg *AdjMatGraph) MiniSpanTreePrim() {
	vertexes := make([]Vertex, amg.vertexNumbers)
	lowcost := make([]int, amg.vertexNumbers)

	lowcost[0] = 0
	vertexes[0] = amg.vertexes[0]

	for i := 1; i < amg.vertexNumbers; i++ {
		lowcost[i] = amg.arcs[0][i]
		vertexes[i] = amg.vertexes[0]
	}

	for i := 1; i < amg.vertexNumbers; i++ {
		min := IntMax

		k := 0
		for j := 1; j < amg.vertexNumbers; j++ {
			if lowcost[j] != 0 && lowcost[j] < min {
				min = lowcost[j]
				k = j
			}
		}

		fmt.Printf("%s, %s\n", vertexes[k].Value, amg.vertexes[k].Value)
		lowcost[k] = 0

		for j := 1; j < amg.vertexNumbers; j++ {
			if lowcost[j] != 0 && amg.arcs[k][j] < lowcost[j] {
				lowcost[j] = amg.arcs[k][j]
				vertexes[j] = amg.vertexes[k]
			}
		}
	}
}
