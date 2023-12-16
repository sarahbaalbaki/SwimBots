package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCopySegmentTree(t *testing.T) {
	type test struct {
		inputSegtree  *Segment
		outputSegtree *Segment
	}

	inputDirectory := "tests/CopySegmentTree/input/"
	outputDirectory := "tests/CopySegmentTree/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))
	//create an array of tests
	tests := make([]test, len(inputFiles))

	//range through the input and output files and set the test values
	for i := range inputFiles {
		tests[i].inputSegtree = ReadSegmentTreeFromFile(inputDirectory, inputFiles[i])
		tests[i].outputSegtree = ReadSegmentTreeFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcome := CopySegmentTree(*test.inputSegtree)
		//check if CopySegmentTree copies the whole segment tree correctly
		if !SegTreeIsTheSame(outcome, test.outputSegtree) {
			t.Errorf("Error! For input test dataset %d your function failed", i)
		} else {
			fmt.Printf("For input test dataset %d :Correct!", i)

		}
	}

}

func (segment *Segment) TestUpdateSegmentPosition(t *testing.T) {
	type test struct {
		inputSeg     *Segment
		inputPrevSeg *Segment
		inputGene    []SegmentGene
		outputSeg    *Segment
	}

	inputDirectory1 := "tests/UpdateSegmentPosition/input1/"
	inputDirectory2 := "tests/UpdateSegmentPosition/input2/"
	outputDirectory := "tests/UpdateSegmentPosition/output/"

	inputFiles1 := ReadFilesFromDirectory(inputDirectory1)
	inputFiles2 := ReadFilesFromDirectory(inputDirectory2)

	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles1), len(outputFiles))
	AssertEqualAndNonzero(len(inputFiles2), len(outputFiles))

	//create an array of tests
	tests := make([]test, len(inputFiles1))

	//range through the input and output files and set the test values
	for i := range inputFiles1 {
		tests[i].inputSeg = ReadSegmentTreeFromFile(inputDirectory1, inputFiles1[i])
		tests[i].inputPrevSeg, tests[i].inputGene = ReadUpdateSegPosInputFromFile(inputDirectory2, inputFiles2[i])
		tests[i].outputSeg = ReadSegmentTreeFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		test.inputSeg.UpdateSegmentPosition(test.inputPrevSeg, test.inputGene)
		//check if positions of the segments in a segment tree are calculated correctly
		if !SegTreeIsTheSame(test.inputSeg, test.outputSeg) {
			t.Errorf("Error! For input test dataset %d your function failed", i)
		} else {
			fmt.Printf("For input test dataset %d :Correct!", i)
		}
	}
}

func (segment *Segment) TestCalculateSegmentPosition(t *testing.T) {
	type test struct {
		inputSeg     *Segment
		inputPrevSeg *Segment
		inputGene    []SegmentGene
		outputPos    OrderedPair
	}

	inputDirectory1 := "tests/CalculateSegmentPosition/input1/"
	inputDirectory2 := "tests/CalculateSegmentPosition/input2/"
	outputDirectory := "tests/CalculateSegmentPosition/output/"

	inputFiles1 := ReadFilesFromDirectory(inputDirectory1)
	inputFiles2 := ReadFilesFromDirectory(inputDirectory2)

	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles1), len(outputFiles))
	AssertEqualAndNonzero(len(inputFiles2), len(outputFiles))
	//create an array of tests
	tests := make([]test, len(inputFiles1))

	//range through the input and output files and set the test values
	for i := range inputFiles1 {
		tests[i].inputSeg = ReadSegmentTreeFromFile(inputDirectory1, inputFiles1[i])
		tests[i].inputPrevSeg, tests[i].inputGene = ReadUpdateSegPosInputFromFile(inputDirectory2, inputFiles2[i])
		tests[i].outputPos = ReadOrderedPairFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		test.inputSeg.CalculateSegmentPosition(test.inputPrevSeg, test.inputGene)
		//check if the position of the current segment is calculated correctly
		if test.inputSeg.position.x != test.outputPos.x || test.inputSeg.position.y != test.outputPos.y {
			t.Errorf("Error! For input test dataset %d your function failed", i)
		} else {
			fmt.Printf("For input test dataset %d :Correct!", i)
		}
	}

}

