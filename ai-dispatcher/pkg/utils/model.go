package utils

import (
	"math"
	"strconv"
	"time"
)

func IsModelExpired(modelID string, granularity, modelMaxUsedTimes int64) bool {
	if modelID == "" {
		return false
	}
	modelCreateTime, err := strconv.ParseInt(modelID, 10, 64)
	if err != nil {
		return false
	}
	nowT := time.Now().Unix()
	delta := nowT - (modelCreateTime/int64(math.Pow(10, 6)) + granularity*modelMaxUsedTimes)
	return delta > 0
}
