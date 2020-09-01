package main

import "fmt"

/*
*	This is a fun programming project to help learn Golang
*	I implement conditional statements such as; if/else and switch
*	Also use the fmt Scan, Print, Printf, Println
*	Functions heavy project
*	The functions have mulitple parameters and some return multiple variables
 */

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("Fuel remaining:", fuel)
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

//Function that will return current balance
func getBalance(balance int) {
	fmt.Println("Your current balance is", balance, "\n")
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Print("Welcome to the planet ", planet, ", please be cautious as we are far from home!\n")
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
	fmt.Println("Lets go to the gas station.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, currentPlanet string, fuel int, balance int) (int, string, int) {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		currentPlanet = planet
		fuelRemaining -= fuelCost
	} else {
		cantFly()
		balance, fuel = gasStation(fuel, currentPlanet, balance)
	}
	return fuelRemaining, currentPlanet, balance
}

//Create a function takeUsHome() here
func takeUsHome(currentPlanet string, fuel int, balance int) (int, string, int) {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(currentPlanet)
	if fuel >= fuelCost {
		switch currentPlanet {
		case "Venus":
			fuelRemaining -= fuelCost
			currentPlanet = "Home"
			fmt.Println("Welcome Back Home :)")
		case "Mercury":
			fuelRemaining -= fuelCost
			currentPlanet = "Home"
			fmt.Println("Welcome Back Home :)")
		case "Mars":
			fuelRemaining -= fuelCost
			currentPlanet = "Home"
			fmt.Println("Welcome Back Home :)")
		case "Home":
			fmt.Println("You are already home, silly!")
		default:
			fmt.Println("I think we're lost")
		}

	} else {
		cantFly()
		balance, fuel = gasStation(fuel, currentPlanet, balance)
	}

	return fuelRemaining, currentPlanet, balance
}
func adjustBalance(currentPlanet string, cost int, fuel int, fuelNeeded int, balance int) (int, int) {
	cost *= fuelNeeded
	if cost > balance {
		fmt.Println("You can't afford to fill up the tank.. yikes.")
	} else {
		//fmt.Printf("You just purchased $%v worth of gas.\n", cost)
		fmt.Printf("You bought %v gallons of gas for $%v.\n\n", fuelNeeded, cost)
		balance -= cost
		fuel = 1000000
	}
	return balance, fuel
}
func calculateFuelNeeded(currentPlanet string, fuel int) int {
	// We get 1000000 per tank so lets get our gallons per tank assuming efficiency of 100mpg
	gallons := 1000000 / 1000
	currentGallons := fuel / 1000
	gallons -= currentGallons
	return gallons
}
func calculateFuelCost(currentPlanet string, fuel int, fuelNeeded int, balance int) (int, int) {
	var cost int
	switch currentPlanet {
	case "Venus":
		fmt.Println("Gas is a bit more expensive on Venus..")
		cost = 30
		balance, fuel = adjustBalance(currentPlanet, cost, fuel, fuelNeeded, balance)
	case "Mercury":
		fmt.Println("Gas is a bit more expensive on Mercury..")
		cost = 60
		balance, fuel = adjustBalance(currentPlanet, cost, fuel, fuelNeeded, balance)
	case "Mars":
		fmt.Println("Gase is cheaper on Mars..")
		cost = 2
		balance, fuel = adjustBalance(currentPlanet, cost, fuel, fuelNeeded, balance)
	case "Home":
		fmt.Println("Not a bad call to fill up while home..")
		cost = 4
		balance, fuel = adjustBalance(currentPlanet, cost, fuel, fuelNeeded, balance)
	}
	return balance, fuel
}

func gasStation(fuel int, currentPlanet string, balance int) (int, int) {
	var fuelNeeded int
	fmt.Print("We are at ", currentPlanet, "'s gas station.\n\n")
	fuelNeeded = calculateFuelNeeded(currentPlanet, fuel)
	balance, fuel = calculateFuelCost(currentPlanet, fuel, fuelNeeded, balance)
	getBalance(balance)
	return balance, fuel
}

func userInput(fuel int, planetChoice string, currentPlanet string, actualInput bool, balance int) (bool, string, int, int) {

	fmt.Println("Which of the planets would you like to travel to? \nWe are currently at", currentPlanet)
	fmt.Print("Venus, Mercury, Mars or Home; q to quit: ")
	fmt.Scan(&planetChoice)
	// And then liftoff!
	switch planetChoice {
	case "Venus":
		actualInput = true
		fuel, currentPlanet, balance = flyToPlanet(planetChoice, currentPlanet, fuel, balance)
		fuelGauge(fuel)
	case "Mercury":
		actualInput = true
		fuel, currentPlanet, balance = flyToPlanet(planetChoice, currentPlanet, fuel, balance)
		fuelGauge(fuel)
	case "Mars":
		actualInput = true
		fuel, currentPlanet, balance = flyToPlanet(planetChoice, currentPlanet, fuel, balance)
		fuelGauge(fuel)
	case "Home":
		actualInput = true
		fuel, currentPlanet, balance = takeUsHome(currentPlanet, fuel, balance)
		fuelGauge(fuel)
	default:
		actualInput = false
	}
	return actualInput, currentPlanet, fuel, balance

}

func main() {
	// Create `planetChoice` and `fuel`
	var fuel int
	var planetChoice string
	var balance int
	currentPlanet := "Home"
	fmt.Println("Welcome to Interstellar Travel Agency!")
	fmt.Println("This travel agency can take you to a few different planets!\n")
	fmt.Print("How much money do you have?: ")
	fmt.Scan(&balance)

	if balance >= 100000 {
		fmt.Println("Wow! That's great.")
		fmt.Println("You can afford the first tank of gas.")
		balance, fuel = gasStation(fuel, currentPlanet, balance)
	} else {
		fmt.Println("Oof. Broke boi.")
		fmt.Println("Don't worry, the first tank of gas is on us.")
		fuel = 1000000
	}
	//One way to use a do-while loop in Go syntax
	actualInput := true
	for ok := true; ok; ok = actualInput {
		actualInput, currentPlanet, fuel, balance = userInput(fuel, planetChoice, currentPlanet, actualInput, balance)
	}
	fmt.Println("Thanks so much for coming!")
	fmt.Print("Your balance after this vacation is $", balance, "\n\n")
}
