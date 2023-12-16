package main

import (
	"fmt"
	"math"
	"math/rand"
)

// SimulatePond takes in initialPond, and simulate the artificial pond numGen of times.
func SimulatePond(numGens int, time float64, numInitialBots, numFood int, viewRange float64, proximity float64, foodEnergy, hungerThreshold float64, maximumAge float64, foodFrequency int, segmentMass, energyLostFactor float64, matingPreference int) []*Pond {
	// Create an initialized pond with specified number of bots
	initialPond := InitializePond(numInitialBots, segmentMass)
	timePoints := make([]*Pond, numGens+1)
	timePoints[0] = initialPond
	//now range over the number of generations and update the pond each time
	for i := 1; i <= numGens; i++ {
		// fmt.Println("generation", i)
		timePoints[i] = UpdatePond(timePoints[i-1], time, i,  numFood, viewRange, proximity, foodEnergy, hungerThreshold, maximumAge, foodFrequency, segmentMass, energyLostFactor, matingPreference)
	}
	return timePoints
}

// UpdatePond update the pond to a new time point
func UpdatePond(oldPond *Pond, time float64, numGen, numFood int, viewRange float64, proximity float64, foodEnergy, hungerThreshold float64, maximumAge float64, foodFrequency int, segmentMass, energyLossFactor float64, matingPreference int) *Pond {
	// create a new Pond
	newPond := CopyPond(oldPond)

	for i := range newPond.swimbots {
		// if the bot already died, we skip updating the bot
		if newPond.swimbots[i] == nil {
			continue
		}
		// Set the goal for all the living bots in the pond
		newPond.SetGoal(i, oldPond, viewRange, hungerThreshold, matingPreference)
		// update the velocity and position
		newPond.swimbots[i].UpdateVelocity(oldPond, energyLossFactor)
		newPond.swimbots[i].UpdatePosition(time) // & ENERGY
		// update age
		newPond.swimbots[i].age += 1
		// if a bot's energy reaches 0, kill the bot!
		if newPond.swimbots[i].energy <= 0 || newPond.swimbots[i].age >= maximumAge {
			newPond.swimbots[i] = nil
		}
	}
	// determine whether the bots can eat or mate
	newPond.EatOrMate(proximity, foodEnergy, segmentMass)
	// add food when we reach "FoodFrequency"
	newPond.AddFood(numGen, numFood, foodFrequency)
	return newPond
}

// EatOrMate let the swimbot eat or mates at this generation
func (pond *Pond) EatOrMate(proximity float64, foodEnergy, segmentMass float64) {
	// make a slice of swimbots that has already mated for this generation
	// because we don't want them to give twins or give two bots in one generation
	alreadyGotLucky := make([]bool, len(pond.swimbots))

	// range through all the swimbots
	for i := range pond.swimbots {
		// determine only if the swimbot is not nil
		if pond.swimbots[i] != nil {
			// if the swimbot doesn't have a visible goal we skip
			if pond.swimbots[i].goal.index == -1 {
				continue
			} else if pond.swimbots[i].goal.isBot { // the goal of the bot isBot
				// There are four scenarios that we need to satisfy in order to mate
				// 1. The goal swimbot still exist (not nil)
				// 2. It's within proximity
				// 3. The swimbot haven't mate in this round
				// 4. The goal swimbot haven't mate in this round
				if pond.swimbots[pond.swimbots[i].goal.index] != nil && pond.swimbots[i].GoalDistance(pond) <= proximity && alreadyGotLucky[i] == false && alreadyGotLucky[pond.swimbots[i].goal.index] == false {
					childIndex := len(pond.swimbots)
					// Generate a child through mating
					child := pond.Mating(i, pond.swimbots[i].goal.index, childIndex, segmentMass)
					pond.swimbots = append(pond.swimbots, child)
					// record the mating swimbots
					alreadyGotLucky[i] = true
					alreadyGotLucky[pond.swimbots[i].goal.index] = true
				}
			} else { // the goal of the bot is food
				// if the foodbit is not nil
				if pond.foodBits[pond.swimbots[i].goal.index] != nil && pond.swimbots[i].GoalDistance(pond) <= proximity {
					pond.swimbots[i].energy += foodEnergy
					// we are using integer as an goal
					pond.foodBits[pond.swimbots[i].goal.index] = nil
				}
			}
		}
	}
}

