package enums

type TaskState int

const (
	Todo TaskState = iota
	InProgress
	Done
)

var taskName = map[TaskState]string{
	Todo:       "todo",
	InProgress: "in progress",
	Done:       "done",
}

func (ts TaskState) String() string {
	return taskName[ts]
}
