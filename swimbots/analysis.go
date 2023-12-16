package main

import (
    "fmt"
    "encoding/csv"
    "os"
    "strconv"
    "math"
)

// GenerateAnalysis() tales in a pond object
// if saves all the data from generation 0 and the last generation into csv files in the corresponding folder
// it also writes the analysis data to the "Results.txt" file
func GenerateAnalysis(pond0, pondN *Pond, numGen int){
    // generate the map for energy for the last generation
    eN:= GetEnergiesMap(pondN)

    // generate the map for translational movement for first and last generation
    t0:= GetTranslationalMovementMap(pond0)
    tN:= GetTranslationalMovementMap(pondN)

    // generate the map for angular movement for first and last generation
    a0:= GetAngularMovementMap(pond0)
    aN:= GetAngularMovementMap(pondN)

    // generate the map for numsegments for first and last generation
    s0:= GetNumSegmentsMap(pond0)
    sN:= GetNumSegmentsMap(pondN)

    // generate the map for age for the last generation
    //age0:= GetAgeMap(timePoints[0])
    ageN:= GetAgeMap(pondN)

    segLen0:= GetSegLengthMap(pond0)
    segLenN:= GetSegLengthMap(pondN)

    // write the generated maps to csv files
    WriteToCSV_int(s0, "csvFiles/segGen0")
    WriteToCSV_float(a0, "csvFiles/angular0")
    WriteToCSV_int(t0, "csvFiles/transl0")
    WriteToCSV_int(segLen0, "csvFiles/segLen0")
    // WriteToCSV_int(e0, "csvFiles/energy0")
    // WriteToCSV_int(age0, "csvFiles/age0")

    WriteToCSV_int(sN, "csvFiles/segGenEnd")
    WriteToCSV_float(aN, "csvFiles/angularEnd")
    WriteToCSV_int(tN, "csvFiles/translEnd")
    WriteToCSV_int(eN, "csvFiles/energylEnd")
    WriteToCSV_int(ageN, "csvFiles/ageEnd")
    WriteToCSV_int(segLenN, "csvFiles/segLenEnd")


    //  create the results text file and write ot the analysis results to it
    fileToWriteTo, err1 := os.Create("Results.txt")
    if err1 != nil {
        panic(err1)
    }

    resultReport:= "Results" + "\n" + "\n" + "This is a summary of the results obtained between generation 0 and " + strconv.Itoa(numGen) + " generation."+"\n"+"\n"
    _, err0 := fileToWriteTo.WriteString(resultReport)
    if err0 != nil {
        fmt.Println(err0)
        fileToWriteTo.Close()
        return
    }

    numBotsInitial:=0
    for i:=0; i<len(pond0.swimbots); i++{
        if pond0.swimbots[i]!=nil{
            numBotsInitial+=1
        }
    }

    numBotsFinal:=0
    for i:=0; i<len(pondN.swimbots); i++{
        if pondN.swimbots[i]!=nil{
            numBotsFinal+=1
        }
    }

    numBotsLeft:= "We start out with "+ strconv.Itoa(numBotsInitial)+" swimbots, and end up with "+strconv.Itoa(numBotsFinal) +" bots."+"\n"+"\n"

    _, err0_0 := fileToWriteTo.WriteString(numBotsLeft)
    if err0_0 != nil {
        fmt.Println(err0_0)
        fileToWriteTo.Close()
        return
    }

    // write in energy results
    resultEnergy:= GetEnergyStats(eN, pondN)
    _, err := fileToWriteTo.WriteString(resultEnergy)
    if err != nil {
        fmt.Println(err)
        fileToWriteTo.Close()
        return
    }

    // write in age results
    resultAge:= GetAgeStats(pondN)
    _, err2 := fileToWriteTo.WriteString(resultAge)
    if err != nil {
        fmt.Println(err2)
        fileToWriteTo.Close()
        return
    }

    // write in segment number results
    resultSeg:= GetSegStats(s0, sN)
    _, err3 := fileToWriteTo.WriteString(resultSeg)
    if err3 != nil {
        fmt.Println(err3)
        fileToWriteTo.Close()
        return
    }

    // write in translational movement results
    resultTransl:= GetTranslStats(t0, tN)
    _, err4 := fileToWriteTo.WriteString(resultTransl)
    if err4 != nil {
        fmt.Println(err4)
        fileToWriteTo.Close()
        return
    }

    // write in rotational movement results
    resultRot:= GetRotStats(a0, aN)
    _, err5 := fileToWriteTo.WriteString(resultRot)
    if err5 != nil {
        fmt.Println(err5)
        fileToWriteTo.Close()
        return
    }

    // write in segment length results
    resultSegLen:= GetSegLengthStats(segLen0, segLenN)
    _, err6 := fileToWriteTo.WriteString(resultSegLen)
    if err6 != nil {
        fmt.Println(err6)
        fileToWriteTo.Close()
        return
    }

}

