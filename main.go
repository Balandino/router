package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"router/engine"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func main() {
	engine.Route()
	compileFile("output.gv")
}

func compileFile(file string) {
	ctx := context.Background()

	dotContent, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", file, err)
	}

	graph, err := graphviz.ParseBytes(dotContent)
	if err != nil {
		log.Fatalf("Failed to parse DOT graph: %v", err)
	}

	defer graph.Close()

	g, err := graphviz.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize graphviz: %v", err)
	}

	defer g.Close()
	writeSVG(graph, g, ctx)
}

func writeSVG(graph *cgraph.Graph, g *graphviz.Graphviz, ctx context.Context) {
	var svgBuf bytes.Buffer
	if err := g.Render(ctx, graph, graphviz.SVG, &svgBuf); err != nil {
		log.Fatalf("Failed to render SVG: %v", err)
	}

	if err := os.WriteFile("output.svg", svgBuf.Bytes(), 0o644); err != nil {
		log.Fatalf("Failed to save output.svg: %v", err)
	}

	fmt.Println("Successfully generated: output.svg")
}