// InitializePond generate randomized swimbots and foodbits at random positions
func InitializePond(numBots int, segmentMass float64) *Pond {
	var p Pond
	p.width = 6000
	initialEnergy := 75.0

	// Initialize swimbots and append them to the slice
	for i := 0; i < numBots; i++ {
		p.swimbots = append(p.swimbots, InitializeSwimbot(initialEnergy, segmentMass))
		p.swimbots[i].family = append(p.swimbots[i].family, i)
	}

	// generate food
	for i := 0; i < 400; i++ {
		var f Food
		// random positions whithin the slightly bigger square from origin
		f.position.x = float64(rand.Intn(5000) + 500)
		f.position.y = float64(rand.Intn(5000) + 500)
		p.foodBits = append(p.foodBits, &f)
	}
	return &p
}

// Mating takes in the index of two swimbots, a index of the child and produce a offspring
func (pond *Pond) Mating(s1, s2 int, childIndex int, segmentMass float64) *Swimbot {
	// calculate the energy for the children

	bot1 := pond.swimbots[s1]
	bot2 := pond.swimbots[s2]

	childEnergy := bot1.energy/2 + bot2.energy/2
	// update the parent's energy level
	bot1.energy *= 0.5
	bot2.energy *= 0.5

	// fmt.Println("S1's main segment genome:", pond.swimbots[s1].segGenes[0])
	// fmt.Println("S2's main segment genome:", pond.swimbots[s2].segGenes[0])

	child := GenerateChild(bot1, bot2, childEnergy, segmentMass)

	// we append the two parents to the family of the child
	child.family = append(child.family, s1)
	child.family = append(child.family, s2)
	child.family = append(child.family, childIndex)

	// append the child to the family of the parents
	bot1.family = append(bot1.family, childIndex)
	bot2.family = append(bot2.family, childIndex)
	// fmt.Printf("A new child%d was given birth by parent %d and %d,cheers!\n", childIndex, s1, s2)
	// fmt.Println("S1's main segment genome:", pond.swimbots[s1].segGenes[0])
	// fmt.Println("S2's main segment genome:", pond.swimbots[s2].segGenes[0])
	// fmt.Println("Children main segment genome:", child.segGenes[0])

	return child
}

// GenerateChild generate a children bot based on its parents' genome and return its pointer
func GenerateChild(s1, s2 *Swimbot, energy float64, segmentMass float64) *Swimbot {
	var child Swimbot
	// set age and energy
	child.age = 0
	child.energy = energy

	child.position.x = (s1.position.x + s2.position.x) * 0.5
	child.position.y = (s1.position.y + s2.position.y) * 0.5

	// we will update the velocity and acceleration in the next round
	// velocity= s1.velocity + s2.velocity/2.0
	// acceleration= s1.acceleration + s2.acceleration/2.0

	// Generate the genes for the child
	child.segGenes, child.botGene = GenerateOffspringGenome(s1, s2)

	// randomize the initial velocity of the child
	child.velocity.x = (rand.Float64()-0.5) * child.botGene.translationalMovement
	child.velocity.y = (rand.Float64()-0.5) * child.botGene.translationalMovement

	// calculate the mass of the child
	child.mass = segmentMass * float64(child.botGene.numSegments)
	// build the segments for the bot
	child.BuildSegments()
	return &child
}

