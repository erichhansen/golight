package golight

import (
	"testing"
)

func TestGetCommandData(t *testing.T) {
	data := getCommandData(command{color: Green})
	
	if (data == nil) {
		t.Errorf("Error, data was null")		
	}	
}