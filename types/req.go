package types

type ReqPostTask struct {
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	LogDir       string `json:"log_dir"`
	CronInterval int    `json:"cron_interval"`
}

type ReqPutTask struct {
	Task
}