// BuildSegments builds the body (the arrangement of segments) of a bot based on its expressing segmentgenes
func (bot *Swimbot) BuildSegments() {
	var mainSeg Segment
	bot.mainSegment = &mainSeg
	//First, decide which segment will be the main segment of the bot
	bot.mainSegment.index = 0

	// The position of the main segment is the same as the bot's position
	bot.mainSegment.position.x = bot.position.x
	bot.mainSegment.position.y = bot.position.y

	// calculat the angle of the mainsegment
	angle := math.Atan(bot.velocity.y/bot.velocity.x)
	if bot.velocity.x < 0 {
		angle += math.Pi
	}
	bot.mainSegment.angle = angle

	for i := 1; i < bot.botGene.numSegments; i++ {
		//currentSegment is a pointer of the current segment of the bot we are looking at
		//it always starts with the main segment
		currentSegment := bot.mainSegment
		//for each segment to be added, we need to determine which existing segment to attach onto by flipping coins.
		for currentSegment.subSegments != nil { // have subSegments
			//First, flip a coin.
			FlipaCoin := rand.Intn(2)
			//If coin == 1, then we will continue looking at the subsegments of the current segment.
			if FlipaCoin == 1 {
				//randomly choose a subsegment to continue
				segIndex := rand.Intn(len(currentSegment.subSegments))
				currentSegment = currentSegment.subSegments[segIndex]
			} else {
				//if coin == 0, we will stop at the current segment
				break
			}
			//or when the currentsegment does not have subsegments, we will stop at the current segment.
		}
		// doesn't have subSegments
		var newSegment Segment
		//the index of the segment corresponds to the sepecific segmentgene
		newSegment.index = i
		newSegment.CalculateSegmentPosition(currentSegment, bot.segGenes)

		//attach the segment to the currentsegment
		currentSegment.subSegments = append(currentSegment.subSegments, &newSegment)
	}
}

// CalculateSegmentPosition takes in an old segment and calculate the position of the current segment based on the segment it is attached to
func (newSeg *Segment) CalculateSegmentPosition(currentSegment *Segment, segGene []SegmentGene)  {

	// calculate the length of the current segment
	l := segGene[newSeg.index][4]
	// calculate the angle of the current segment by adding angle to parent to the angle of the previous segment
	newSeg.angle = currentSegment.angle + segGene[newSeg.index][3]
	// calculate the center of the new segments
	newSeg.position.x = currentSegment.position.x + 0.5*segGene[currentSegment.index][4]*math.Cos(math.Pi-currentSegment.angle) + 0.5*l*math.Cos(math.Pi-newSeg.angle)
	newSeg.position.y = currentSegment.position.y + 0.5*segGene[currentSegment.index][4]*math.Sin(currentSegment.angle) + 0.5*l*math.Sin(newSeg.angle)

	if math.IsNaN(newSeg.position.x) {
		panic("The newSegment position y is not a number.")
	}
	if math.IsNaN(newSeg.position.y) {
		panic("The newSegment position x is not a number.")
	}
}

// func GenerateOffspringGenome(botGene1, botGene2 CommonGene, segGene1, segGene2 SegmentGenes) (SegmentGene, CommonGene){
func GenerateOffspringGenome(s1, s2 *Swimbot) ([]SegmentGene, CommonGene) {

	// do for the three genes:
	//   choose random number between 0 and 1
	//   if 0-> take from bot1
	//   if 1-> take from bot2

	// for the three common gene we are choosing randomly from the parent
	var offspringCommonGene CommonGene

	num := rand.Intn(2)
	if num == 0 {
		offspringCommonGene.angularMovement = s1.botGene.angularMovement
	} else {
		offspringCommonGene.angularMovement = s2.botGene.angularMovement
	}

	num = rand.Intn(2)
	if num == 0 {
		offspringCommonGene.translationalMovement = s1.botGene.translationalMovement
	} else {
		offspringCommonGene.translationalMovement = s2.botGene.translationalMovement
	}

	num = rand.Intn(2)
	if num == 0 {
		offspringCommonGene.numSegments = s1.botGene.numSegments
	} else {
		offspringCommonGene.numSegments = s2.botGene.numSegments
	}

	// Each of the SegmentGenes is generated by combining the corresponding segmentgenes of parents.
	offspringSegmentGene := make([]SegmentGene, 8)
	for i := range offspringSegmentGene {
		//for each segmentgene, a random crossoverpoint is generated.
		crosspoint := rand.Intn(6)
		//The 0-crosspoint part of the segmentgene will be inherited from one parent and the rest from the other parent.
		offspringSegmentGene[i] = GenerateSegmentGene(s1.segGenes[i], s2.segGenes[i], crosspoint)
	}

	return offspringSegmentGene, offspringCommonGene
}

