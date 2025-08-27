package main

import "testing"

func TestGetPriceText(t *testing.T) {
	open, _, _ := testApp.GetPriceText()

	if open.Text != "Open: R$18295.6726 BRL" {
		t.Errorf("Expected: %s, but have %s", "Open: R$18295.6726 BRL", open.Text)
	}
}
