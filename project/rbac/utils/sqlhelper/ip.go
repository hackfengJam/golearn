package sqlhelper

import (
	"go-common-master/library/net/ip"
	"math"
)

const maxMask uint32 = 32

func MaskIntToString(mask uint32) string {
	// mask int to string  16 -> 255.255.0.0
	if mask > 32 {
		return ""
	}
	return ip.InetNtoA(uint32(math.Pow(2, float64(maxMask)) - math.Pow(2, float64(maxMask-mask))))
}

func MaskStringToInt(mask string) uint32 {
	// mask string to int  255.255.0.0 -> 16
	sum := ip.InetAtoN(mask)
	var max uint32 = 0x7fffffff + 1

	var maskInt uint32 = 0
	for {
		if sum&max == max {
			maskInt++
		} else {
			break
		}
		sum = sum << 1
	}
	return maskInt
}
