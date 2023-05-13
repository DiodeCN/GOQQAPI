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
	triggeredQQs      *ini.File
	triggeredMilitary *ini.File
	mu                sync.Mutex
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

func ReadIni(qq int64, filename string) (string, error) {
	iniFile, err := ini.Load(filename)
	if err != nil {
		return "", err
	}

	section, err := iniFile.GetSection(fmt.Sprintf("%d", qq))
	if err != nil {
		return "", err
	}

	dateKey, err := section.GetKey("date")
	if err != nil {
		return "", err
	}

	if dateKey.String() == time.Now().Format("2006-01-02") {
		textKey, err := section.GetKey("text")
		if err != nil {
			return "", err
		}
		return textKey.String(), nil
	}

	return "", nil
}

func GenerateAbility(qq int64) string {
	seed := time.Now().UnixNano() + qq
	rand.Seed(seed)
	administrative := rand.Intn(7)
	diplomatic := rand.Intn(7)
	military := rand.Intn(7)

	text, err := ReadIni(qq, "Ability.ini")
	if err == nil && text != "" {
		return fmt.Sprintf("主人今天已经查过啦，\r\n"+"[CQ:at,qq=%d]%s", qq, text)
	}

	text = fmt.Sprintf("行政：%d 外交：%d 军事：%d", administrative, diplomatic, military)

	mu.Lock()
	defer mu.Unlock()

	section, _ := triggeredQQs.NewSection(fmt.Sprintf("%d", qq))
	section.NewKey("date", time.Now().Format("2006-01-02"))
	section.NewKey("text", text)
	triggeredQQs.SaveTo("Ability.ini")

	return fmt.Sprintf("[CQ:at,qq=%d]%s", qq, text)
}

func GenerateMilitary(qq int64) string {
	seed := time.Now().UnixNano() + qq
	rand.Seed(seed)
	firepower := rand.Intn(7)
	assault := rand.Intn(7)
	mobility := rand.Intn(7)
	siege := rand.Intn(7)

	text, err := ReadIni(qq, "Military.ini")
	if err == nil && text != "" {
		return fmt.Sprintf("主人今天已经查过啦，\r\n"+"[CQ:at,qq=%d]%s", qq, text)
	}

	text = fmt.Sprintf("火力:%d 冲击:%d 机动:%d 围城:%d", firepower, assault, mobility, siege)

	mu.Lock()
	defer mu.Unlock()

	section, _ := triggeredMilitary.NewSection(fmt.Sprintf("%d", qq))
	section.NewKey("date", time.Now().Format("2006-01-02"))
	section.NewKey("text", text)
	triggeredMilitary.SaveTo("Military.ini")

	return fmt.Sprintf("[CQ:at,qq=%d]%s", qq, text)
}