// GenerateSegmentGene takes in two gene and perform a crossover of genome
func GenerateSegmentGene(s1Gene, s2Gene SegmentGene, crosspoint int) SegmentGene {
	s1GeneCopy := make(SegmentGene, len(s1Gene))
	s2GeneCopy := make(SegmentGene, len(s2Gene))
	// In order to obtain the deep copy
	copy(s1GeneCopy, s1Gene)
	copy(s2GeneCopy, s2Gene)
	// Append the two genes back to the segment
	offspringSegmentGene := append(s1GeneCopy[:crosspoint], s2GeneCopy[crosspoint:]...)
	return offspringSegmentGene
}

// UpdateVelocity updates the velocity f the bot
func (bot *Swimbot) UpdateVelocity(pond *Pond, energyLossFactor float64) {
	// if the goal is -1, it couldn't find a goal
	// keep swimming towards the same direction
	if bot.goal.index != -1 {
		// swim towards its goal
		var deltax float64
		var deltay float64

		switch bot.goal.isBot {
		// if the goal of the bot is a bot
		case true:
			// calculate the distance against the bot
			deltax = pond.swimbots[bot.goal.index].position.x - bot.position.x
			deltay = pond.swimbots[bot.goal.index].position.y - bot.position.y
		// if the goal of the bot is a food
		case false:
			// calculate the distance against the food
			deltax = pond.foodBits[bot.goal.index].position.x - bot.position.x
			deltay = pond.foodBits[bot.goal.index].position.y - bot.position.y

		}
		// we calculate new and old angle by calculating acosine
		newangle := math.Acos(deltax / bot.GoalDistance(pond))
		oldangle := math.Acos(bot.velocity.x / bot.botGene.translationalMovement)

		// handle if the newAngle return NaN
		if math.IsNaN(newangle) {
			fmt.Println("goal is", bot.goal.index)
			fmt.Println("The bot position are:", bot.position, pond.swimbots[bot.goal.index].position)
			panic("New angle is not a number!")
		}
		if math.IsNaN(oldangle) {
			fmt.Println("Bot velocity is", bot.velocity.x)
			fmt.Println("Translational movement is ", bot.botGene.translationalMovement)
			fmt.Println(oldangle)
			panic("Old angle is not a number!")
		}
		// Restrict the turning angle with angularMovement gene
		if math.Abs(newangle-oldangle) > bot.botGene.angularMovement {
			if newangle-oldangle <= 0 {
				newangle = oldangle - bot.botGene.angularMovement
			} else {
				newangle = oldangle + bot.botGene.angularMovement
			}
		}

		bot.mainSegment.angle = newangle

		// Calculate the velocity using the new restricted angle
		bot.velocity.x = bot.botGene.translationalMovement * math.Cos(newangle)
		if deltay < 0 {
			bot.velocity.y = -bot.botGene.translationalMovement * math.Sin(newangle)
		} else {
			bot.velocity.y = bot.botGene.translationalMovement * math.Sin(newangle)

		}

	} else {
		if bot.position.x >= pond.width || bot.position.x <= 0 {
			bot.velocity.x = -bot.velocity.x
		}
		if bot.position.y >= pond.width || bot.position.y <= 0{
			bot.velocity.y = -bot.velocity.y
		}
	}
	// decrease the energy according to the speed.
	speed := math.Sqrt(math.Pow(bot.velocity.x, 2)+ math.Pow(bot.velocity.y, 2))
	// the loss of energy is porportioned to the square of speed and bot's mass
	bot.energy -= energyLossFactor*speed*speed*bot.mass
}

// UpdatePosition update the position of a swimbot based on its velocity
func (bot *Swimbot) UpdatePosition(time float64) {
	bot.position.x += bot.velocity.x * time
	bot.position.y += bot.velocity.y * time

	bot.mainSegment.position.x = bot.position.x
	bot.mainSegment.position.y = bot.position.y
	// handle the case where velocity is updated to not a number
	if math.IsNaN(bot.velocity.x){
		fmt.Println("Velocity == NaN")
	}
	if math.IsNaN(bot.velocity.y){
		fmt.Println("Velocity == NaN")
	}

	for i := range bot.mainSegment.subSegments {
		bot.mainSegment.subSegments[i].UpdateSegmentPosition(bot.mainSegment, bot.segGenes)
	}
}

