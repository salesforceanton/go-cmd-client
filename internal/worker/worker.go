package worker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	typings "github.com/salesforceanton/go-cmd-client/internal"
	"github.com/salesforceanton/go-cmd-client/internal/logger"
)

func Run(input typings.RunTaskInput) typings.Result {
	logger.LogInfo("Task execution", fmt.Sprintf("Task [%s] start...", input.Command))

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	path, err := exec.LookPath(input.Bin)
	if err != nil {
		logger.LogInfo("Task execution", fmt.Sprintf("Task [%s] finished with Error!", input.Command))
		return prepareResult("", err)
	}

	cmd := exec.Command(path, input.Params...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()

	if len(stderr.String()) != 0 {
		errMsg := strings.Trim(stderr.String(), "\n")
		logger.LogInfo("Task execution", fmt.Sprintf("Task [%s] finished with Error!", input.Command))
		return prepareResult("", errors.New(errMsg))
	}

	logger.LogInfo("Task execution", fmt.Sprintf("Task [%s] finished Successfully!", input.Command))
	return prepareResult(strings.Trim(stdout.String(), "\n"), nil)
}

func RunWithContext(ctx context.Context, input typings.RunTaskInput) typings.Result {
	resultChan := make(chan typings.Result)

	go func() {
		resultChan <- Run(input)
	}()

	select {
	case <-ctx.Done():
		return typings.Result{
			Status:  typings.Error,
			Details: "Task time to execute has been exceed",
		}
	case result := <-resultChan:
		return result
	}
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
