package service_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimee_GetDaysUntilMyDate(t *testing.T) {

	testCases2 := []struct {
		name           string
		expectDiffDate int32
		startTime      time.Time
		endTime        time.Time
		validTime      bool
	}{
		{
			name:           "Valid: 2 days remaining",
			expectDiffDate: 2,
			startTime:      time.Date(2000, 0, 0, 0, 0, 0, 0, time.UTC),
			endTime:        time.Date(2000, 0, 2, 0, 0, 0, 0, time.UTC),
			validTime:      true,
		},
		{
			name: "Invalid",
			// expectDiffDate: 0,
			startTime: time.Date(2000, 0, 2, 0, 0, 0, 0, time.UTC),
			endTime:   time.Date(2000, 0, 0, 0, 0, 0, 0, time.UTC),
			validTime: false,
		},
	}

	for _, tc := range testCases2 {

		t.Run(tc.name, func(t *testing.T) {
			result, err := DaysLeft(tc.startTime, tc.endTime)

			if tc.validTime {
				assert.NoError(t, err)
				assert.Equal(t, result, tc.expectDiffDate)
			} else {
				assert.Error(t, err)
			}

		})
	}
}
