package tx

import (
	"github.com/vechain/thor/acc"
	"github.com/vechain/thor/cry"
)

// Log represents a contract log event. These events are generated by the LOG opcode and
// stored/indexed by the node.
type Log struct {
	// address of the contract that generated the event
	Address acc.Address
	// list of topics provided by the contract.
	Topics []cry.Hash
	// supplied by the contract, usually ABI-encoded
	Data []byte
}