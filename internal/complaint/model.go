package complaint

import "time"

type Complaint struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
