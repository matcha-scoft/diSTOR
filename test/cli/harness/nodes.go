package harness

import (
	"github.com/multiformats/go-multiaddr"
)

// Nodes is a collection of Kubo nodes along with operations on groups of nodes.
type Nodes []*Node

func (n Nodes) Init(args ...string) Nodes {
	for _, node := range n {
		node.Init()
	}
	return n
}

func (n Nodes) Connect() Nodes {
	for i, node := range n {
		for j, otherNode := range n {
			if i == j {
				continue
			}
			node.Connect(otherNode)
		}
	}
	for _, node := range n {
		firstPeer := node.Peers()[0]
		if _, err := firstPeer.ValueForProtocol(multiaddr.P_P2P); err != nil {
			log.Panicf("unexpected state for node %d with peer ID %s: %s", node.ID, node.PeerID(), err)
		}
	}
	return n
}

func (n Nodes) StartDaemons() Nodes {
	for _, node := range n {
		node.StartDaemon()
	}
	return n
}

func (n Nodes) StopDaemons() Nodes {
	for _, node := range n {
		node.StopDaemon()
	}
	return n
}
