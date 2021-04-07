# golang-snowflaker
Snowflaker(golang/snowflaker) is a golang project. It is used to generate primary key id with [Snowflake algorithm](https://en.wikipedia.org/wiki/Snowflake_ID).

```
[Sign Bit: 1] - [Timestamp: 41] - [Device ID: 10] - [Sequence Number: 12]

0 - 00000000 00000000 00000000 00000000 00000000 0 - 00000000 00 - 00000000 0000
```

## Introduction

Snowflaker is a kotlin project. It is used to generate primary key id with Snowflake algorithm.

The generated id is of type 'Long' (64 bits). It contains four parts:

- Part 1 (1 bit): this bit should always be 0 to ensure that the id is a positive number.
- Part 2 (41 bits): the current system time in milliseconds.
- Part 3 (10 bits): the device id.
- Part 4 (12 bits): the sequence number in the same millisecond.

## How to?

You can find the test code in [snowflaker_test.go][1]

```golang
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
```


[1]: https://github.com/kmvdata/golang-snowflaker/blob/main/snowflaker_test.go