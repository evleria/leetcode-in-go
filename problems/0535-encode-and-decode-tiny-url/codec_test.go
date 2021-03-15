package _535_encode_and_decode_tiny_url

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDecodeEncode(t *testing.T) {
	var urls []string
	for i := 0; i < 100; i++ {
		urls = append(urls, fmt.Sprintf("http://fake-url.com/%d", i))
	}

	m := make(map[string]string, len(urls))
	codec := Constructor()

	for _, url := range urls {
		m[url] = codec.encode(url)
	}

	for long, short := range m {
		assert.Check(t, is.Equal(long, codec.decode(short)))
	}
}
