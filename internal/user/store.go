package user

import (
	"github.com/google/uuid"
	"sync"
)

type UserStore interface {
	GetUser(id uuid.UUID) (User, error)
	GetUsers() ([]User, error)
	CreateUser() (User, error)
	UpdateUser(id uuid.UUID) error
	DeleteUser(id uuid.UUID) error
}

type MemoryStore struct {
	users []User
	mutex sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: make([]User, 0),
	}
}

func (s *MemoryStore) GetUser(id uuid.UUID) (User, error) {
	return User{}, nil
}

func (s *MemoryStore) GetUsers() ([]User, error) {
	return s.users, nil
}

func (s *MemoryStore) CreateUser() (User, error) {
	return User{}, nil
}

func (s *MemoryStore) UpdateUser(id uuid.UUID) error {
	return nil
}

func (s *MemoryStore) DeleteUser(id uuid.UUID) error {
	return nil
}
