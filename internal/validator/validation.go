package validation

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
)

func ValidateSignature(address, signatureHex string) (bool, error) {
	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, err
	}

	message := []byte("Authenticate to DeNet service")
	msgHash := accounts.TextHash(message)

	pubKey, err := crypto.SigToPub(msgHash, signature)
	if err != nil {
		return false, err
	}

	recoveredAddress := crypto.PubkeyToAddress(*pubKey).Hex()

	return recoveredAddress == address, nil
}
