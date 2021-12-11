package uuid

import (
	"github.com/google/uuid"
	base58 "github.com/jbenet/go-base58"
	gouuid "github.com/satori/go.uuid"
)

const base58Alphabet = base58.FlickrAlphabet

func New() string {
	return uuid.New().String()
}

func Base58Encode(v4 string) string {
	id, _ := gouuid.FromString(v4)
	return base58.EncodeAlphabet(id.Bytes(), base58Alphabet)
}
