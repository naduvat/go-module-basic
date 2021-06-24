package hello

import (
    "testing"
)

func TestHello(t *testing.T) {
    want := "hello"
    got := Greet()
    if want != got {
        t.Errorf("Greet(), want=%s, got=%s", want, got)
    }
}
