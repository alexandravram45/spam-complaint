package complaint

import (
	"encoding/json"
	"os"
)

type Repository struct {
	filePath string
}

func NewRepository(path string) *Repository {
	return &Repository{filePath: path}
}

func (r *Repository) Save(complaints []Complaint) error {
	data, err := json.MarshalIndent(complaints, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}
