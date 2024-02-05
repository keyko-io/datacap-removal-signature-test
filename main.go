package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	verifregst "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	"github.com/filecoin-project/go-state-types/crypto"
	filsigner "github.com/jsign/go-filsigner/wallet"
)

func SerializeSignature(signature *crypto.Signature) string {
	if signature == nil {
		return ""
	}

	dataStr := hex.EncodeToString(signature.Data)
	return fmt.Sprintf("Serialized Signature: %s", dataStr)
}

func main() {
	// assign datacap to a client
	datacapToRemove := big.NewInt(1000)
	clientAddress := "t13kfk4abtjrujaflzghcuksd73zasf33l44iupci"
	addr, err := address.NewFromString(clientAddress)
	if err != nil {
		fmt.Println("ERROR in client adress parsing")
	}

	removeProposal := verifregst.RemoveDataCapProposal{
		VerifiedClient:    addr,
		DataCapAmount:     datacapToRemove,
		RemovalProposalID: verifregst.RmDcProposalID{ProposalID: 0},
	}

	buf := bytes.Buffer{}
	if err := removeProposal.MarshalCBOR(&buf); err != nil {
		fmt.Println("Error in CBOR marshalling:", err)
		return
	}

	removeProposalSer := buf.Bytes()

	privKey := "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a22776d496b31734d4f2f5474374c3675725878413246396171677547686447713763356b63493548765767593d227d"
	privKey2 := "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a22574139324c5267736b7a4b546b2b584c4668524f4e526277547855316c6f564172496d78443831536a53493d227d"
	// Call the separate signing function
	signedData1, err := filsigner.WalletSign(privKey, removeProposalSer)
	if err != nil {
		fmt.Println("Error signing datacap removal proposal:", err)
		return
	}

	signedData2, err := filsigner.WalletSign(privKey2, removeProposalSer)
	if err != nil {
		fmt.Println("Error signing datacap removal proposal:", err)
		return
	}
	serializedSignature1 := SerializeSignature(signedData1)
	serializedSignature2 := SerializeSignature(signedData2)
	fmt.Println("Signature 1: ", serializedSignature1)
	fmt.Println("Signature 2: ", serializedSignature2)
}
