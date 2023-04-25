package service

import (
	"errors"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft(endTime time.Time) (int32, error) {
	difference := time.Until(endTime) // endTime.Sub(startTime)
	days := difference.Hours() / 24

	if days < 0 {
		return 0, errors.New("sql: Rows are closed")
	}

	return int32(days), nil
}
