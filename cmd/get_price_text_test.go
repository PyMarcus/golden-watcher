package main

import (
	"testing"
)

func TestGetPriceText(t *testing.T) {
	open, _, _ := testApp.GetPriceText()
	res := open.Text
	if res != "Abertura: R$18295.6726 BRL" {
		t.Errorf("Expected: %s, but have %s", "Abertura: R$18295.6726 BRL", res)
	}
}
