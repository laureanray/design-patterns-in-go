package main

import "fmt"

func main() {
  s := GetInstance()

  fmt.Println(s.count)

  s.AddOne()

  b := GetInstance()

  fmt.Println(b.count)

  b.AddOne()
}
