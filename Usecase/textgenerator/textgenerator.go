package textgenerator

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateAbility(qq int64) string {
	seed := time.Now().UnixNano() + qq
	rand.Seed(seed)
	randomNumber := rand.Intn(7*7*7) // 7的三次方
	A := randomNumber / 49
	B := (randomNumber % 49) / 7
	C := randomNumber % 7

	text := fmt.Sprintf("[CQ:at,qq=%d]%d%d%d", qq, A, B, C)
	return text
}
