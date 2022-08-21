package main

import (
	"testing"

	"fyne.io/fyne/test"
)

func Test_makeUI(t *testing.T) {

	var testCfg AppConfig

	edit, preview := testCfg.makeUI()

	test.Type(edit, "Hello")

	if preview.String() != "Hello" {

		t.Error("Failed,	did	not	find	expected	value	in	preview")

	}
}
