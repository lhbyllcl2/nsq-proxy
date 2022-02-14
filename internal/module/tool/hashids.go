package tool

import (
	"strings"
	"time"

	"github.com/speps/go-hashids/v2"
)

func AppIdEncode(id int64, salt string) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 12
	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64([]int64{id})
	return strings.ToLower(e)
}
func AppSecretEncode(id int64, salt string) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 30
	nowTimeTimestamp := time.Now().UnixNano() + id
	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64([]int64{nowTimeTimestamp})
	return e
}
