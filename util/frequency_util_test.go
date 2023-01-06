package util

import (
	"fmt"
	"testing"
	"time"
)

func TestFrequencyUtil_CheckFrequency(t *testing.T) {

	f := FrequencyUtil{
		Frequency:    5,
		OverflowType: OVER_FLOWTYPE_ERROR,
	}

	count := 0

	for i := 0; i <= 50; i++ {
		now := time.Now().String()
		err := f.CheckFrequency()
		if err != nil {
			fmt.Println(now, err.Error())
		} else {
			count++
			fmt.Println(now, "ok:", count)
		}
		time.Sleep(time.Millisecond * 100)
	}

}
