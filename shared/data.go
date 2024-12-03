package shared

import "time"

type Task struct {
	ID        string
	Name      string
	Status    string
	CreatedAt time.Time
}
