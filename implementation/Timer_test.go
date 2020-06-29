package implementation_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimer(t *testing.T) {
	var t implementation.Timer
	now := time.Now().Unix()
	expectedNow := t.CurrentTimeStamp()
	assert.True(t, expectedNow >= now)
}
