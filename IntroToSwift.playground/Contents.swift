import UIKit

// Strings
var str = "Hello, playground"   // Name representation of a value
var number = 20
var decimal = 50.5

let newDecimal = 40.5 // let makes variable constant
//newDecimal = 6.5 // Cant change value

// var newMessage: String // type annotation, only this variable can be of type newMessage


// Numbers + Type Safety
var myBankaccount: Double = -500.0
var my2ndInt: UInt = 100 // UInt should not be given to negative integers
var someVal: Float = 5.5
var sum = myBankaccount * Double(someVal)

// Terminal notes
// To make directory: mkdir terminal - directory name
// To make a file - touch myfile.__filetype
// to delete file - rm __filename
// rm -R --directory name recursively deletes everything in it
// cd.. goes backward one level


// Version Control
// You have to be in the folder of the repository you want to make
// git add filename
// git commit -m "added our first file"
// git status -> checks current status of repository


// To add a new branch
// git checkout -b filename -- creates new branch
// git add -A
// git commit -m "adding the new code to the new branch"
// git checkout ----
// delete branches
    // git branch -D branch name


// Add to the master branch
    // git merge anotherbranchname


// Conditionals
var trigger = true
if trigger {
    print("get water")
}   else {
    print("false")
}

func funkA () { print("function") }
funkA()

// Array
var names = ["jon", "jacob", "sadman"]
var countries = [String]() // Empty array of strings
countries.append("America")

var top3Colors = [String](repeating: "Boring Brown", count: 3)
top3Colors[0] = "Blue"

// Loops
for _ in 0...top3Colors.count - 1 { // for index from 0 to length - 1
    print("hello")
}

// Stack View
// use before relying on manual constraints for UI
// Think about vertical stack
// Then think about horizontal stack










