package main

import (
  "fmt"
  "strings"
  "strconv"
)

type TestService struct {}

func (svc TestService) GetCharacterCountOccurencesInResult(n int, d int) int {
  count := 0
  for i := 0; i <= n; i++ {
    count += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
  }
  return count
}

func main() {
  // Execution Initiated
  fmt.Println("Hello, Interviewer!")
  fmt.Println("Executing Test Suite.")
  
  // Bootstrap Testing
  service:= TestService{}
  result:= service.GetCharacterCountOccurencesInResult(121,1)
  
  // Test Cases
  case1 := 82
  if result != case1 {
    fmt.Printf("ERROR: Case 1: The expected count %d is not equal to the actual count %d", case1, result)
  } else {
    fmt.Printf("PASS: Case 1: The expected count %d is equal to the actual count %d", case1, result)
  }
}
