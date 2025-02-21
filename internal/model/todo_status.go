package model

type TodoStatus int

const (
	StatusPending TodoStatus = iota
	StatusInProgress
	StatusPaused
	StatusCompleted
	StatusCanceled
)

func (s TodoStatus) String() string {
	return [...]string{
		"PENDING",
		"IN_PROGRESS",
		"PAUSED",
		"COMPLETED",
		"CANCELED",
	}[s]
}

func (s TodoStatus) IsValid() bool {
	return s >= StatusPending && s <= StatusCanceled
}
