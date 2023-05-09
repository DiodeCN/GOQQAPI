package textgenerator

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

var (
	triggeredQQs *ini.File
	mu           sync.Mutex
)

func init() {
	var err error
	triggeredQQs, err = ini.Load("triggeredQQs.ini")
	if err != nil {
		if os.IsNotExist(err) {
			triggeredQQs = ini.Empty()
			triggeredQQs.SaveTo("triggeredQQs.ini")
		} else {
			panic(err)
		}
	}
}

func GenerateAbility(qq int64) string {
	seed := time.Now().UnixNano() + qq
	rand.Seed(seed)
	randomNumber := rand.Intn(7*7*7) // 7的三次方
	A := randomNumber / 49
	B := (randomNumber % 49) / 7
	C := randomNumber % 7

	text := fmt.Sprintf("%d%d%d", A, B, C)

	mu.Lock()
	defer mu.Unlock()

	section := triggeredQQs.Section("Triggered")
	key, err := section.GetKey(fmt.Sprintf("%d", qq))
	if err == nil {
		text = "111" + text
	} else {
		section.NewKey(fmt.Sprintf("%d", qq), "")
		triggeredQQs.SaveTo("triggeredQQs.ini")
		text = "222" + text
	}

	return fmt.Sprintf("[CQ:at,qq=%d]%s", qq, text)
}