func (segment *Segment)UpdateSegmentPosition(prevSegment *Segment, segGene []SegmentGene) {
	segment.CalculateSegmentPosition(prevSegment, segGene)
	if segment.subSegments != nil {
		for i := range segment.subSegments {
			segment.subSegments[i].UpdateSegmentPosition(segment, segGene)
		}
	}
}

// InitializeSwimbot generate a swimbot with randomize genome at a random position and
func InitializeSwimbot(initialEnergy, segmentMass float64) *Swimbot {
	var bot Swimbot

	bot.age = 0.0 // couldn't this be an int?
	bot.energy = initialEnergy

	bot.position.x = float64(rand.Intn(4000) + 1000) // random positions to each swimbot within the given radius
	bot.position.y = float64(rand.Intn(4000) + 1000)

	// generate genome for the bot
	bot.botGene, bot.segGenes = RandomGenome()


	angle := rand.Float64()*2*math.Pi
	bot.velocity.x = math.Cos(angle) * bot.botGene.translationalMovement
	bot.velocity.y = math.Sin(angle) * bot.botGene.translationalMovement

	// mass is associate with the number of segments
	bot.mass = segmentMass * float64(bot.botGene.numSegments)

	bot.BuildSegments()

	return &bot
}

// RandomGenome generates a random genome for the initialization of swimbots
func RandomGenome() (CommonGene, []SegmentGene) {
	var common CommonGene
	// we have to multiply this by something when we calculate velocity, maybe 45 degree?
	common.angularMovement = rand.Float64() * math.Pi / 4.0
	// let's time 10 when we scale it
	common.translationalMovement = rand.Float64() * 10
	// should generate values 2 to 8
	common.numSegments = rand.Intn(7) + 2
	// each bot have 8 segement gene
	segGenes := make([]SegmentGene, 8)
	// each segment gene have 6 traits
	for i := 0; i < 8; i++ {
		segGenes[i] = make(SegmentGene, 6)
		redInt := rand.Intn(256)
		greenInt := rand.Intn(256)
		blueInt := rand.Intn(256)

		segGenes[i][0] = float64(redInt)
		segGenes[i][1] = float64(greenInt)
		segGenes[i][2] = float64(blueInt)

		// angleToParent
		segGenes[i][3] = rand.Float64()*math.Pi - (math.Pi / 2.0) // should take on values from -(pi)/2 to +(pi)/2 -- using radians b/c trig functions take that
		// length
		segGenes[i][4] = (rand.Float64() * 15.0) + 5.0 // should take on values from 5.0 to 20.0
		// width
		segGenes[i][5] = (rand.Float64() * 3.7) + 0.3 // should take on values from 0.3 to 4.0
	}
	return common, segGenes
}

// SetGoal takes a pond and sets a bot's goal for the current timestep based on its goal from the prior timestep and its current energy/state.
func (newPond *Pond) SetGoal(i int, oldPond *Pond, viewRange float64, hungerThreshold float64, matingPreference int) {
	// in many cases, bot should simply keep the same goal it had from the prior timestep
	needsNewGoal := false
	bot := newPond.swimbots[i]
	// check cases in which bot should have its goal updated, changing needsNewGoal to true if applicable
	if bot.goal.index == -1 {
		// this occurs if an execution of FindNewGoal() function from prior timestep fails to find appropriate goal within bot's view range
		needsNewGoal = true
	} else {
		// if the bot's prior goal doesn't match its current state, it will need to find a new goal
		if (bot.goal.isBot && bot.energy < hungerThreshold) || (!bot.goal.isBot && bot.energy >= hungerThreshold) {
			needsNewGoal = true
		}
		// if the food bit/bot that is its goal no longer exists or has moved out of view, it will also need a new goal
		// is the goal a bot?
		if bot.goal.isBot {
			currentGoalMate := oldPond.swimbots[bot.goal.index]
			if currentGoalMate == nil || bot.GoalDistance(oldPond) > viewRange {
				needsNewGoal = true
			}
		} else {
			// goal is a food bit
			currentGoalFood := newPond.foodBits[bot.goal.index]
			// find the new goal if the food is gone or out of range
			if currentGoalFood == nil || bot.GoalDistance(oldPond) > viewRange {
				needsNewGoal = true
			}
		}

	}
	// if any of the prior conditions were triggered, bot will update its goal; otherwise it simply retains the prior one
	if needsNewGoal {
		bot.goal = bot.FindNewGoal(oldPond, viewRange, hungerThreshold, matingPreference)
	}
}

