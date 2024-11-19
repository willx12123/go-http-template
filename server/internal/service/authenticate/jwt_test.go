package authenticate

import (
	"testing"
)

func TestEncodeToken(t *testing.T) {
	_, err := EncodeToken(9527)
	if err != nil {
		t.Fatalf("Generate digest fail: %+v", err)
	}
}

func TestDecodeToken(t *testing.T) {
	var userID uint = 9527

	digest, _ := EncodeToken(userID)
	token, err := DecodeToken(digest)
	if err != nil {
		t.Fatalf("Parse digest fail: %+v", err)
	}

	if token.UserID != userID {
		t.Fatalf("Parse digest error, UserID should be: %d, but received: %d.", userID, token.UserID)
	}
}
