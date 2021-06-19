package common

import (
	"math/rand"
	"os"
	"time"
)

// RandomString return an n bit string
func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// MkDir create dir if not exists
func MkDir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
	}
	return err
}

// MkDirAll create dir if not exists
func MkDirAll(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return err
}
