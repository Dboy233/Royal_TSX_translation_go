package util

import (
	"errors"
	"fmt"
	"time"
)

const (
	// OVER_FLOWTYPE_ERROR 返回error
	OVER_FLOWTYPE_ERROR = iota
	// OVER_FLOWTYPE_AWAIT 等待时间
	OVER_FLOWTYPE_AWAIT
)

type FrequencyUtil struct {
	//每秒调用的频率
	Frequency int
	//当frequency 频率达到上限的时候处理策略[OVER_FLOWTYPE_ERROR]返回error,[OVER_FLOWTYPE_AWAIT]等待时间
	OverflowType int
	//调用计数
	requestCount int
	//调用时间
	requestNowTime int64
}

func (f *FrequencyUtil) CheckFrequency() error {

	now := time.Now().UnixMilli()

	//时差大于1秒直接重置请求次数和请求时间
	timeDiff := now - f.requestNowTime
	if timeDiff >= 1000 {
		f.requestNowTime = now
		f.requestCount = 0
	}

	//在一秒以内频率超限处理
	if f.requestCount >= f.Frequency && timeDiff <= 1000 {

		if f.OverflowType == OVER_FLOWTYPE_ERROR {
			return errors.New("频率达到每秒上限")
		} else {
			var sleepTime = 1000 - timeDiff
			fmt.Println("频率达到上限，进行休眠补时 ", sleepTime)
			time.Sleep(time.Millisecond * time.Duration(sleepTime))
		}

	}

	//累加请求次数
	f.requestCount++

	return nil
}
