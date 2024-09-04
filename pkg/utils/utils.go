package utils

import (
	"encoding/json"
	"os"
	"test/pkg/models"
)

func ReadProgramListFromFile(filepath string) (*models.ProgramList, error) {
	var ProgramList models.ProgramList
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &ProgramList)
	if err != nil {
		return nil, err
	}

	return &ProgramList, nil
}

func WriteProgramListToFile(filepath string, ProgramList *models.ProgramList) error {
	dumpData, err := json.Marshal(&ProgramList)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, dumpData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadProgramDataFromFile(filepath string) (*models.ProgramData, error) {
	var ProgramData models.ProgramData
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &ProgramData)
	if err != nil {
		return nil, err
	}

	return &ProgramData, nil
}

func WriteProgramDataToFile(filepath string, ProgramData *models.ProgramData) error {
	dumpData, err := json.Marshal(&ProgramData)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, dumpData, 0644)
	if err != nil {
		return err
	}

	return nil
}
