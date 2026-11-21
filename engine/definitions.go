// Package engine contains the main code called to build the gv file
package engine

// digraph represents the main digraph container for all other resources
// only 1 should ever be made during program execution
type digraph struct {
	name     string
	contents []resource
	links    []link
}

// resource represents all resources and data to be stored in the resultant .gz file, excluding the digraph.
// Though there is 1 Go resource type, it can essentially be two different types of struct - containers and resources.
// Resources are the simplest form and just contain data in the value field to be written to the resultant .gz file,
// often a single line to be written, which could be a setting like rankdir=LR, or a node or link.  Containers
// represent grouping resources like subgraph.  They can contain values and also other resources. If a resource struct has
// UNUSED for containerType, it's just a resource and shouldn't have any contents, just a value set.  Otherwise
// it is a container style resource which has contents to be evaluated when building the .gz.
type resource struct {
	name          string     // Optional, will only be used for containers
	containerType string     // Required for containers, otherwise should be the const UNUSED
	contents      []resource // Contains resources that could be other containers or just resources with values
	value         string     // The string to be written to the file
}

// link represents the edges between nodes
type link struct {
	from  string
	to    string
	label string
}
