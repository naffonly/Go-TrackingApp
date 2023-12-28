package random

import (
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandomStc() string {
	length := 10
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	ranString := fmt.Sprintf("%s%s", dateNow(), string(b))
	return ranString
}

func dateNow() string {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	date := fmt.Sprintf("%d%d%d", year, month, day)
	result := date[2:]
	return result
}
