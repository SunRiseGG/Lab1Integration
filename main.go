package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Current project version:", BuildVersion)
  fmt.Println("Input your expression: ")
  fmt.Print("~> ")
  text, _ := reader.ReadString('\n')
  text = strings.Replace(text, "\n", "", -1)
  text = strings.Replace(text, "\r", "", -1)

  result, err := PrefixToPostfix(text)
  fmt.Println("-----------------------------------")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(result)
}
