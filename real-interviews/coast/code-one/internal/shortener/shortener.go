package shortener

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var (
	cache = map[string]string{}
)

func Hash(url string) string {
	b := []byte(url)

	hasher := sha256.New()
	hasher.Write(b)

	out := hex.EncodeToString(hasher.Sum(nil))
	cutString := out[0:11]

	cache[cutString] = url

	return cutString
}

func GetOriginalURL(shortUrl string) (*string, error) {
	val, ok := cache[shortUrl]
	if !ok {
		return nil, fmt.Errorf("url does not exist")
	}

	return &val, nil
}
