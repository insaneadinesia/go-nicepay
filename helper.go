package nicepay

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func Sha256(s string) string {
	h := sha256.New()
	io.WriteString(h, s)
	bs := h.Sum(nil)

	result := fmt.Sprintf("%x", bs)

	return result
}
