package main

import (
	"fmt"
	"testing"
)

func TestApp_getPriceText(t *testing.T) {

	open, _, _ := testApp.getPriceText()
	fmt.Println("open:", open.Text)
	if open.Text != "Open: $1969.7850 USD" {
		t.Error("wrong price returned", open.Text)
	}

}