// FindNewGoal takes a pond and returns a new goal for the swimbot based on its current state/energy. It returns a goal with index field equal to -1 if no objects of the appropriate type are currently within the bot's view range.
func (bot *Swimbot) FindNewGoal(pond *Pond, viewRange, hungerThreshold float64, matingPreference int) Goal {
	var newGoal Goal
	// bot is hungry, will pursue its closest food bit
	if bot.energy < hungerThreshold {
		newGoal.isBot = false
		// will be updated if any foodbits within view are found
		closestFoodIndex := -1

		var shortestDist float64

		// range through all food bits in pond, keeping track of which is closest (within bot's view)
		for i := range pond.foodBits {
			if pond.foodBits[i] != nil {
				dist := bot.DistanceToFood(pond.foodBits[i])
				// if the food is closer update the index and distance
				if dist < viewRange && (closestFoodIndex == -1 || dist < shortestDist) {
					closestFoodIndex = i
					shortestDist = dist
				}
			}
		}
		// if no food bits were found within bot's view, this will still hold a value of -1
		newGoal.index = closestFoodIndex

	} else { // bot is not hungry, will pursue another bot that is within view and not part of its family
		newGoal.isBot = true

		suitableBotIndices := make([]int, 0)

		for i := range pond.swimbots {
			// if it's not himself and not nil
			if pond.swimbots[i] != bot && pond.swimbots[i] != nil {
				potentialMate := pond.swimbots[i]
				// choose bots that is not related to the current bot
				if bot.DistanceToSwimbot(potentialMate) <= viewRange && !bot.RelatedTo(i) {
					suitableBotIndices = append(suitableBotIndices, i)
				}
			}
		}
		// fmt.Println("bot's potentialMate", suitableBotIndices)
		if len(suitableBotIndices) == 0 {
			newGoal.index = -1
		} else {
			// choose the mate based on mating preferences
			// 0: choose a mate randomly
			// 1: choose swimbot with more segments.
			// 2: choose swimbot with less segments.
			// 3: choose swimbot that's faster
			// 4: similar number of segments.
			// 5: similar length of mainSegment

			// first set the goal to the first in the potential mating list
			newGoal.index = suitableBotIndices[0]
			// based on the matingPreferences, we choose the mate differently
			if matingPreference == 0 {
				newGoal.index = suitableBotIndices[rand.Intn(len(suitableBotIndices))]
			} else if matingPreference == 1 {
				for i := range suitableBotIndices {
					if pond.swimbots[suitableBotIndices[i]].botGene.numSegments >= pond.swimbots[newGoal.index].botGene.numSegments {
						newGoal.index = suitableBotIndices[i]
					}
				}
			} else if matingPreference == 2 {
				for i := range suitableBotIndices {
					if pond.swimbots[suitableBotIndices[i]].botGene.numSegments <= pond.swimbots[newGoal.index].botGene.numSegments {
						newGoal.index = suitableBotIndices[i]
					}
				}
			} else if matingPreference == 3 {
				for i := range suitableBotIndices {
					// if the target bot is faster, we update the goal
					if pond.swimbots[suitableBotIndices[i]].botGene.translationalMovement >= pond.swimbots[newGoal.index].botGene.translationalMovement {
						newGoal.index = suitableBotIndices[i]
					}
				}
			} else if matingPreference == 4 {
				newGoal.index = suitableBotIndices[0]
				// calculate the difference between the mainsegment of the current swimbot and the target swimbot
				segmentDiff := math.Abs(float64(pond.swimbots[newGoal.index].botGene.numSegments - bot.botGene.numSegments))
				// range through all the suitable bots
				for i := 1; i < len(suitableBotIndices); i++ {
					// if the new segment difference is smaller than the current one, we replace the target with new index
					if segmentDiff >= math.Abs(float64(pond.swimbots[suitableBotIndices[i]].botGene.numSegments - bot.botGene.numSegments)) {
						newGoal.index = suitableBotIndices[i]
						segmentDiff = math.Abs(float64(pond.swimbots[suitableBotIndices[i]].botGene.numSegments - bot.botGene.numSegments))
					}
				}
			} else if matingPreference == 5 {
				newGoal.index = suitableBotIndices[0]
				lengthDiff := math.Abs(float64(pond.swimbots[newGoal.index].segGenes[0][4] - bot.segGenes[0][4]))
				for i := 1; i < len(suitableBotIndices); i++ {
					// if the difference in length is smaller, we replace it with the new one
					if lengthDiff >= math.Abs(float64(pond.swimbots[suitableBotIndices[i]].segGenes[0][4] - bot.segGenes[0][4])) {
						newGoal.index = suitableBotIndices[i]
						lengthDiff = math.Abs(float64(pond.swimbots[suitableBotIndices[i]].segGenes[0][4] - bot.segGenes[0][4]))
					}
				}
			}
		}
	}

	return newGoal
}

