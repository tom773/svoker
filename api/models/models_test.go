package models

import (
	"fmt"
	"testing"
)

var tableid = "71wh85i204kt8zx"

func TestGetTable(t *testing.T) {
	table, err := GetTable(tableid)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Println(table)
}
