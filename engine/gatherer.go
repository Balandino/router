package engine

func Route() {
	ptrDigraph := getDigraph("gogv")

	ptrVPC := getVPC("c0000")
	ptrPrivSubnet := getPrivateSubnet("sub-1")

	ptrEC2 := getEC2("i-123")
	ptrPrivSubnet.addResource(ptrEC2)

	ptrEC2 = getEC2("i-789")
	ptrPrivSubnet.addResource(ptrEC2)

	ptrVPC.addResource(ptrPrivSubnet)

	ptrPrivSubnet2 := getPrivateSubnet("sub-2")
	// ptrPrivSubnet2.addBlankNode()
	ptrVPC.addResource(ptrPrivSubnet2)

	ptrDigraph.addContainer(ptrVPC)
	ptrDigraph.addLink("i-123", "i-789", "My Link")
	ptrDigraph.addLink("sub-1", "sub-2", "subnet link")

	writeToFile(*ptrDigraph)
}