// DistanceToSwimbot takes a pointer to a swimbot and returns the distance between that swimbot and the one calling the function.
func (bot *Swimbot) DistanceToSwimbot(otherBot *Swimbot) float64 {
	// handle the case if the bot or its goal's distance is not a number
	if math.IsNaN(bot.position.x) {
		fmt.Println(bot)
		panic("Bot position is NaN")
	}
	if math.IsNaN(otherBot.position.x) {
		fmt.Println(otherBot)
		panic("Other Bot position is NaN")
	}
	// calculate the distance
	deltaX := otherBot.position.x - bot.position.x
	deltaY := otherBot.position.y - bot.position.y
	distance := math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))

	return distance
}

// DistanceToFood takes a pointer to a food bit and returns the distance between the food bit and the swimbot caling the function.
func (bot *Swimbot) DistanceToFood(foodBit *Food) float64 {
	deltaX := foodBit.position.x - bot.position.x
	deltaY := foodBit.position.y - bot.position.y

	return math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))
}

// RelatedTo takes a pointer to a swimbot and returns whether it is part of the family of the swimbot calling the function.
func (bot *Swimbot) RelatedTo(botIndex int) bool {
	for i := range bot.family {
		// fmt.Println("bot's family, otherbot")
		// fmt.Println(bot.family[i], botIndex)
		// if the index is the same then they are related
		if bot.family[i] == botIndex {
			return true
		}
	}
	return false
}

// GoalDistance takes a pond and returns the distance between the bot calling the function and its current goal.
func (bot *Swimbot) GoalDistance(p *Pond) float64 {
	var dist float64

	if bot.goal.isBot {
		dist = bot.DistanceToSwimbot(p.swimbots[bot.goal.index])
		// handle the case if the distance is not a number
		if math.IsNaN(dist) {
			fmt.Println("Distance to swimbot is NaN")
		}
	} else {
		dist = bot.DistanceToFood(p.foodBits[bot.goal.index])
		// handle the case if the distance is not a number
		if math.IsNaN(dist) {
			fmt.Println("Distance to food is NaN")
		}
	}

	return dist
}

// CopySegment copy the order of segments of a bot to a copied swimbot
func CopySegmentTree(seg Segment) *Segment {
	var mainSegNew Segment
	mainSegNew.RecursiveCopy(&seg)
	return &mainSegNew
}

