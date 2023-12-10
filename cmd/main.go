package main

import (
	"context"
	"fmt"
	"time"

	typings "github.com/salesforceanton/go-cmd-client/internal"
	"github.com/salesforceanton/go-cmd-client/internal/config"
	"github.com/salesforceanton/go-cmd-client/internal/logger"
	"github.com/salesforceanton/go-cmd-client/internal/parser"
	"github.com/salesforceanton/go-cmd-client/internal/worker"
)

func main() {
	// Set timer to get execution time
	start := time.Now()

	// Set logger config
	logger.SetConfiguration()

	// Init app configuration from env variables
	cfg, err := config.InitConfig()
	if err != nil {
		logger.LogError("Runtime", err)
		return
	}

	// Read task list file
	logger.LogInfo("Runtime", "Prepare Task List...")

	err = parser.ReadTasklist()
	if err != nil {
		logger.LogError("Runtime", err)
		return
	}

	// Process tasks and grab results
	logger.LogInfo("Runtime", "Tasks processing...")

	for i, task := range parser.TaskList {
		ctx, _ := context.WithTimeout(context.Background(), cfg.TaskTimeout)

		var result typings.Result

		if task.PositiveOutcome != nil && task.NegativeOutcome != nil {
			result = worker.RunConditionWithContext(ctx, cfg.BinPth, task)
		} else {
			result = worker.RunWithContext(
				ctx,
				typings.RunTaskInput{
					Bin:     cfg.BinPth,
					Command: task.Name,
					Params:  parser.RetainArgs(task.Name, task.Params),
				})
		}

		parser.SetResult(i, result)
	}

	// Save tasklist with execution results
	logger.LogInfo("Runtime", "All tasks have been done! Saving results...")

	parser.Save()

	elapsed := time.Since(start)
	logger.LogInfo("Runtime", fmt.Sprintf("Process complite for %s - find results in task_list.json!", elapsed))
}
