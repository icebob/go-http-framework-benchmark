package util

import (
	"math/rand"
	"time"
)

type JsonReq struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type JsonRes struct {
	Action string `json:"action"`
	Num1   int    `json:"num1"`
	Num2   int    `json:"num2"`
	Result int    `json:"result"`
}

func Add(num1, num2 int) int {
	//time.Sleep(10 * time.Millisecond)
	return num1 + num2
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var methods = []string{"GET", "POST"}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateRandomRoutes(n int) []string {
	res := make([]string, n)
	for j := 0; j <= 100; j++ {
		path := "/" + RandString(20)
		res = append(res, path)
	}

	return res
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
