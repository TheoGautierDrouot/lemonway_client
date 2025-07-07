package openapi

import (
	"testing"
)

func TestClientCreation(t *testing.T) {
	client := NewAPIClient(NewConfiguration())
	if client == nil {
		t.Error("no client")
		return
	}
}
