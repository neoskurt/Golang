package commands

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
}

var Tasks []Task