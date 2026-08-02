package conv

import "encoding/base64"

// takes a byte slice and convert it into a base 64 string
func B64UrlEncoding(b []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}
