# Nose2Node

> Sniff. Spread. Sync.

Nose2Node is a decentralized, gossip-based messaging system where every node is both sender and spreader. Built on peer-to-peer communication, this project explores distributed systems, eventual consistency, and intentional message drift through a playful, experimental lens.

## What is Nose2Node?

In Nose2Node, messages aren't sent directly to a server. Instead, they're whispered node-to-node, spreading organically like rumors. Each node receives a complete message and passes it onward to its peers, allowing the network to sync naturally—but never instantly.

## Key Features

- Decentralized Gossip Protocol: No central server, just peer-to-peer message passing.
- Intentional Drift: Messages may subtly mutate, creating multiple narrative versions across nodes.
- Eventual Consistency: Every node eventually hears everything—but at its own pace.
- Network Visualization: See messages hop, mutate, and propagate visually.

## Tech Stack

- Go: Lightweight concurrency with goroutines and channels, easy networking.
- WebSockets: Reliable, real-time node communication.
- React & D3.js: Interactive frontend and dynamic network visualizations.

## Why Gossip?

Gossip protocols excel at fault-tolerance, scalability, and simplicity. They're at the heart of systems like Cassandra, BitTorrent, and blockchain networks, providing resilience and graceful degradation under chaos.

## Goals of the Project

- Demonstrate practical distributed system concepts.
- Experiment with the natural drift of messages across a network.
- Provide a clear, visual way to understand how information propagates in decentralized systems.

## Getting Started

(Coming soon: installation and setup guide.)

## Future Improvements

- Enhanced drift visualizations
- Advanced AI-powered message mutation
- Robust fault injection scenarios

---

Nose2Node: because sometimes consistency is overrated.
