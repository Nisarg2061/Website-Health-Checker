package main

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T){
  destination := "google.com"
  expected := fmt.Sprintf("[UP] %v is reachable.\n", destination)
  result := Check("google.com", "80")

  if result != expected {
    t.Error("Expected: " + expected + " got: " + result)
  }
}
