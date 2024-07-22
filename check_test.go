package main

import (
	"fmt"
	"testing"
)

func TestCheckReachable(t *testing.T){
  destination := "google.com"
  expected := fmt.Sprintf("[UP] %v is reachable.\n", destination)
  //uses an reachable host
  result := Check("google.com", "80")

  if result != expected {
    t.Error(fmt.Sprintf("Expected %v got %v", expected, result))
  }
}

func TestCheckUnreachable(t *testing.T){
  destination := "synfullsolutions.com"
  expected2 := fmt.Sprintf("[DOWN] %v is unreachable.\n", destination)
  //uses an unreachable host
  result := Check("synfullsolutions.com", "80")

  if result !=  expected2{
    t.Error(fmt.Sprintf("Expected %v got %v", expected2, result))
  }
}
