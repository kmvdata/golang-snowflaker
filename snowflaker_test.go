package snowflaker

import (
	"testing"
	"time"
)

func TestSnowflakerGetId(t *testing.T) {
	snowflaker := NewSnowflaker(0)
	// Test sequenceNumber
	println("==========")
	for i := 0; i < 10; i++ {
		println(snowflaker.GetId())
	}

	println("==========")
	// Test skip sequenceNumber
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(10) * time.Millisecond)
		println(snowflaker.GetId())
	}
}
