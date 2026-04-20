package domain

import (
	"encoding/base64"
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	tests := []struct {
		input    []byte