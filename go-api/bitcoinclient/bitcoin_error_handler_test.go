package bitcoinclient

import (
	"testing"
)

func TestAuthenticateRPCResponse(t *testing.T) {
	correctResponse := `{"result":{"txid":"b1fea52486ce0c62bb442b530a3f0132b826c74e473d1f2c220bfa78111c5082","hash":"b1fea52486ce0c62bb442b530a3f0132b826c74e473d1f2c220bfa78111c5082","version":1,"size":134,"vsize":134,"locktime":0,"vin":[{"coinbase":"04ffff001d0102","sequence":4294967295}],"vout":[{"value":50.00000000,"n":0,"scriptPubKey":{"asm":"04d46c4968bde02899d2aa0963367c7a6ce34eec332b32e42e5f3407e052d64ac625da6f0718e7b302140434bd725706957c092db53805b821a85b23a7ac61725b OP_CHECKSIG","hex":"4104d46c4968bde02899d2aa0963367c7a6ce34eec332b32e42e5f3407e052d64ac625da6f0718e7b302140434bd725706957c092db53805b821a85b23a7ac61725bac","reqSigs":1,"type":"pubkey","addresses":["1PSSGeFHDnKNxiEyFrD1wcEaHr9hrQDDWc"]}}]},"error":null,"id":null}`
	result := bitcoinClient.AuthenticateRPCResponse(correctResponse)

	if result != nil {
		t.Errorf("Result should have passed since we passed a valid response")
	}

	failedResponse := `{"result":null,"error":{"code":-32601,"message":"Method not found"},"id":null}`
	resultFailedResponse := bitcoinClient.AuthenticateRPCResponse(failedResponse)

	if resultFailedResponse == nil {
		t.Errorf("Result should have failed")
	}
}