func TestUpdateVelocity(t *testing.T) {
	type test struct {
		inputPond *Pond
		inputBot  *Swimbot
		outputBot *Swimbot
	}

	inputDirectory1 := "tests/UpdateVelocity/input1/"
	inputDirectory2 := "tests/UpdateVelocity/input2/"
	outputDirectory := "tests/UpdateVelocity/output/"

	inputFiles1 := ReadFilesFromDirectory(inputDirectory1)
	inputFiles2 := ReadFilesFromDirectory(inputDirectory2)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles1), len(outputFiles))
	AssertEqualAndNonzero(len(inputFiles2), len(outputFiles))

	//create an array of tests
	tests := make([]test, len(inputFiles1))

	//range through the input and output files and set the test values
	for i := range inputFiles1 {
		tests[i].inputPond = ReadPondFromFile(inputDirectory1, inputFiles1[i])
		tests[i].inputBot = ReadSwimbotFromFile(inputDirectory2, inputFiles2[i])
		tests[i].outputBot = ReadSwimbotFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		test.inputBot.UpdateVelocity(test.inputPond)
		//check if the bot's velocity is updated correctly
		if !SwimbotistheSame(test.inputBot, test.outputBot) {
			t.Errorf("Error! For input test dataset %d your function failed", i)
		} else {
			fmt.Printf("For input test dataset %d :Correct!", i)
		}
	}
}

func TestMating(t *testing.T) {
	type test struct {
		inputPond *Pond
		indexS1   int
		indexS2   int
		indexKid  int
		outputBot *Swimbot
	}

	inputDirectory1 := "tests/Mating/input1/"
	inputDirectory2 := "tests/Mating/input2/"
	outputDirectory := "tests/Mating/output/"

	inputFiles1 := ReadFilesFromDirectory(inputDirectory1)
	inputFiles2 := ReadFilesFromDirectory(inputDirectory2)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles1), len(outputFiles))
	AssertEqualAndNonzero(len(inputFiles2), len(outputFiles))

	//create an array of tests
	tests := make([]test, len(inputFiles1))

	//range through the input and output files and set the test values
	for i := range inputFiles1 {
		tests[i].inputPond = ReadPondFromFile(inputDirectory1, inputFiles1[i])
		tests[i].indexS1, tests[i].indexS2, tests[i].indexKid = ReadThreeIntFromFile(inputDirectory2, inputFiles2[i])
		tests[i].outputBot = ReadSwimbotFromFile(outputDirectory, outputFiles[i])
	}

	for i, test := range tests {
		outcome := test.inputPond.Mating(test.indexS1, test.indexS2, test.indexKid)
		//check if the mating function produces a child normally
		if !ChildbotistheSame(outcome, test.outputBot) {
			t.Errorf("Error! For input test dataset %d your function failed", i)
		} else {
			fmt.Printf("For input test dataset %d :Correct!", i)
		}
	}

}

