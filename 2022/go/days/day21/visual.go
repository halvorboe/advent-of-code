package day21

import (
	"fmt"
	"log"

	graphviz "github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func (ast *AST) Show() {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	addNode(graph, ast.getReference(hashId(ROOT)))

	// 3. write to file directly
	if err := g.RenderFilename(graph, graphviz.PNG, "./graph.png"); err != nil {
		log.Fatal(err)
	}
}

func addNode(graph *cgraph.Graph, node *ASTNode) *cgraph.Node {
	if node.operand == VARI {
		curr, err := graph.CreateNode("VARIABLE")
		if err != nil {
			panic(err)
		}
		curr.SetColor("red")
		return curr
	} else if node.operand == VALU {
		curr, err := graph.CreateNode(fmt.Sprintf("%d = %d", node.name, node.value))
		if err != nil {
			panic(err)
		}
		curr.SetColor("gray")
		return curr
	} else {
		left := addNode(graph, node.left)
		right := addNode(graph, node.right)
		curr, err := graph.CreateNode(fmt.Sprintf("%d = %c", node.name, node.operand))
		if err != nil {
			panic(err)
		}
		if node.canMemoize {
			curr.SetColor("green")
		} else {
			curr.SetColor("blue")
		}

		_, err = graph.CreateEdge(fmt.Sprintf("%d -> %d", node.left.name, node.name), left, curr)
		if err != nil {
			panic(err)
		}
		_, err = graph.CreateEdge(fmt.Sprintf("%d -> %d", node.right.name, node.name), right, curr)
		if err != nil {
			panic(err)
		}

		return curr
	}

}
