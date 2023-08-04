package utils

import "strconv"

func ByteToMB(size int64) string {
	const megabyte = 0.000001
	return strconv.FormatFloat(float64(size) * megabyte, 'f', 4, 32) + "MB"
}