func ReadUpdateSegPosInputFromFile(directory string, inputFile os.FileInfo) (*Segment, []SegmentGene) {
	fileName := inputFile.Name() //grab file name

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//create a new segment(which acts as a previous segment in the UpdateSegmentPosition function, so only the position of the segment matters)
	var presegment Segment
	preseg := &presegment

	//create a new slice of segmentgenes
	SegGenes := make([]SegmentGene, 0)

	contentIndex := 0
	for _, inputLine := range inputLines {
		if inputLine == "-" {
			contentIndex += 1
			continue
		}

		//contentIndex == 0 indicates collecting information of the segment
		if contentIndex == 0 {
			currentLine := strings.Split(inputLine, " ")
			// A line contains position.x, position.y, angle, index fields of a segment orderly
			preseg.position.x, err = strconv.ParseFloat(currentLine[0], 64)
			if err != nil {
				panic(err)
			}
			preseg.position.y, err = strconv.ParseFloat(currentLine[1], 64)
			if err != nil {
				panic(err)
			}
			preseg.angle, err = strconv.ParseFloat(currentLine[2], 64)
			if err != nil {
				panic(err)
			}
			preseg.index, err = strconv.Atoi(currentLine[3])

			if err != nil {
				panic(err)
			}
		}
		//-----------------------------------------------------------------------------
		//contentIndex == 1 indicates collecting information of the segment gene slice
		if contentIndex == 1 {
			//A line represents a gene slice of a segment in the total SegmentGene slice
			var gene SegmentGene
			currentLine := strings.Split(inputLine, " ")
			//A line contains six float values corresponding to six genes of a segment
			for i := range currentLine {
				value, err := strconv.ParseFloat(currentLine[i], 64)
				if err != nil {
					panic(err)
				}
				gene = append(gene, value)
			}

			SegGenes = append(SegGenes, gene)
		}

	}
	return preseg, SegGenes
}

// SwimbotistheSame checks whether two swimbots are identical in terms of their properties
func SwimbotistheSame(bot1 *Swimbot, bot2 *Swimbot) bool {
	if bot1.position.x != bot2.position.x || bot2.position.y != bot1.position.y {
		return false
	} else if bot1.velocity.x != bot2.velocity.x || bot2.velocity.y != bot1.velocity.y {
		return false
	} else if bot1.energy != bot2.energy {
		return false
	} else if bot1.mass != bot2.mass {
		return false
	} else if bot1.age != bot2.age {
		return false
	} else if bot1.goal.isBot != bot2.goal.isBot || bot1.goal.index != bot2.goal.index {
		return false
	} else if !SliceIsTheSame(bot1.family, bot2.family) {
		return false
	}
	return true
}

// SwimbotistheSame checks whether two childbots are identical in terms of their properties
// This is used to check whether a childbot is generated properly.
// the velocity and mass of the childbots are not checked beacuse these properties will be generated randomly
func ChildbotistheSame(bot1 *Swimbot, bot2 *Swimbot) bool {
	if bot1.position.x != bot2.position.x || bot2.position.y != bot1.position.y {
		return false
	} else if bot1.energy != bot2.energy {
		return false
	} else if bot1.age != bot2.age {
		return false
	} else if bot1.goal.isBot != bot2.goal.isBot || bot1.goal.index != bot2.goal.index {
		return false
	} else if !SliceIsTheSame(bot1.family, bot2.family) {
		return false
	}
	return true
}

// This function checks if two slices of intergers are exactly identical or not.
func SliceIsTheSame(sl1, sl2 []int) bool {
	//First check if the two slices have the same length
	if len(sl1) != len(sl2) {
		return false
	} else {
		//range through all the elements in a slice and check if it's identical to the corresponding element in the other slice
		for i := range sl1 {
			if sl1[i] != sl2[i] {
				return false
			}
		}
	}
	return true
}

