package graph

import (
	"fmt"
	"github.com/Luna-CY/go-learn/data_structs"
)

type AdjTabGraph struct {
	vertexes []Vertex
	arcs     []*node

	vertexNumbers, arcNumber int

	traverseMarks []bool
}

func (atg *AdjTabGraph) Create(vertexes []Vertex, arcs []Arc) {
	atg.vertexes = vertexes
	atg.vertexNumbers = len(vertexes)

	atg.arcs = make([]*node, atg.vertexNumbers)
	for index, vertex := range vertexes {
		atg.arcs[index] = &node{vertex: vertex, index: index}
	}

	for _, arc := range arcs {
		node := &node{vertex: atg.vertexes[arc.Vj-1], index: arc.Vj - 1}
		node.next = atg.arcs[arc.Vi-1].next
		atg.arcs[arc.Vi-1].next = node
	}
}

func (atg *AdjTabGraph) dfs(index int) {
	atg.traverseMarks[index] = true

	arc := atg.arcs[index]
	fmt.Printf("Name: %s Value: %s\n", arc.vertex.Name, arc.vertex.Value)

	next := arc.next

	for nil != next {
		if !atg.traverseMarks[next.index] {
			atg.dfs(next.index)
		}

		next = next.next
	}
}

func (atg *AdjTabGraph) DFSTraverse() {
	atg.traverseMarks = make([]bool, atg.vertexNumbers)

	for index, isAccessed := range atg.traverseMarks {
		if !isAccessed {
			atg.dfs(index)
		}
	}
}

// breadth first search
func (atg *AdjTabGraph) BFSTraverse() {
	atg.traverseMarks = make([]bool, atg.vertexNumbers) // init marks

	queue := data_structs.Queue{}
	for index, isAccessed := range atg.traverseMarks {
		atg.traverseMarks[index] = true

		if !isAccessed {
			queue.Put(index)

			for !queue.IsEmpty() {
				index = queue.Get().(int)

				arc := atg.arcs[index]
				fmt.Printf("Name: %s Value: %s\n", arc.vertex.Name, arc.vertex.Value)

				next := arc.next

				for nil != next {
					if !atg.traverseMarks[next.index] {
						atg.traverseMarks[next.index] = true
						queue.Put(next.index)
					}

					next = next.next
				}
			}
		}
	}
}

type node struct {
	vertex Vertex
	index  int
	next   *node
}
