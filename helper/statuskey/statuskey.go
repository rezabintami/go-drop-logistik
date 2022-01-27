package statuskey

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
)

func IsValid(orderID, StatusCode, Amount, SignKey, authKey string) error {
	key := fmt.Sprintf("%s%s%s%s", orderID, StatusCode, Amount, authKey)
	h512 := sha512.New()
	io.WriteString(h512, key)
	if fmt.Sprintf("%x", h512.Sum(nil)) != SignKey {
		return fmt.Errorf("%w: Invalid sign key", errors.New("bad request"))
	}
	return nil
}
