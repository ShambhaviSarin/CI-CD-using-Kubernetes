package greet

import "testing"

func TestHello_Empty(t *testing.T) {
    got := Hello("")
    want := "Hello, World!"
    if got != want {
        t.Fatalf("Hello(\"\") = %q; want %q", got, want)
    }
}

func TestHello_Name(t *testing.T) {
    got := Hello("Alice")
    want := "Hello, Alice!"
    if got != want {
        t.Fatalf("Hello(\"Alice\") = %q; want %q", got, want)
    }
}
