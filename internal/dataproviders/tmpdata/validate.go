package tmpdata

import "github.com/google/uuid"

func (s *Storage) AddToValidate(ids []uuid.UUID) {
	s.toValidate.Push(ids)
}

func (s *Storage) ValidateList() []uuid.UUID {
	return s.toValidate.Pop()
}
