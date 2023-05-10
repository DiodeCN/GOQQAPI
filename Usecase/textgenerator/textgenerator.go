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
	triggeredMilitary *ini.File
	mu           sync.Mutex
)

func init() {
	var err error
	triggeredQQs, err = ini.Load("Ability.ini")
	if err != nil {
		if os.IsNotExist(err) {
			triggeredQQs = ini.Empty()
			triggeredQQs.SaveTo("Ability.ini")
		} else {
			panic(err)
		}
	}

	triggeredMilitary, err = ini.Load("Military.ini")
	if err != nil {
		if os.IsNotExist(err) {
			triggeredMilitary = ini.Empty()
			triggeredMilitary.SaveTo("Military.ini")
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

	section, _ := triggeredQQs.NewSection(fmt.Sprintf("%d", qq))
	section.NewKey("date", time.Now().Format("2006-06-28"))
	section.NewKey("text", text)
	triggeredQQs.SaveTo("Ability.ini")

	return fmt.Sprintf("[CQ:at,qq=%d]%s", qq, text)
}

func GenerateMilitary(qq int64) string {
	seed := time.Now().UnixNano() + qq
	rand.Seed(seed)
	A := rand.Intn(7)
	B := rand.Intn(7)
	C := rand.Intn(7)
	D := rand.Intn(7)

	text := fmt.Sprintf("%d%d%d%d", A, B, C, D)

	mu.Lock()
	defer mu.Unlock()

	section, _ := triggeredMilitary.NewSection(fmt.Sprintf("%d", qq))
	section.NewKey("date", time.Now().Format("2006-06-28"))
	section.NewKey("text", text)
	triggeredMilitary.SaveTo("Military.ini")

	return fmt.Sprintf("[CQ:at,qq=%d]%s", qq, text)
}
