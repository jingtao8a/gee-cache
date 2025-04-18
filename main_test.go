package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplitN(t *testing.T) {
	parts := strings.SplitN("fdlsaj/fdsajkfdsa", "/", 2)
	for _, part := range parts {
		fmt.Println(part)
	}
}
