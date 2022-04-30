package utils

import "time"

type stopWatch struct {
	initTimestamp int64
}

func NewStopWatch() stopWatch {
	return stopWatch{
		initTimestamp: time.Now().Unix(),
	}
}

func (s stopWatch) ElapsedTime() int64 {
	return time.Now().Unix() - s.initTimestamp
}
