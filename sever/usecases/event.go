package usecases

import (
	"gitaction/sever/domain"
)

type AddInMemoryRepository interface {
	Add(domain.Event) error
}

type AddEventInMemory struct {
	repository AddInMemoryRepository
}

func (usecase *AddEventInMemory) Save(e domain.Event) error {
	err := usecase.repository.Add(e)
	return err
}

func NewAddEventInMemory(r AddInMemoryRepository) *AddEventInMemory {
	return &AddEventInMemory{repository: r}
}
