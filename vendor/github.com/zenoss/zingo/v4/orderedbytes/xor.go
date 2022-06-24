package orderedbytes

import "math"

func xorbytes(arr []byte) {
	for i, b := range arr {
		arr[i] = xorbyte(b)
	}
}

func xorbyte(b byte) byte {
	return b ^ math.MaxUint8
}
