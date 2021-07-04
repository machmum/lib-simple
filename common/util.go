package libcommon

import (
	"encoding/json"
	"fmt"
	"os"
)

// WrapWithHostname append hostname after base
func WrapWithHostname(base string) string {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	return fmt.Sprintf("%s-%s", base, hostname)
}

// NewRequestID accept id return id-hostname
// id MUST be unique
func NewRequestID(id string) string {
	return WrapWithHostname(id)
}

func JsonMarshallIgnoreError(src interface{}) (dest []byte) {
	dest, _ = json.Marshal(src)
	return
}
