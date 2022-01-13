package jibby

import (
	"errors"
	"fmt"
	"testing"
)

func TestParseError(t *testing.T) {
	var buf []byte
	_, err := Unmarshal([]byte(`{"hi": "hello", "worldo": powa}`), buf)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
	wrapped := fmt.Errorf("wrapped: %w", err)

	var pe *ParseError
	if !errors.As(err, &pe) {
		t.Fatal("error wasn't a ParseError")
	}
	if !errors.As(wrapped, &pe) {
		t.Fatal("wrapped error wasn't a ParseError")
	}
	// expectedErrMsg := "blah"
	// if pe.Error() != expectedErrMsg {
	// 	t.Fatalf("wtf expected %s but got %s", expectedErrMsg, pe.Error())
	// }
}
