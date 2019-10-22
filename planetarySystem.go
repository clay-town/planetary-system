//This is my first go program!
// made with the help of Golang documenation and tutorialedge.net
package main

import (
        "fmt"
        "bufio"
        "os"
        "strings"
        "encoding/json"
        "io/ioutil"
        "time"
        "math/rand"
)

type Planets struct {
  Planets []Planet `json:"planets"`
}
type Planet struct {
  Name          string  `json:"name"`
  Description   string  `json: "description"`
}

func main(){
  rand.Seed(time.Now().UnixNano()) //why do we have to see this? I need to read more about it (found on StackOverflow)
  var reader = bufio.NewReader(os.Stdin)
  rLoc := rand.Intn(8)

  var iName = greeting(reader)
  fmt.Print("Nice to meet you, ",iName," My name is Eliza, I'm an old friend of Alexa. \nShall I randomly choose a planet for you to visit? (Y or N)\n")
  var rand = yesOrNo(reader)
  switch rand{
  case true:
    // do true things
    name, desc := readJsonFile(rLoc)
    printStuff(name, desc)

  case false:
    // do false things
    fmt.Print("Name the planet you would like to visit\n")
    var loc = validatePlanet(reader)
    name, desc := readJsonFile(loc)
    printStuff(name, desc)
  }
}

func printStuff(name string, desc string) {
  fmt.Println("Welcome to " + name + "...")
  fmt.Println(desc)
}

func validatePlanet(reader *bufio.Reader) int{
  iPlanet, _ := reader.ReadString('\n')
  switch strings.TrimRight(iPlanet, "\n") {
  case "Mercury","mercury":
    return 0
  case "Venus","venus":
    return 1
  case "Earth", "earth":
    return 2
  case "Mars", "mars":
    return 3
  case "Jupiter", "jupiter":
    return 4
  case "Saturn", "saturn":
    return 5
  case "Uranus", "uranus":
    return 6
  case "Neptune", "neptune":
    return 7
  case "Pluto", "pluto":
    return 8
  }
  fmt.Print("Please enter a valid planet name \n")
  return validatePlanet(reader)
}

func readJsonFile(loc int) (string, string){
  jsonFile, _ := os.Open("planetarySystem.json")
  byteValue, _ := ioutil.ReadAll(jsonFile)
  var planets Planets
  json.Unmarshal(byteValue, &planets)
  defer jsonFile.Close()
  return planets.Planets[loc].Name, planets.Planets[loc].Description
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
