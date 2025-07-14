package datetime_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/unvurn/core/datetime"
)

type dummySpan struct {
	Start time.Time
	End   *time.Time
}

func TestDateFilter(t *testing.T) {
	span := &dummySpan{Start: time.Now().Add(-10 * time.Minute)}

	result := datetime.Between(span.Start, span.End, time.Now())

	assert.Equal(t, result, true)
}

func TestDateFilterBetween(t *testing.T) {
	end := time.Now().Add(3 * time.Minute)
	span := &dummySpan{Start: time.Now().Add(-10 * time.Minute), End: &end}

	result := datetime.Between(span.Start, span.End, time.Now())

	assert.Equal(t, result, true)
}
