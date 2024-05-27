package models

import (
	"testing"
)

var tableid = "71wh85i204kt8zx"

func TestGetTable(t *testing.T) {
	_, err := GetTable(tableid)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
