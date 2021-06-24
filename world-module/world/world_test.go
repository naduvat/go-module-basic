package world

import (
    "testing"
)

func TestWorld(t *testing.T) {
    want := "world"
    got := Greet()
    if want != got {
        t.Errorf("Greet(), want=%s, got=%s", want, got)
    }
}
