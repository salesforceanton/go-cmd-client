package worker

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"

	typings "github.com/salesforceanton/go-cmd-client/internal"
)

func Run(input typings.RunTaskInput) typings.Result {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	path, err := exec.LookPath(input.Bin)
	if err != nil {
		return prepareResult("", err)
	}

	args := []string{input.Command}
	copy(input.Params, args)

	cmd := exec.Command(path, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()

	if len(stderr.String()) != 0 {
		errMsg := strings.TrimRight(stderr.String(), "\n")
		return prepareResult("", errors.New(errMsg))
	}

	return prepareResult(strings.TrimRight(stdout.String(), "\n"), nil)
}

func prepareResult(result string, err error) typings.Result {
	if err != nil {
		return typings.Result{
			Status:  typings.Error,
			Details: err.Error(),
		}
	}

	return typings.Result{
		Status:  typings.Success,
		Details: result,
	}

}