func ReadSwimbotFromFile(directory string, file os.FileInfo) *Swimbot {
	fileName := file.Name() //grab file name

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}
	//create a new swimbot
	var bot Swimbot
	b := &bot

	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//The first line contains information of a bot's goal(isBot && index)
	currentLine := strings.Split(inputLines[0], " ")
	int1, err := strconv.Atoi(currentLine[0])
	if err != nil {
		panic(err)
	}
	if int1 == 0 {
		bot.goal.isBot = false
	} else {
		bot.goal.isBot = true
	}
	bot.goal.index, err = strconv.Atoi(currentLine[1])
	if err != nil {
		panic(err)
	}

	//The second line contains information of the bot's age
	bot.age, err = strconv.ParseFloat(inputLines[1], 64)
	if err != nil {
		panic(err)
	}

	//The third line contains information of the bot's energy level
	bot.energy, err = strconv.ParseFloat(inputLines[2], 64)
	if err != nil {
		panic(err)
	}

	//The fourth line contains information of the bot's position
	currentLine = strings.Split(inputLines[3], " ")
	bot.position.x, err = strconv.ParseFloat(currentLine[0], 64)
	if err != nil {
		panic(err)
	}
	bot.position.y, err = strconv.ParseFloat(currentLine[1], 64)
	if err != nil {
		panic(err)
	}

	//The fifth line contains information of the bot's velocity
	currentLine = strings.Split(inputLines[4], " ")
	bot.velocity.x, err = strconv.ParseFloat(currentLine[0], 64)
	if err != nil {
		panic(err)
	}
	bot.velocity.y, err = strconv.ParseFloat(currentLine[1], 64)
	if err != nil {
		panic(err)
	}

	//The sixth line contains information of the bot's mass
	bot.mass, err = strconv.ParseFloat(inputLines[5], 64)
	if err != nil {
		panic(err)
	}

	//The seventh line contains information of the bot's family
	currentLine = strings.Split(inputLines[6], " ")
	for _, str := range currentLine {
		value, err := strconv.Atoi(str)
		bot.family = append(bot.family, value)
		if err != nil {
			panic(err)
		}
	}

	//The seventh line contains information of the bot's botgene(commongene)
	currentLine = strings.Split(inputLines[7], " ")
	bot.botGene.angularMovement, err = strconv.ParseFloat(currentLine[0], 64)
	if err != nil {
		panic(err)
	}
	bot.botGene.translationalMovement, err = strconv.ParseFloat(currentLine[1], 64)
	if err != nil {
		panic(err)
	}
	bot.botGene.numSegments, err = strconv.Atoi(currentLine[2])
	if err != nil {
		panic(err)
	}

	return b
}

func ReadThreeIntFromFile(directory string, file os.FileInfo) (int, int, int) {
	fileName := file.Name()

	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	int1, err := strconv.Atoi(outputLines[0])
	if err != nil {
		panic(err)
	}

	int2, err := strconv.Atoi(outputLines[1])
	if err != nil {
		panic(err)
	}

	int3, err := strconv.Atoi(outputLines[2])
	if err != nil {
		panic(err)
	}

	return int1, int2, int3
}

