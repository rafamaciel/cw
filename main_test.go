package main

import (
	//"fmt"
	"github.com/stretchr/testify/assert" //"reflect"
	"testing"
	"time"
)

func TesttimestampToTime(t *testing.T) {
	assert := assert.New(t)

	a := "2017-03-12"
	assert.Equal(timestampToTime(&a), time.Date(2017, 3, 12, 0, 0, 0, 0, time.UTC),
		"wrong parsing for input %s", a)

	a = "2017-03-12T18"
	assert.Equal(timestampToTime(&a), time.Date(2017, 3, 12, 18, 0, 0, 0, time.UTC),
		"wrong parsing for input %s", a)

	a = "2017-03-12T18:22"
	assert.Equal(timestampToTime(&a), time.Date(2017, 3, 12, 18, 22, 0, 0, time.UTC),
		"wrong parsing for input %s", a)

	a = "2017-03-12T18:22:23"
	assert.Equal(timestampToTime(&a), time.Date(2017, 3, 12, 18, 22, 23, 0, time.UTC),
		"wrong parsing for input %s", a)

	a = "18"
	y, m, d := time.Now().Date()
	assert.Equal(timestampToTime(&a), time.Date(y, m, d, 18, 0, 0, 0, time.UTC),
		"wrong parsing for input %s", a)

	a = "18:31"
	y, m, d = time.Now().Date()
	assert.Equal(timestampToTime(&a), time.Date(y, m, d, 18, 31, 0, 0, time.UTC),
		"wrong parsing for input %s", a)
}

func TestHumanReadableTimeToTime(t *testing.T) {
	assert := assert.New(t)

	s := "32h"
	dd, _ := time.ParseDuration(s)
	x := time.Now().UTC().Add(-dd)

	y, m, d := x.Date()

	assert.Equal(timestampToTime(&s), time.Date(y, m, d, x.Hour(), x.Minute(), 0, 0, time.UTC), "wrong parsing for input %s", s)

	s = "50m"
	dd, _ = time.ParseDuration(s)
	x = time.Now().UTC().Add(-dd)

	y, m, d = x.Date()

	assert.Equal(timestampToTime(&s), time.Date(y, m, d, x.Hour(), x.Minute(), 0, 0, time.UTC), "wrong parsing for input %s", s)

	s = "2h30m"
	dd, _ = time.ParseDuration(s)
	x = time.Now().UTC().Add(-dd)

	y, m, d = x.Date()

	assert.Equal(timestampToTime(&s), time.Date(y, m, d, x.Hour(), x.Minute(), 0, 0, time.UTC), "wrong parsing for input %s", s)
}
