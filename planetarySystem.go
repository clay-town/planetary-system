//This is my first go program!
package main

import "fmt"

func main(){
  reader:= bufio.NewReader()
  var greeting = "Welcome to the Solar System!\nThere are 9 planets to explore.\n What is your name?\n\n"
  var nResponse = "Nice to meet you, <name>. My name is Eliza, I'm an old friend of Alexa. \nShall I randomly choose a planet for you to visit? (Y or N)\n\n"
  var dest = "Traveling to Pluto...Arrived at Pluto. I don't care what they say - it's a planet\n"

  fmt.Printf(greeting)
  fmt.Printf(nResponse)
  fmt.Printf(dest)
}


//ReadRune() for single character input
//ReadString() for name input
