package tool

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func Sign(m map[string]string, secretKey string) string {
	size := len(m)
	if size == 0 {
		return ""
	}
	keys := make([]string, size)
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	pairs := make([]string, size)
	for i, key := range keys {
		pairs[i] = key + "=" + m[key]
	}
	str := strings.Join(pairs, "&") + secretKey
	byteData := []byte(str)
	has := md5.Sum(byteData)
	md5str := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str)
}
