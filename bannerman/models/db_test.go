package models

import (
	"testing"
)

func TestPing(t *testing.T) {
	if err := db.Ping(); err != nil {
		t.Error(err)
	}
}
