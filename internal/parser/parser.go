package parser

import (
	"encoding/json"
	"errors"
	"os"

	typings "github.com/salesforceanton/go-cmd-client/internal"
)

const targetFilepath = "task_list.json"

var TaskList []typings.Task

func ReadTasklist() error {
	rawDataIn, err := os.ReadFile(targetFilepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawDataIn, &TaskList)
	if err != nil {
		return errors.New("task_list.json is invalid")
	}

	return nil
}

func SetResult(index int, result typings.Result) {
	TaskList[index].Result = result
}

func Save() error {
	rawDataOut, err := json.Marshal(&TaskList)
	if err != nil {
		return err
	}

	err = os.WriteFile(targetFilepath, rawDataOut, 0)
	if err != nil {
		return err
	}
	return nil
}

func RetainArgs(task typings.Task) []string {
	var result []string
	result = append(result, task.Name)

	for flag, value := range task.Params {
		result = append(result, flag)
		result = append(result, value)
	}

	return result
}
