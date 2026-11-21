package engine

func Route() {
	ptrDigraph := getDigraph("gogv")

	// ptrVPC := getVPC("c0000")
	// ptrEC2 := getEC2("i-123")
	// ptrVPC.addResource(ptrEC2)
	//
	// ptrEC2 = getEC2("i-789")
	// ptrVPC.addResource(ptrEC2)
	//
	// ptrVPC.addLink("i-123", "i-789", "My Link")

	ptrVPC := getVPC("c0000")
	ptrPrivSubnet := getPrivateSubnet("sub-0c123d321341")

	ptrEC2 := getEC2("i-123")
	ptrPrivSubnet.addResource(ptrEC2)

	ptrEC2 = getEC2("i-789")
	ptrPrivSubnet.addResource(ptrEC2)

	ptrVPC.addResource(ptrPrivSubnet)
	ptrVPC.addLink("i-123", "i-789", "My Link")

	ptrDigraph.addContainer(ptrVPC)
	writeToFile(*ptrDigraph)
}
