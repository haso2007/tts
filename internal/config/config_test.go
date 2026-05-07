package config

import "testing"

func TestStripMarkdownRemovesEmojiIcons(t *testing.T) {
	processor := &SSMLProcessor{}

	input := "- \U0001f534 **Status**: fix \U0001f3af now \u2705 1\ufe0f\u20e3 item"
	got := processor.StripMarkdown(input)
	want := "Status: fix now item"

	if got != want {
		t.Fatalf("StripMarkdown() = %q, want %q", got, want)
	}
}

func TestStripEmojiIconsKeepsPlainDigits(t *testing.T) {
	input := "step 1 and version 2.0"
	got := stripEmojiIcons(input)

	if got != input {
		t.Fatalf("stripEmojiIcons() = %q, want %q", got, input)
	}
}