func ReadPondFromFile(directory string, file os.FileInfo) *Pond {
	fileName := file.Name()

	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//create a new pond object
	var pond Pond
	p := &pond

	//first, read lines and split along blank space
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	contentIndex := 0
	botIndex := -1
	foodIndex := 0
	var property int

	for _, inputLine := range inputLines {
		if inputLine == "S" {
			//contentIndex being 1 means collecting information of swimbots
			contentIndex += 1
			continue
		}

		if inputLine == "F" {
			//contentIndex being 2 means collecting information of
			contentIndex += 1
			continue
		}

		//contentIndex being 1 means collecting basic information of the pond
		if contentIndex == 0 {
			currentLine := strings.Split(inputLine, " ")
			//currentLine contains length of the swimbot slice, length of the foodbit slice and width of the pond
			numBots, err := strconv.Atoi(currentLine[0])
			if err != nil {
				panic(err)
			}
			numFood, err := strconv.Atoi(currentLine[1])
			if err != nil {
				panic(err)
			}
			pond.width, err = strconv.ParseFloat(currentLine[2], 64)
			if err != nil {
				panic(err)
			}
			pond.swimbots = make([]*Swimbot, numBots)
			pond.foodBits = make([]*Food, numFood)
		}

		//the information of each swimbot will be separated by a '-', followed by each line containing one property of the swimbot.
		if contentIndex == 1 {
			if inputLine == "-" {
				botIndex += 1
				property = 0
				continue
			}

			if property == 0 {
				currentLine := strings.Split(inputLine, " ")
				int1, err := strconv.Atoi(currentLine[0])
				if err != nil {
					panic(err)
				}
				if int1 == 0 {
					pond.swimbots[botIndex].goal.isBot = false
				} else {
					pond.swimbots[botIndex].goal.isBot = true
				}
				pond.swimbots[botIndex].goal.index, err = strconv.Atoi(currentLine[1])
				if err != nil {
					panic(err)
				}
				property += 1
			} else if property == 1 {
				pond.swimbots[botIndex].age, err = strconv.ParseFloat(inputLines[1], 64)
				if err != nil {
					panic(err)
				}
				property += 1

			} else if property == 2 {
				pond.swimbots[botIndex].energy, err = strconv.ParseFloat(inputLines[2], 64)
				if err != nil {
					panic(err)
				}
				property += 1
			} else if property == 3 {
				currentLine := strings.Split(inputLine, " ")
				pond.swimbots[botIndex].position.x, err = strconv.ParseFloat(currentLine[0], 64)
				if err != nil {
					panic(err)
				}
				pond.swimbots[botIndex].position.y, err = strconv.ParseFloat(currentLine[1], 64)
				property += 1
			} else if property == 4 {
				currentLine := strings.Split(inputLine, " ")
				pond.swimbots[botIndex].velocity.x, err = strconv.ParseFloat(currentLine[0], 64)
				if err != nil {
					panic(err)
				}
				pond.swimbots[botIndex].velocity.y, err = strconv.ParseFloat(currentLine[1], 64)
				if err != nil {
					panic(err)
				}
				property += 1
			} else if property == 5 {
				pond.swimbots[botIndex].mass, err = strconv.ParseFloat(inputLines[5], 64)
				if err != nil {
					panic(err)
				}
				property += 1
			} else if property == 6 {
				currentLine := strings.Split(inputLine, " ")
				for _, str := range currentLine {
					value, err := strconv.Atoi(str)
					pond.swimbots[botIndex].family = append(pond.swimbots[botIndex].family, value)
					if err != nil {
						panic(err)
					}

				}
				if err != nil {
					panic(err)
				}
				property += 1
			} else if property == 7 {
				currentLine := strings.Split(inputLine, " ")
				pond.swimbots[botIndex].botGene.angularMovement, err = strconv.ParseFloat(currentLine[0], 64)
				if err != nil {
					panic(err)
				}
				pond.swimbots[botIndex].botGene.translationalMovement, err = strconv.ParseFloat(currentLine[1], 64)
				if err != nil {
					panic(err)
				}
				pond.swimbots[botIndex].botGene.numSegments, err = strconv.Atoi(currentLine[2])
				if err != nil {
					panic(err)
				}
				property += 1
			}
		}

		if contentIndex == 2 {
			//'nil' means teh foodbit has been eaten
			if inputLine == "nil" {
				pond.foodBits[foodIndex] = nil

			} else {
				//each line contains the position of a foodbit
				currentLine := strings.Split(inputLine, " ")

				pond.foodBits[foodIndex].position.x, err = strconv.ParseFloat(currentLine[0], 64)
				if err != nil {
					panic(err)
				}
				pond.foodBits[foodIndex].position.y, err = strconv.ParseFloat(currentLine[1], 64)
				if err != nil {
					panic(err)
				}
			}
			foodIndex += 1

		}
	}

	return p
}

func ReadFilesFromDirectory(directory string) []os.FileInfo {
	dirContents, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Error reading directory: " + directory)
	}

	return dirContents
}

func AssertEqualAndNonzero(length0, length1 int) {
	if length0 == 0 {
		panic("No files present in given directory.")
	}
	if length1 == 0 {
		panic("No files present in given directory.")
	}
	if length0 != length1 {
		panic("Number of files in directories doesn't match.")
	}
}

