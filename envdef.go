package envdef

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Diff(source, dist string, overwrite bool) (*Result, error) {
	insertSlice := InsertSlice{}
	updateSlice := UpdateSlice{}
	deleteSlice := DeleteSlice{}
	noChangeSlice := NoChangeSlice{}

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
			// nochange
			if !overwrite || v == distEnv[k] {
				noChangeSlice = append(noChangeSlice, envFormat(k, v))
				continue
			}

			if v != distEnv[k] {
				updateSlice = append(updateSlice, envFormat(k, v))
				continue
			}
		}

		// insert
		insertSlice = append(insertSlice, envFormat(k, v))
	}

	for k, v := range distEnv {
		if _, ok := sourceEnv[k]; !ok {
			deleteSlice = append(deleteSlice, envFormat(k, v))
		}
	}

	return &Result{
		InsertSlice:   insertSlice,
		UpdateSlice:   updateSlice,
		DeleteSlice:   deleteSlice,
		NoChangeSlice: noChangeSlice,
	}, nil
}

func envFormat(k, v string) string {
	format := "%v=%v"
	return fmt.Sprintf(format, k, v)
}

func Read(path string) (map[string]string, error) {
	e, err := godotenv.Read(path)
	if err != nil {
		return nil, err
	}

	return e, nil
}
