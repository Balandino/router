// Package engine contains the main code called to build the gv file
package engine

// digraph represents the main digraph container for all other resources
// only 1 should ever be made during program execution
type digraph struct {
	name     string
	contents []resource
}

// resource represents all resources and data in the resultant .gz file, excluding the digraph.
// Resources essentially split into two types - containers and resources.  Resources are the lowest
// type and just contain a value to be written to the resultant .gz file, often a single line to be written,
// which could be a setting like rankdir=LR, or a node or link.  Containers represent grouping
// resources like subgraph.  They can contain values and also resources.
type resource struct {
	name          string     // Required for containers, otherwise should be the const UNUSED
	containerType string     // Required for containers, otherwise should be the const UNUSED
	contents      []resource // Contains resources that could be other containers or just resources with values
	value         string     // The string to be written to the file
}