// GetEnergiesMap() takes in a pointer to a pond
// returns a map of the energy values of swimbots at this stage and how many correspond to each
func GetEnergiesMap(pond *Pond) map[int]int{
    energyMap := make(map[int]int)

    for n:=0; n<100; n++{
        if n%5==0{
            energyMap[n]=0
        }
    }

    for i:=0; i<len(pond.swimbots); i++{
        if pond.swimbots[i]!=nil{
            //fmt.Println(ond.swimbots[i].energy)
            eBot:= int(pond.swimbots[i].energy)
            for k, _:= range energyMap{
                if eBot>=k && eBot<k+5{
                    energyMap[k]+=1
                }
            }
        }
    }
    return energyMap
}

// GetAgeMap takes in a pointer to a pond
// returns a map of the ages  of swimbots at this stage and how many correspond to each
func GetAgeMap(pond *Pond) map[int]int{
    ageMap := make(map[int]int)

    // get the maximum age
    max := GetMaximumAge(pond)

    // set the bins such that each we have 10 bins from 0 to max age
    for n:=0; n<10; n++{
        ageMap[int(n*int(max)/10)]=0
    }

    // loop over the swimbots and add in the corresponding bins
    for i:=0; i<len(pond.swimbots); i++{
        if pond.swimbots[i]!=nil{
            ageBot:= int(pond.swimbots[i].age)
            for k, _ := range ageMap {
                if ageBot>=k && ageBot<k+int(max/10) {
                    ageMap[k]+=1
                }
            }
        }
    }
    return ageMap
}
// GetSegLengthMap takes in a pointer to a pond
// returns a map of the length of the main segments for swimbots at this stage and how many correspond to each
func GetSegLengthMap(pond *Pond) map[int]int{
    segMap := make(map[int]int)

    // set the bins such that each we have 10 bins from 0 to max age
    // range of seg length: 5-20
    for n := 5; n < 20; n = n + 2 {
        segMap[n] = 0
    }

    // loop over the swimbots and add in the corresponding bins
    for i:=0; i<len(pond.swimbots); i++{
        if pond.swimbots[i]!=nil{
            lengthBot:= int(pond.swimbots[i].segGenes[0][4])
            for k, _ := range segMap {
                if lengthBot>=k && lengthBot<k+2 {
                    segMap[k]+=1
                }
            }
        }
    }
    return segMap
}

// GetTranslationalMovementMap() takes in a pointer to a pond
// returns a map of the translational movement values of swimbots at this stage and how many correspond to each
func GetTranslationalMovementMap(pond *Pond) map[int]int{
    translationalMovementMap := make(map[int]int)

    // set the keys of the map based on range of close valuessince we are working with floats
    for n:=0; n<10; n++{
        translationalMovementMap[n]=0
    }

    for i:=0; i<len(pond.swimbots); i++{
        if pond.swimbots[i]!=nil{
            transMov:= pond.swimbots[i].botGene.translationalMovement
            obtainedNum, _ := math.Modf(transMov)
            translationalMovementMap[int(obtainedNum)]+=1
        }
    }
    return translationalMovementMap
}

// GetANgularMovementMap() takes in a pointer to a pond
// returns a map of the angular movement values of swimbots at this stage and how many correspond to each
func GetAngularMovementMap(pond *Pond) map[float64]int{
    angularMovementMap := make(map[float64]int)

    // set the keys of the map based on range of close valuessince we are working with floats
    for n:=0; n<10; n++{
        angularMovementMap[float64(n)*math.Pi*0.025]=0
    }

    // loop over the swimbots and add in the corresponding bins
    for i:=0; i<len(pond.swimbots); i++{
        if pond.swimbots[i]!=nil{
            angMov:= pond.swimbots[i].botGene.angularMovement
            for k, _ := range angularMovementMap {
                if angMov>=k && angMov<k+math.Pi*0.025{
                    angularMovementMap[k]+=1
                }
            }
        }
    }
    return angularMovementMap
}

//GetNumSegmentsMap() gets the number of segments frequency
// returns a map with key value pairs corresponding to the number of segments and their frequency of occurence
func GetNumSegmentsMap(pond *Pond) map[int]int{
    numSegmentsMap := make(map[int]int)

    // set the keys of the map based on range of close valuessince we are working with floats
    for n:=2; n<9; n++{
        numSegmentsMap[n]=0
    }

    // everytime we encounter a key, we increment it's value
    for i:=0; i<len(pond.swimbots); i++{
        if pond.swimbots[i]!=nil{
            nSeg:= pond.swimbots[i].botGene.numSegments
            numSegmentsMap[nSeg]+=1
        }
    }
    return numSegmentsMap
}

