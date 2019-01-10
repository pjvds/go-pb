package pb

import (
	"fmt"
	"strings"
)

// Printable returns a human readable version of a byte array. The bytes that correspond with
// ASCII printable characters [32-127) are passed through. Other bytes are
// replaced with \x followed by a two character zero-padded hex code for byte.
func Printable(d []byte) string {
	var sb strings.Builder
	for _, b := range d {
		if b >= 32 && b < 127 && b != '\\' {
			sb.WriteByte(b)
			continue
		}
		if b == '\\' {
			sb.WriteString("\\\\")
			continue
		}
		sb.WriteString(fmt.Sprintf("\\x%02x", b))
	}
	return sb.String()
}
