package rlp

import "golang.org/x/crypto/sha3"

func RLPHash(x interface{}) []byte {
	hw := sha3.NewLegacyKeccak256()
	Encode(hw, x)
	return hw.Sum(nil)
}
