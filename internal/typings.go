package typings

type Status string

const (
	Success Status = "SUCCESS"
	Error   Status = "ERROR"
)

type Result struct {
	Status  Status `json:"status"`
	Details string `json:"details"`
}

type Task struct {
	Name   string            `json:"name"`
	Params map[string]string `json:"params"`
	Result Result            `json:"result"`
}

type RunTaskInput struct {
	Bin     string
	Command string
	Params  []string
}
