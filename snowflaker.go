package snowflaker

import "time"

const (
	SNOWFLAKER_TAKE_BITS_12 int64 = 0xfff         //  Take 12 bits with '1' in binary.
	SNOWFLAKER_TAKE_BITS_10 int64 = 0x3ff         //  Take 10 bits with '1' in binary.
	SNOWFLAKER_TAKE_BITS_41 int64 = 0x1ffffffffff // Take 41 bits with '1' in binary.
)

type Snowflaker struct {
	deviceId       int64
	sequenceNumber int64
	lastTimeMills  int64
}

func NewSnowflaker(iDeviceId int64) *Snowflaker {
	if iDeviceId < 0 || iDeviceId >= 1024 {
		iDeviceId = 0
	}
	return &Snowflaker{
		deviceId:       iDeviceId,
		sequenceNumber: 0,
		lastTimeMills:  currentMillisecond(),
	}
}

// GetId() 获取父级节点的父级节点
func (sf *Snowflaker) GetId() int64 {
	currentTimeMillis := currentMillisecond()
	if currentTimeMillis == sf.lastTimeMills {
		sf.sequenceNumber = sf.sequenceNumber + 1
	} else {
		sf.lastTimeMills = currentTimeMillis
		sf.sequenceNumber = 0
	}
	timeMillsBits := currentTimeMillis & SNOWFLAKER_TAKE_BITS_41 << 22
	deviceIdBits := sf.deviceId & SNOWFLAKER_TAKE_BITS_10 << 12
	sequenceNumberBits := sf.sequenceNumber & SNOWFLAKER_TAKE_BITS_12
	return timeMillsBits | deviceIdBits | sequenceNumberBits
}

func currentMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
