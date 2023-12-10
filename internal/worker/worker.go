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
	"github.com/salesforceanton/go-cmd-client/internal/parser"
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

func RunCondition(bin string, condition typings.Task) typings.Result {
	result := Run(typings.RunTaskInput{
		Bin:     bin,
		Command: condition.Name,
		Params:  parser.RetainArgs(condition.Name, condition.Params),
	})

	outcome := isPositiveResult(result.Details)
	if outcome {
		return Run(typings.RunTaskInput{
			Bin:     bin,
			Command: condition.PositiveOutcome.Name,
			Params:  parser.RetainArgs(condition.PositiveOutcome.Name, condition.PositiveOutcome.Params),
		})
	}
	return Run(typings.RunTaskInput{
		Bin:     bin,
		Command: condition.NegativeOutcome.Name,
		Params:  parser.RetainArgs(condition.NegativeOutcome.Name, condition.NegativeOutcome.Params),
	})
}

func RunConditionWithContext(ctx context.Context, bin string, task typings.Task) typings.Result {
	resultChan := make(chan typings.Result)

	go func() {
		resultChan <- RunCondition(bin, task)
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

func isPositiveResult(details string) bool {
	return strings.Contains(details, "True")
}
