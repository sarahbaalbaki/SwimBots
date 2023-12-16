package main

import (
	"fmt"
	"gifhelper"
	"math/rand"
)

func main() {

	rand.Seed(0)

	var isDefault string

	// basic
	var numGen int
	var time float64

	// parameters
	var numInitialBots int
	var numFood int
	var viewRange float64
	var proximity float64
	var foodEnergy float64
	var hungerThreshold float64
	var maximumAge float64
	var foodFrequency int
	var segmentMass float64
	var energyLossFactor float64
	var matingPreference int

	fmt.Println("Welcome to swimbot genepool simulation.")
	fmt.Println("Would you like to simulate the genepool with default parameters? (y/n)")
	fmt.Scan(&isDefault)

	if isDefault == "y" {
		fmt.Println("Simulating genepool with default parameters.")

		numGen = 1000
		fmt.Println("Number of generations: ", numGen)

		time = 1
		fmt.Println("Time interval: ", time)

		numInitialBots = 200
		fmt.Println("Initial number of bots: ", numInitialBots)

		numFood = 5
		fmt.Println("Number of food bits we add every time we add food: ", numFood)

		viewRange = 300
		fmt.Println("View range of a swimbot: ", viewRange)

		proximity = 10
		fmt.Println("Proximity for swimbot to eat or mate:", proximity)

		foodEnergy = 50
		fmt.Println("Energy of the foodBits", foodEnergy)

		hungerThreshold = 50
		fmt.Println("Energy threshold for hungry: ", hungerThreshold)

		maximumAge = 1000
		fmt.Println("Maximum age of a bot: ", maximumAge)

		foodFrequency = 5
		fmt.Println("The frequency of putting in food: ", 5)

		segmentMass = 10.0
		fmt.Println("The mass of each segment:", segmentMass)

		energyLossFactor = 0.0005
		fmt.Println("The energy loss factor is:", energyLossFactor)

		matingPreference = 0
		fmt.Println("The mating prefernce is randomized.")

	} else if isDefault == "n" {

		fmt.Println("Simulating genepool with your own parameters!")

		// numGen
		fmt.Println("How many generations would you like to simulate?")
		fmt.Println("Please input a integer. (The default value is 1000)")
		fmt.Scan(&numGen)

		// time
		fmt.Println("What's the time interval for each gneration?")
		fmt.Println("Please input a float64. (The default value is 1.0)")
		fmt.Scan(&time)

		// numInitialBots
		fmt.Println("How many swimbots would you like to have in the initial pond?")
		fmt.Println("Please input a integer. (The default value is 200)")
		fmt.Scan(&numInitialBots)

		// numFood
		fmt.Println("How many food bits do you want to add everytime?")
		fmt.Println("Please input a integer. (The default value is 5)")
		fmt.Scan(&numFood)

		// viewRange
		fmt.Println("How far do you want a swimbot to see to pick it's goal to swim towards?")
		fmt.Println("Please input a float64. (The default value is 300)")
		fmt.Scan(&viewRange)

		// proximity
		fmt.Println("How close does a swimbot have to get to its goal in order to eat or mate?")
		fmt.Println("Please input a float64. (The default value is 10)")
		fmt.Scan(&proximity)

		// foodEnergy
		fmt.Println("How much energy does a swimbot gain when eating a food?")
		fmt.Println("Please input a float64. (The default value is 50.0)")
		fmt.Scan(&foodEnergy)

		// hungerThreshold
		fmt.Println("What's the hunger threshold of the swimbot?")
		fmt.Println("Please input a float64. (The default value is 50)")
		fmt.Scan(&hungerThreshold)

		// hungerThreshold
		fmt.Println("What's the maximum age of a swimbot?")
		fmt.Println("Please input a float64. (The default value is 1000)")
		fmt.Scan(&maximumAge)

		// foodFrequency
		fmt.Println("How often do you want to throw food into the pond?")
		fmt.Println("Please input a int. (The default value is 5)")
		fmt.Scan(&foodFrequency)

		// foodFrequency
		fmt.Println("What's the basic value for mass of a segment in swimbots?")
		fmt.Println("Please input a float64. (The default value is 10.0)")
		fmt.Scan(&segmentMass)

		// energyLossFactor
		fmt.Println("How much energy was loss when the swimbot swim?")
		fmt.Println("Please input a float64. (The default value is 0.0005)")
		fmt.Scan(&energyLossFactor)

		// matingPreference
		fmt.Println("How should the swimbots pick their mate?")
		fmt.Println("0: randomly choose a mate.")
		fmt.Println("1: choose mate with more segments.")
		fmt.Println("2: choose mate with less segments.")
		fmt.Println("3: choose mate that's faster.")
		fmt.Println("4: choose mate that have similar number of segments.")
		fmt.Println("5: choose mate with similar main segment length.")
		fmt.Println("Please input a integer. (The default value is 0)")
		fmt.Scan(&matingPreference)

		fmt.Println("Number of generations: ", numGen)
		fmt.Println("Time interval: ", time)
		fmt.Println("Initial number of bots: ", numInitialBots)
		fmt.Println("Number of food bits we add every time we add food: ", numFood)
		fmt.Println("View range of a swimbot: ", viewRange)
		fmt.Println("The proximity is:", proximity)
		fmt.Println("The food energy is:", foodEnergy)
		fmt.Println("Energy threshold for hungry: ", hungerThreshold)
		fmt.Println("Maximum age of a bot: ", maximumAge)
		fmt.Println("The frequency of putting in food: ", foodFrequency)
		fmt.Println("The mass of each segment: ", segmentMass)
		fmt.Println("The energy loss factor: ", energyLossFactor)
		fmt.Println("The mating preference: ", matingPreference)

	} else {
		panic("Invalid answer!")
	}

	fmt.Println("Parameters received. Start Simulation!")

	timePoints := SimulatePond(numGen, time, numInitialBots, numFood, viewRange, proximity, foodEnergy, hungerThreshold, maximumAge, foodFrequency, segmentMass, energyLossFactor, matingPreference)
	images := AnimateSystem(timePoints, 2000, 1, 10)
	fmt.Println("Images drawn!")

	// making gif for the simulations
	fmt.Println("Making GIF.")
	gifhelper.ImagesToGIF(images, "pond")
	fmt.Println("Animated GIF produced!")

	fmt.Println("Analyzing result.")
	GenerateAnalysis(timePoints[0], timePoints[len(timePoints)-1], numGen)
	fmt.Println("txt file produced.")
	fmt.Println("Existing normally.")

}
