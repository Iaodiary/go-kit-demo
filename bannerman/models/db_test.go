package models

import (
	"testing"
)

func TestPing(t *testing.T) {
	if pingResult := ping(); pingResult != nil {
		t.Error(pingResult)
	}
}