// RecursiveCopy copy the subsegments of the segments recursively
func (new *Segment) RecursiveCopy(old *Segment) {
	new.CopySegment(old)
	// if the old have subsegment
	if old.subSegments != nil {
		// range through the subsegments and copy all the fields
		for i := range old.subSegments {
			var newSubSeg Segment
			new.subSegments = append(new.subSegments, &newSubSeg)
			new.subSegments[i].RecursiveCopy(old.subSegments[i])
		}
	}
}

// CopySegment copy all the fields of a segment besides the subsegments
func (SegNew *Segment) CopySegment(seg *Segment) {
	SegNew.position.x = seg.position.x
	SegNew.position.y = seg.position.y
	SegNew.angle = seg.angle
	SegNew.index = seg.index
}

// AddFood add new food to the pond
func (pond *Pond) AddFood(numGen int, numFood int, foodFrequency int) {
	if numGen%foodFrequency == 0 { // we are at a timestep that is a multiple of 20, add food at this step
		for i := 0; i < numFood; i++ { // add food bits at random positions on the board
			var f Food
			f.position.x = float64(rand.Intn(5000) + 500)
			f.position.y = float64(rand.Intn(5000) + 500)
			pond.foodBits = append(pond.foodBits, &f)
		}
	}
}

// CopyPond copy all the fields in the pond to a new pond
func CopyPond(oldPond *Pond) *Pond {
	var newPond Pond

	newPond.width = oldPond.width
	numBots := len(oldPond.swimbots)
	newPond.swimbots = make([]*Swimbot, numBots)

	for i := range oldPond.swimbots {
		if oldPond.swimbots[i] == nil {
			newPond.swimbots[i] = nil
		} else {
			// copy every field in at the bottom
			var SwimbotNew Swimbot
			SwimbotNew.goal = oldPond.swimbots[i].goal
			SwimbotNew.age = oldPond.swimbots[i].age
			SwimbotNew.energy = oldPond.swimbots[i].energy
			SwimbotNew.position.x = oldPond.swimbots[i].position.x
			SwimbotNew.position.y = oldPond.swimbots[i].position.y
			SwimbotNew.velocity.x = oldPond.swimbots[i].velocity.x
			SwimbotNew.velocity.y = oldPond.swimbots[i].velocity.y
			SwimbotNew.acceleration.x = oldPond.swimbots[i].acceleration.x
			SwimbotNew.acceleration.y = oldPond.swimbots[i].acceleration.y
			SwimbotNew.mass = oldPond.swimbots[i].mass

			// CHECK FAMILY POINTERS!!!!!!!!!!!!!!!
			fam := make([]int, len(oldPond.swimbots[i].family))
			fam = oldPond.swimbots[i].family
			SwimbotNew.family = fam

			var newCommonGene CommonGene
			newCommonGene.angularMovement = oldPond.swimbots[i].botGene.angularMovement
			newCommonGene.translationalMovement = oldPond.swimbots[i].botGene.translationalMovement
			newCommonGene.numSegments = oldPond.swimbots[i].botGene.numSegments
			SwimbotNew.botGene = newCommonGene

			segGenesNew := make([]SegmentGene, 8)
			for k := range segGenesNew {
				segGenesNew[k] = make(SegmentGene, 6)
				for j := 0; j < 6; j++ {

					segGenesNew[k][j] = oldPond.swimbots[i].segGenes[k][j]
				}
			}

			SwimbotNew.segGenes = segGenesNew
			// We need to copy the subsegments recursively
			SwimbotNew.mainSegment = CopySegmentTree(*oldPond.swimbots[i].mainSegment)
			newPond.swimbots[i] = &SwimbotNew
		}
	}

	numFood := len(oldPond.foodBits)
	newPond.foodBits = make([]*Food, numFood)
	copy(newPond.foodBits, oldPond.foodBits)

	// for k := range oldPond.foodBits {
	// 	// var newFoodBit Food
	// 	// newFoodBit.position.x = oldPond.foodBits[k].position.x
	// 	// newFoodBit.position.y = oldPond.foodBits[k].position.y
	// 	// newPond.foodBits[k] = &newFoodBit
	// 	newPond.foodBits[k] = oldPond.foodBits[k]
	// }

	return &newPond
}
