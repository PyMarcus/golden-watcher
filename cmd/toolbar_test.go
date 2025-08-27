package main

import "testing"

func TestGetToolBar(t *testing.T) {
	tb := testApp.getToolBar()

	if len(tb.Items) != 4 {
		t.Errorf("Expected 4 items in toolbar, received: %d", len(tb.Items))
	}
}