// WriteToCSV_int() take a mapp that maps integers to integers
// it writes out the map into a csv file
func WriteToCSV_int(m map[int]int, filename string){
    //create the csv file
    csvFile, err1 := os.Create(filename+".csv")
    if err1!=nil{
        panic(err1)
    }
    defer csvFile.Close()

    // create the writer object
    csvwriter := csv.NewWriter(csvFile)
    defer csvwriter.Flush()

    // add the first line to the csv file
    // this line will have the column names
    // the first column has the same name as the filename which will represent the variable
    // the second column will have the counts (values) corresponding to the keys in column one
    firstLine:= []string{filename, "counts"}
    if err := csvwriter.Write(firstLine); err !=nil{
        panic(err)
    }

    // range over the map
    // add the keys and values to the csv file
    //fmt.Println("S1: ", s1)
    for key, value := range m {
        csvLine := make([]string, 0)
        csvLine = append(csvLine, strconv.Itoa(key), strconv.Itoa(value))
        //fmt.Println("csvLine", csvLine)
        if err := csvwriter.Write(csvLine); err !=nil{
            panic(err)
        }
    }
}

// WriteToCSV_float() take a mapp that maps float64 keys to integer values
// it writes out the map into a csv file
func WriteToCSV_float(m map[float64]int, filename string){
    //create the csv file
    csvFile, err1 := os.Create(filename+".csv")
    if err1!=nil{
        panic(err1)
    }
    defer csvFile.Close()

    // create the writer object
    csvwriter := csv.NewWriter(csvFile)
    defer csvwriter.Flush()

    // add the first line to the csv file
    // this line will have the column names
    // the first column has the same name as the filename which will represent the variable
    // the second column will have the counts (values) corresponding to the keys in column one
    firstLine:= []string{filename, "counts"}
    if err := csvwriter.Write(firstLine); err !=nil{
        panic(err)
    }

    // range over the map
    // add the keys and values to the csv file
    for key, value := range m {
        csvLine := make([]string, 0)
        kNew:= fmt.Sprintf("%v", key)
        csvLine = append(csvLine, kNew, strconv.Itoa(value))
        // fmt.Println("csvLine", csvLine)
        if err := csvwriter.Write(csvLine); err !=nil{
            panic(err)
        }
    }
}

// GetEnergyStats returns a string with the max and average enerfy at the end of a simulation
func GetEnergyStats(m map[int]int, pond *Pond) string{
    // get max energy
    maxEnergy:= GetMaximumEnergy(pond)

    // get average energy
    avgEn:= GetAverage(m)

    // return the result to be typed into the file
    resultEnergy:= "The maximum energy in the last generation was " + fmt.Sprintf("%f", maxEnergy)+ " and the mean energy was "+ fmt.Sprintf("%f", avgEn) + "." + "\n"+"\n"
    return resultEnergy
}

func GetMaximumEnergy(pond *Pond) float64 {
    // get the maximum age
    max := 0.0
    for i:=0; i<len(pond.swimbots); i++{
        if pond.swimbots[i]!=nil{
            if pond.swimbots[i].energy>max{
                max= pond.swimbots[i].energy
            }
        }
    }
    return max
}

// GetAgeStats returns a string with the max, min, and average age at the end of a simulation
func GetAgeStats(pondN *Pond) string {
    // get max age
    maxAge:= GetMaximumAge(pondN)

    // get min age
    minAge:= GetMinAge(pondN)

    // get average energy
    avgAge:=GetAverageAge(pondN)

    // return the result to be typed into the file
    resultAge:= "The maximum age in the last generation was " +  strconv.Itoa(maxAge)+ " , the minimum was " + strconv.Itoa(minAge)+ " and the average age was "+ fmt.Sprintf("%f", avgAge) + "." + "\n"+"\n"
    return resultAge
}

func GetMaximumAge(pond *Pond) int {
    // get the maximum age
    max := 0.0
    for i:=0; i<len(pond.swimbots); i++{
        if pond.swimbots[i]!=nil{
            if pond.swimbots[i].age>max{
                max= pond.swimbots[i].age
            }
        }
    }
    return int(max)
}

func GetAverageAge(pond *Pond) float64 {
    sum := 0.0
    count := 0
    for i := range pond.swimbots {
        if pond.swimbots[i] != nil {
            sum += pond.swimbots[i].age
            count += 1
        }
    }
    return sum/float64(count)
}

