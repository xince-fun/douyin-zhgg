package snowflake

import (
	"github.com/godruoyi/go-snowflake"
	"time"
)

var startTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

func init() {
	snowflake.SetStartTime(startTime)
}

func ID() int64 {
	return int64(snowflake.ID())
}

func FromLowerTime(t time.Time) int64 {
	return t.Add(time.Millisecond).
		Sub(startTime).Milliseconds() <<
		int64(snowflake.MachineIDLength+snowflake.SequenceLength)
}

func FromUpperTime(t time.Time) int64 {
	return t.Sub(startTime).Milliseconds() << int64(snowflake.MachineIDLength+snowflake.SequenceLength)
}

func Time(id int64) time.Time {
	sid := snowflake.ParseID(uint64(id))
	return sid.GenerateTime()
}
