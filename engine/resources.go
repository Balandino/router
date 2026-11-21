package engine

import "fmt"

// getResource returns the building block type of the program.
// containerType and name should be UNUSED if just a resource with a value
func getResource(containerType, name string) *resource {
	return &resource{
		name:          name,
		containerType: containerType,
		contents:      []resource{},
		value:         "",
	}
}

// getSubgraph builds a container type resource that can contain both values and other resources
func getSubgraph(name string) *resource {
	// Containers don't seem to render without the prefix cluster_ in the name
	return getResource("subgraph", fmt.Sprintf("cluster_%s", RemoveDashes(name)))
}

// getPrefilledResource builds a resource just containing a value
func getPrefilledResource(value string) *resource {
	return &resource{
		name:          UNUSED,
		containerType: UNUSED,
		contents:      []resource{},
		value:         value,
	}
}

// getDigraph returns the digraph.  Should only be called once during the program
func getDigraph(name string) *digraph {
	ptrDigraph := digraph{
		name:     name,
		contents: []resource{},
	}

	ptrDigraph.multiAddContent([]string{
		`rankdir=LR`,
		`bgcolor="black"`,
		`node [shape=box, style=rounded, color="white", fontcolor="white", fontname=Helvetica, fontsize=12]`,
		`edge [fontname=Helvetica, fontsize=10, fontcolor="white", color="white"]`,
		`graph [fontcolor="white", color="white"]`,
	})

	return &ptrDigraph
}

// getVPC returns a vpc styled resource prefilled with a top left icon
func getVPC(name string) *resource {
	ptrVPC := getSubgraph(name)
	ptrVPC.addContent(fmt.Sprintf(`labelloc="t" labeljust="l" color="purple" label=<
			<TABLE BORDER="0" CELLBORDER="0" CELLSPACING="4" CELLPADDING="0">
				<TR>
					<TD><IMG SRC="icons/Architecture-Group-Icons_07312026/Virtual-private-cloud-VPC_32.svg"/></TD>
               <TD><B>%s</B></TD>
            </TR>
         </TABLE>
      >`, name))

	return ptrVPC
}

// getEC2 returns a ec2 styles resource prefilled with an icon
func getEC2(instanceID string) *resource {
	return getPrefilledResource(fmt.Sprintf(`%s [shape=none, fontcolor="white", label=<
       <TABLE BORDER="0" CELLBORDER="0" CELLSPACING="2">
          <TR>
			 	<TD><IMG SRC="icons/Architecture-Service-Icons_07312026/Arch_Compute/32/Arch_Amazon-EC2_32.svg"/></TD>
			 </TR>
          <TR>
			 	<TD>%s</TD>
			 </TR>
       </TABLE>
        >]`, RemoveDashes(instanceID), instanceID))
}