// SegTreeIsTheSame recursivey checks if two segmenttrees are identical
func SegTreeIsTheSame(segtree1, segtree2 *Segment) bool {
	//First check is the current segments have the same properties
	if !SegmentIsTheSame(segtree1, segtree2) {
		return false

		//If the segment has subsegments, check if all the subsegments are also identical
	} else if segtree1.subSegments != nil {
		//first check if they have the same amount of subsegments
		if len(segtree1.subSegments) != len(segtree2.subSegments) {
			return false
		} else {
			for i := range segtree1.subSegments {
				if !SegmentIsTheSame(segtree1.subSegments[i], segtree2.subSegments[i]) {
					return false
				}
			}
		}
	}

	return true
}

// SegmentIsTheSame checks whether the properties(position,angletoparent and index) of two segments are the same or not
func SegmentIsTheSame(s1, s2 *Segment) bool {

	if s1.position.x != s2.position.x {
		return false
	} else if s1.position.y != s2.position.y {
		return false
	} else if s1.angle != s2.angle {
		return false
	} else if s1.index != s2.index {
		return false
	}
	return true
}

func ReadSegmentTreeFromFile(directory string, file os.FileInfo) *Segment {
	var mainseg Segment
	ms := &mainseg
	fileName := file.Name()

	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//create a slice of parentsegments to keep track of which segment to attach to
	ParentSegments := make([]*Segment, len(inputLines))
	//CurrentParentIndex is a pointer of which parentsegment's subsegments are being visited now
	CurrentParentIndex := 0
	for row, inputLine := range inputLines {
		//The first line contains the information of the mainsegment
		if row == 0 {
			currentLine := strings.Split(inputLine, " ")
			ms.position.x, err = strconv.ParseFloat(currentLine[0], 64)
			if err != nil {
				panic(err)
			}
			ms.position.y, err = strconv.ParseFloat(currentLine[1], 64)
			if err != nil {
				panic(err)
			}
			ms.angle, err = strconv.ParseFloat(currentLine[2], 64)
			if err != nil {
				panic(err)
			}
			ms.index, err = strconv.Atoi(currentLine[3])
			if err != nil {
				panic(err)
			}

			//let the CurrenParentIndex point to the main segment
			//Include the main segment to the slice of parentsegments
			ParentSegments[CurrentParentIndex] = ms
			CurrentParentIndex -= 1
			continue
		}

		if inputLine == "EnteringSubsegments" {
			//when visiting the subsegments of a segment, move the pointer forward
			CurrentParentIndex += 1
			continue
		}

		if inputLine == "LeavingSubsegments" {
			//when finished visiting the subsegments of a segment, move the pointer backward
			CurrentParentIndex -= 1
			continue
		}

		//create a segment for each line containing infomration
		var seg Segment
		s := &seg
		currentLine := strings.Split(inputLine, " ")
		s.position.x, err = strconv.ParseFloat(currentLine[0], 64)
		if err != nil {
			panic(err)
		}
		s.position.y, err = strconv.ParseFloat(currentLine[1], 64)
		if err != nil {
			panic(err)
		}
		s.angle, err = strconv.ParseFloat(currentLine[2], 64)
		if err != nil {
			panic(err)
		}

		s.index, err = strconv.Atoi(currentLine[3])
		if err != nil {
			panic(err)
		}

		//if the subsegments of the current segment is to be visited, append the segment to the parent slice
		if inputLines[row+1] == "EnteringSubsegments" {
			ParentSegments[CurrentParentIndex].subSegments = append(ParentSegments[CurrentParentIndex].subSegments, s)
		}

	}

	return ms
}

func ReadOrderedPairFromFile(directory string, file os.FileInfo) OrderedPair {
	fileName := file.Name() //grab file name

	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}
	//trim out extra space and store as a slice of strings, each containing one line.
	lines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")
	currentLine := strings.Split(lines[0], " ")

	var position OrderedPair
	position.x, err = strconv.ParseFloat(currentLine[0], 64)
	if err != nil {
		panic(err)
	}
	position.y, err = strconv.ParseFloat(currentLine[1], 64)
	if err != nil {
		panic(err)
	}

	return position
}
