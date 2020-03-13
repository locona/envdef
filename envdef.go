package envdef

import (
	"fmt"

	"github.com/joho/godotenv"
)

var iconFormat = "%v %v"
var envFormat = "%v=%v"

func Diff(source, dist string) (*Result, error) {
	var (
		insertSlice   InsertSlice
		updateSlice   UpdateSlice
		deleteSlice   DeleteSlice
		noChangeSlice NoChangeSlice
	)

	sourceEnv, err := Read(source)
	if err != nil {
		return nil, err
	}

	distEnv, err := Read(dist)
	if err != nil {
		return nil, err
	}

	for k, v := range sourceEnv {
		// update
		if _, ok := distEnv[k]; ok {
			if v != distEnv[k] {
				updateSlice = append(updateSlice, fmt.Sprintf(envFormat, k, v))
				continue
			}

			noChangeSlice = append(noChangeSlice, fmt.Sprintf(envFormat, k, v))

			continue
		}

		// insert
		insertSlice = append(insertSlice, fmt.Sprintf(envFormat, k, v))
	}

	for k, v := range distEnv {
		if _, ok := sourceEnv[k]; !ok {
			deleteSlice = append(deleteSlice, fmt.Sprintf(envFormat, k, v))
		}
	}

	return &Result{
		InsertSlice:   insertSlice,
		UpdateSlice:   updateSlice,
		DeleteSlice:   deleteSlice,
		NoChangeSlice: noChangeSlice,
	}, nil
}

func Read(path string) (map[string]string, error) {
	e, err := godotenv.Read(path)
	if err != nil {
		return nil, err
	}

	return e, nil
}
