//This is my first go program!
package main

import (
        "fmt"
        "bufio"
        "os"
        "strings"
)

func main(){
  var reader = bufio.NewReader(os.Stdin)


  var iName = greeting(reader)
  fmt.Print("Nice to meet you, ",iName," My name is Eliza, I'm an old friend of Alexa. \nShall I randomly choose a planet for you to visit? (Y or N)\n")
  var rand = yesOrNo(reader)
  switch rand{
  case true:
    // do true things
    fmt.Print("true\n")
  case false:
    // do false things
    fmt.Print("false\n")
  }
}

func greeting(reader *bufio.Reader) string{
  //var dest = "Traveling to Pluto...Arrived at Pluto. I don't care what they say - it's a planet\n"
  fmt.Print("Welcome to the Solar System!\nThere are 9 planets to explore.\n What is your name?\n")
  iName, _ := reader.ReadString('\n') //_ returns a subset of the return values. ReadString returns two values string and error
  return iName
}

func yesOrNo(reader *bufio.Reader) bool{
  iYN, _ := reader.ReadString('\n')
  var input = strings.TrimRight(iYN, "\n")

  switch input {
  case "y", "Y", "Yes", "yes":
    return true
  case "n", "N", "No", "no":
    return false
  }
  fmt.Print("Please enter y or n \n")
  return yesOrNo(reader)

}



// func scanner() {
//   scanner := bufio.NewScanner(os.Stdin)
//   for scanner.Scan() {
//     fmt.Println(scanner.Text())
//     return
//   }
// }

//ReadRune() for single character input
//ReadString() for name input,

// this works, but only take the input upto the first whitespace
// fmt.Print("Enter text: ")
// var input string
// fmt.Scanln(&input)
// fmt.Print(input)
