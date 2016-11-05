package main

import (
	"reflect"
	"testing"
	"time"
)

func TestStartTimer(t *testing.T) {
	var timer *time.Timer
	testTimer := StartTimer(5)
	if reflect.TypeOf(testTimer) != reflect.TypeOf(timer) {
		t.Error("Start Timer did not create a new timer")
	}
}
