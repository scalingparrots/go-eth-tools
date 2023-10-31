package message

import (
	"encoding/hex"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/storyicon/sigverify"
)

// SignMessage signs a message using a private key
func SignMessage(message string, privateKey string) (string, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	data := []byte(message)
	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), privateKeyECDSA)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(signature), nil
}

// VerifySignature verifies a signature using a public key
func VerifySignature(message string, signature string, address string) (bool, error) {
	valid, err := sigverify.VerifyEllipticCurveHexSignatureEx(
		ethcommon.HexToAddress(address),
		[]byte(message),
		signature,
	)

	if err != nil {
		return false, err
	}

	return valid, nil
}

// PrepareMessage prepares a message to be signed
func PrepareMessage(msg, addr string) string {
	return "\x19Ethereum Signed Message:\n" + string(len(msg)) + msg + addr
}
