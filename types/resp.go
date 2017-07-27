package types

type Resp struct {
	Success bool   `json:"success"`
	Errmsg  string `json:"errmsg"`
}

type RespGetTask struct {
	Resp
	Task Task `json:"task,omitempty"`
}

type RespGetTasks struct {
	Resp
	Tasks []Task `json:"tasks,omitempty"`
}

type RespPutTask struct {
	Resp
	Task Task `json:"task,omitempty"`
}

type RespStartTask struct {
	Resp
}

type RespStopTask struct {
	Resp
}

type RespPostTask struct {
	Resp
	Task *Task `json:"task,omitempty"`
}

type RespDeleteTask struct {
	Resp
}
