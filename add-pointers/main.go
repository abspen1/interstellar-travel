package main

import "fmt"

/*
This is a fun programming project to help learn Golang
I implement conditional statements such as; if/else and switch
Also use the fmt Scan, Print, Printf, Println
Functions heavy project
Most of these functions are actually go modules that pass our struct pointer
*/

// Traveler struct
type Traveler struct {
	initalMoney  int
	balance      int
	location     string
	planetChoice string
	fuel         int
	fuelNeeded   int
}

func main() {
	// Create user with type traveler
	traveler := Traveler{
		location: "Earth",
	}

	fmt.Println("Welcome to Interstellar Travel Agency!")
	fmt.Print("This travel agency can take you to a few different planets!\n")
	fmt.Print("How much money do you have? Keep in mind it's pricey to travel: ")
	fmt.Scan(&traveler.balance)
	traveler.initalMoney = traveler.balance

	if traveler.balance >= 100000 {
		fmt.Println("Wow! That's great.")
		fmt.Println("You can afford the first tank of gas.")
		traveler.gasStation()
	} else {
		fmt.Println("Oof. Broke boi.")
		fmt.Println("Don't worry, the first tank of gas is on us.")
		traveler.fuel = 1000000
	}
	//One way to use a do-while loop in Go syntax
	actualInput := true
	for ok := true; ok; ok = actualInput {
		actualInput = traveler.userInput()
	}
	fmt.Println("Thanks so much for coming!")
	fmt.Print("Your balance after this vacation is $", traveler.balance, "\n\n")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
	fmt.Println("Lets go to the gas station.")
}

// Create the function flyToPlanet() here
func (t *Traveler) flyToPlanet() {
	var fuelCost int
	fuelCost = calculateFuel(t.planetChoice)
	if t.fuel >= fuelCost {
		t.location = t.planetChoice
		fmt.Print("Welcome to the planet ", t.location, ", please be cautious as we are far from home!\n")
		t.fuel -= fuelCost
	} else {
		cantFly()
		t.gasStation()
		fmt.Printf("Do you still want to travel to %s if so enter 'yes': ", t.planetChoice)
		var response string
		fmt.Scan(&response)
		if response == "yes" {
			t.flyToPlanet()
		}
	}
}

//Create a function takeUsHome() here
func (t *Traveler) takeUsHome() {
	var fuelCost int
	fuelCost = calculateFuel(t.location)
	if t.fuel >= fuelCost {
		switch t.location {
		case "Venus":
			t.fuel -= fuelCost
			t.location = "Earth"
		case "Mercury":
			t.fuel -= fuelCost
			t.location = "Earth"
		case "Mars":
			t.fuel -= fuelCost
			t.location = "Earth"
		case "Earth":
			fmt.Println("You are already home, silly!")
		default:
			fmt.Println("I think we're lost")
		}
		fmt.Println("Welcome Back Home :)")

	} else {
		cantFly()
		t.gasStation()
		fmt.Print("Do you still want to travel home: ")
		var response string
		fmt.Scan(&response)
		if response == "yes" {
			t.takeUsHome()
		}
	}
}
func (t *Traveler) adjustBalance(cost int) {
	cost *= t.fuelNeeded
	if cost > t.balance {
		fmt.Println("You can't afford to fill up the tank.. yikes.")
	} else {
		//fmt.Printf("You just purchased $%v worth of gas.\n", cost)
		fmt.Printf("You bought %v gallons of gas for $%v.\n\n", t.fuelNeeded, cost)
		t.balance -= cost
		t.fuel = 1000000
	}
}
func (t *Traveler) calculateFuelNeeded() {
	// We get 1000000 per tank so lets get our gallons per tank assuming efficiency of 100mpg
	t.fuelNeeded = 1000000 / 1000
	currentGallons := t.fuel / 1000
	t.fuelNeeded -= currentGallons
}

func (t *Traveler) calculateFuelCost() {
	var cost int
	switch t.location {
	case "Venus":
		fmt.Println("Gas is a bit more expensive on Venus..")
		cost = 30
	case "Mercury":
		fmt.Println("Gas is a bit more expensive on Mercury..")
		cost = 60
	case "Mars":
		fmt.Println("Gase is cheaper on Mars..")
		cost = 2
	case "Earth":
		fmt.Println("Not a bad call to fill up while home..")
		cost = 4
	}
	t.adjustBalance(cost)
}

func (t *Traveler) gasStation() {
	fmt.Print("We are at ", t.location, "'s gas station.\n\n")
	t.calculateFuelNeeded()
	t.calculateFuelCost()
	fmt.Print("Your current balance is ", t.balance, ".\n\n")
}

func (t *Traveler) userInput() bool {

	fmt.Println("Which of the planets would you like to travel to? \nWe are currently on", t.location)
	fmt.Print("Venus, Mercury, Mars or Earth; q to quit: ")
	fmt.Scan(&t.planetChoice)
	// And then liftoff!
	switch t.planetChoice {
	case "Venus":
		t.flyToPlanet()
	case "Mercury":
		t.flyToPlanet()
	case "Mars":
		t.flyToPlanet()
	case "Earth":
		t.takeUsHome()
	case "q":
		return false
	default:
		fmt.Println("Enter something valid please :)")
	}
	fmt.Println("Fuel remaining:", t.fuel)
	if t.balance != 0 || t.fuel > 300000 {
		return true
	}
	fmt.Println("You're out of money!")
	if t.location != "Earth" {
		fmt.Printf("You're stranded on %s, I told you to be careful!", t.location)
	} else {
		fmt.Println("At least you didn't run out of money away from home!")
	}
	return false
}
