package engine

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// UNUSED is the universal non-use string for structs
const UNUSED = "-"

// writeToFile writes the built resources to a .gz file
func writeToFile(digraph digraph) {
	var output strings.Builder
	output.WriteString("digraph ")

	fmt.Fprintf(&output, "%s {\n", digraph.name)

	for _, content := range digraph.contents {
		content := processContents(content, 1)
		output.WriteString(content)
	}

	for _, link := range digraph.links {
		fmt.Fprintf(&output, "%s -> %s %s\n", link.from, link.to, link.label)
	}

	output.WriteString("}")
	// fmt.Println(output.String())

	err := os.WriteFile("output.gv", []byte(output.String()), 0o644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

// processContents recursively builds resources into a string for writing to the file
func processContents(content resource, tabDepth int) string {
	var output strings.Builder

	spacingPrefix := strings.Repeat("   ", tabDepth)
	if content.containerType == UNUSED {
		if content.name != UNUSED {
			return spacingPrefix + content.name + " " + formatBlock(content.value, tabDepth)
		}
		return formatBlock(content.value, tabDepth)
	}

	fmt.Fprintf(&output, "%s%s %s {\n", spacingPrefix, content.containerType, content.name)
	for _, c := range content.contents {
		output.WriteString(processContents(c, tabDepth+1))
	}

	fmt.Fprintf(&output, "%s}\n", spacingPrefix)
	return output.String()
}

// RemoveDashes is needed as certain resources won't render if they have a - in
func RemoveDashes(s string) string {
	return strings.ReplaceAll(s, "-", "")
}

// addResource adds another resource to the contents of another
func (ptrContainer *resource) addResource(ptrResource *resource) {
	ptrContainer.contents = append(ptrContainer.contents, *ptrResource)
}

// addResource adds a new resource to the digraph
func (ptrDigraph *digraph) addResource(ptrResource *resource) {
	ptrDigraph.contents = append(ptrDigraph.contents, *ptrResource)
}

// addContent adds a new resource existing content of a resource
func (ptrContainer *resource) addContent(value string) {
	prefilledResource := getPrefilledResource(value)
	ptrContainer.contents = append(ptrContainer.contents, *prefilledResource)
}

// addContent adds a new resource to the existing content of the digraph
func (ptrDigraph *digraph) addContent(value string) {
	prefilledResource := getPrefilledResource(value)
	ptrDigraph.contents = append(ptrDigraph.contents, *prefilledResource)
}

// multiAddContent enables adding multiple new resources to the existing content of the digraph
func (ptrDigraph *digraph) multiAddContent(values []string) {
	for _, value := range values {
		prefilledResource := getPrefilledResource(value)
		ptrDigraph.contents = append(ptrDigraph.contents, *prefilledResource)
	}
}

// multiAddContent enables adding multiple new resources to the existing content of a resource
func (ptrContainer *resource) multiAddContent(values []string) {
	for _, value := range values {
		prefilledResource := getPrefilledResource(value)
		ptrContainer.contents = append(ptrContainer.contents, *prefilledResource)
	}
}

// addContainer adds a container resource, such as a subgraph, which should contain its own resources
func (ptrDigraph *digraph) addContainer(ptrContainer *resource) {
	ptrDigraph.contents = append(ptrDigraph.contents, *ptrContainer)
}

// addNewline adds spacing to the digraph for better formatting in the output file
func (ptrDigraph *digraph) addNewline() {
	prefilledResource := getPrefilledResource("")
	ptrDigraph.contents = append(ptrDigraph.contents, *prefilledResource)
}

// addLink adds a link/edge between nodes to the digraph
func (ptrDigraph *digraph) addLink(from, to, label string) {
	fmtLabel := fmt.Sprintf(`[label="%s"]`, label)
	ptrDigraph.links = append(ptrDigraph.links, *getLink(RemoveDashes(from), RemoveDashes(to), fmtLabel))
}

// // addLink adds a link/edge between nodes
// func (ptrDigraph *digraph) addLink(from, to, label string) {
// 	ptrDigraph.addContent(fmt.Sprintf(`%s -> %s [label="%s"]`, RemoveDashes(from), RemoveDashes(to), label))
// }

// formatBlock cleans up multi-line text (like HTML tables or DOT attributes)
// by removing common source-code indentation (dedenting) and re-indenting each
// line according to the specified nesting depth (tabDepth).
func formatBlock(text string, tabDepth int) string {
	// Standardize tabs to 3 spaces to ensure consistent indentation width.
	text = strings.ReplaceAll(text, "\t", "   ")

	// Strip leading/trailing newlines from the whole block, then split into lines.
	lines := strings.Split(strings.Trim(text, "\r\n"), "\n")

	// Step 1: Calculate the common leading indentation (margin) across non-empty lines.
	minIndent := -1
	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		// Skip blank or whitespace-only lines so they don't skew the minimum indentation.
		if trimmed == "" {
			continue
		}
		// Count the number of leading spaces on this line.
		indent := len(l) - len(strings.TrimLeft(l, " "))
		// Keep track of the smallest leading indent found so far.
		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}
	// If the entire text was blank, set minimum indent to 0.
	if minIndent == -1 {
		minIndent = 0
	}

	// Step 2: Re-indent each line according to tabDepth.
	// Build the indentation prefix (3 spaces per nesting level).
	prefix := strings.Repeat("   ", tabDepth)
	var sb strings.Builder
	for _, line := range lines {
		// Preserve blank lines without leaving trailing whitespace.
		if strings.TrimSpace(line) == "" {
			sb.WriteString("\n")
			continue
		}
		// Strip the common leading indentation from the line (dedent).
		if len(line) >= minIndent {
			line = line[minIndent:]
		}
		// Write the current nesting prefix, the stripped line, and a newline.
		sb.WriteString(prefix)
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	return sb.String()
}
