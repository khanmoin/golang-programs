package main

import "fmt"

func main() {
  // one way of declaring a string
  var text1 = "This is a string"

  // using shorthand
  text2 := "This is another sting"

  // string spaning multiple lines
  text3 := `This is a string 
  that spans multiple lines`

  fmt.Println(text1)
  fmt.Println(text2)
  fmt.Println(text3)
}
