package main 

import (
  "testing"
  "strings"
  "fmt"
)

func TestGetAccessToken(t *testing.T) {
  s, err := getAccessToken()
  if err != nil {
    t.Fatal(err)
  }
  if s == "" {
    t.Fatal(fmt.Errorf("Null string returned"))
  }
  if !strings.Contains(s, "bearer") {
    t.Fatal(fmt.Errorf("Invalid token return"))
  }
}