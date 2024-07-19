package main

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T){
  destination := "google.com"
  expected := fmt.Sprintf("[UP] %v is reachable.\n", destination)
  expected2 := fmt.Sprintf("[DOWN] %v is unreachable.\n", destination)
  //uses an unreachable host
  result := Check("synfullsolutions.com", "80")

  if result != expected && result != expected2{
    t.Error(fmt.Sprintf("Expected %v or %v got %v", expected, expected2, result))
  }
}
