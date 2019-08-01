/*
 * Copyright (c) 2015 Łukasz S.
 * Distributed under the terms of GPL-2 License.
 */

package src

import (
	"math/rand"
)

func BoolToUint8(b bool) uint8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func RandUint8() uint8 {
	return uint8(rand.Intn(256))
}
