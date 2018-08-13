package client

import (
	"bytes"
	"encoding/gob"

	"github.com/simple-rules/harmony-benchmark/blockchain"
	"github.com/simple-rules/harmony-benchmark/proto"
	"github.com/simple-rules/harmony-benchmark/proto/node"
)

// The specific types of message under CLIENT category
type ClientMessageType byte

const (
	TRANSACTION ClientMessageType = iota
	// TODO: add more types
)

// The types of messages used for CLIENT/TRANSACTION
type TransactionMessageType int

const (
	PROOF_OF_LOCK TransactionMessageType = iota // The proof of accept or reject returned by the leader to the client tnat issued cross shard transactions.
)

// [leader] Constructs the proof of accept or reject message that will be sent to client
func ConstructProofOfAcceptOrRejectMessage(proofs []blockchain.CrossShardTxProof) []byte {
	byteBuffer := bytes.NewBuffer([]byte{byte(proto.CLIENT)})
	byteBuffer.WriteByte(byte(TRANSACTION))
	byteBuffer.WriteByte(byte(PROOF_OF_LOCK))
	encoder := gob.NewEncoder(byteBuffer)

	encoder.Encode(proofs)
	return byteBuffer.Bytes()
}

// [client] Constructs the unlock to commit or abort message that will be sent to leaders
func ConstructUnlockToCommitOrAbortMessage(txsAndProofs []blockchain.Transaction) []byte {
	byteBuffer := bytes.NewBuffer([]byte{byte(proto.NODE)})
	byteBuffer.WriteByte(byte(node.TRANSACTION))
	byteBuffer.WriteByte(byte(node.UNLOCK))
	encoder := gob.NewEncoder(byteBuffer)
	encoder.Encode(txsAndProofs)
	return byteBuffer.Bytes()
}
