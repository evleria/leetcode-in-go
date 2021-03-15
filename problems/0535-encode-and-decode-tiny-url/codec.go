package _535_encode_and_decode_tiny_url

import (
	"crypto/sha256"
	"encoding/base64"
)

var base = "http://tinyurl.com/"

type Codec struct {
	m map[string]string
}

func Constructor() Codec {
	return Codec{
		m: map[string]string{},
	}
}

func (c *Codec) encode(longUrl string) string {
	hash := sha256.Sum256([]byte(longUrl))
	suffix := base64.StdEncoding.EncodeToString(hash[:16])[:8]
	c.m[suffix] = longUrl

	return base + suffix
}

func (c *Codec) decode(shortUrl string) string {
	suffix := shortUrl[len(base):]
	return c.m[suffix]
}
