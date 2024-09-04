package main

import (
	"log"
	"test/pkg/utils"
)

func main() {
	programList, err := utils.ReadProgramDataFromFile("cmd/input_test.json")
	if err != nil {
		panic(err)
	}
	log.Println(programList)
	utils.WriteProgramDataToFile("cmd/output_test.json", programList)
}