func GetMinAge(pond *Pond) int {
    // get the maximum age
    min := 10000000.0
    for i:=0; i<len(pond.swimbots); i++{
        if pond.swimbots[i]!=nil{
            if pond.swimbots[i].age<min{
                min= pond.swimbots[i].age
            }
        }
    }
    return int(min)
}

// GetSegStats returns a string with the max, min, and average segment number at the begining and end of a simulation
func GetSegStats(m0 map[int]int, mN map[int]int) string {
    // get max age
    // fmt.Println(mN)
    maxSeg0:= GetMostFrequent(m0)
    avgSeg0:= GetAverage(m0)

    maxSegN:= GetMostFrequent(mN)
    avgSegN:= GetAverage(mN)

    // return the result to be typed into the file
    resultSeg:= "The most frequent segment number in the last generation was " + strconv.Itoa(maxSegN)+ " and the average segment number was "+ fmt.Sprintf("%f", avgSegN) + ", \n"+
    "compared to the most frequent being "+ strconv.Itoa(maxSeg0)+ " in the first generation, and the average being " + fmt.Sprintf("%f", avgSeg0)+"."+"\n"+"\n"
    return resultSeg
}

// GetSegStats returns a string with the max, min, and average segment number at the begining and end of a simulation
func GetSegLengthStats(m0 map[int]int, mN map[int]int) string{
    // get max age
    maxSeg0:= GetMostFrequent(m0)
    avgSeg0:= GetAverage(m0)

    maxSegN:= GetMostFrequent(mN)
    avgSegN:= GetAverage(mN)

    // return the result to be typed into the file
    resultSeg:= "The most frequent range of length of the first (main) segment in the last generation was " + strconv.Itoa(maxSegN) + " and the average was "+ fmt.Sprintf("%f", avgSegN) + ", \n"+
    "compared to the most frequent range of length of the first segment being  "+ strconv.Itoa(maxSeg0)+ " in the first generation, and the average being " + fmt.Sprintf("%f", avgSeg0)+"."+"\n"+"\n"
    return resultSeg
}

// GetTranslStats returns a string with the max and average enerfy at the end of a simulation
func GetTranslStats(m0, mN map[int]int) string{
    // get max energy
    // maxTransl0:= GetMax(m0)
    // maxTranslN:= GetMax(mN)

    // get min energy
    // minTransl0:= GetMin(m0)
    // minTranslN:= GetMin(mN)

    // get average energy
    avTransl0:= GetAverage(m0)
    avTranslN:= GetAverage(mN)

    // return the result to be typed into the file
    resultTransl:= "The average translational movement in the last generation was " + fmt.Sprintf("%f", avTranslN)+ " while in the first generation it was "+ fmt.Sprintf("%f", avTransl0)+"."+"\n"+"\n"

    return resultTransl
}

// GetRotStats returns a string with the max and average enerfy at the end of a simulation
func GetRotStats(m0, mN map[float64]int) string{

    // get average energy
    avRot0:= fmt.Sprintf("%f",GetAverageFloatMap(m0))
    avRotN:= fmt.Sprintf("%f",GetAverageFloatMap(mN))

    // return the result to be typed into the file
    // fmt.Println(m0)
    // fmt.Println(avRot0)
    resultRot:= "The average rotational movement in the last generation was " + avRotN+ " while in the first generation it was "+ avRot0+"."+"\n"+"\n"

    return resultRot
}

// get the max of all values in a map
func GetMostFrequent(m map[int]int) int{
    mostFreqKey := 0
    // fmt.Println("Map:", m)
    // fmt.Println("mostFreqKey:", m[mostFreqKey])
    for key, value := range m {
        if value > m[mostFreqKey]{
            mostFreqKey = key
        }
    }
    // fmt.Println("mostFreqKey:", mostFreqKey)
    return mostFreqKey
}

// get the min of all values in a mpa
func GetMin(m map[int]int) int{
    var min int
    // fmt.Println(m)
    for i := range m {
        min = i
        // fmt.Println(min)
        break
    }
    for key, value := range m {
        if value < m[min]{
            min= key
        }
    }
    return min
}

// get the average of all values in a map
func GetAverage(m map[int]int) float64{
    a:=0
    swimbotCount:= 0
    for key, val := range m {
        swimbotCount += val
        a += key * val
    }

    avgSeg := float64(a)/float64(swimbotCount)
    return avgSeg
}

// get the average of all values in a map for float to int
func GetAverageFloatMap(m map[float64]int) float64{
    a:=0.0
    swimbotCount:= 0.0
    for key,val := range m {
        swimbotCount += float64(val)
        a+= key * float64(val)
    }

    avg:= a/swimbotCount
    return avg
}
