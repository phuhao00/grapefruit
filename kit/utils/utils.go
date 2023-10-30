package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const keyChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetTimeString() string {
	now := time.Now()
	return fmt.Sprintf("%s%d", now.Format("20060102150405"), now.UnixNano()%1e9)
}

func GetRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	key := make([]byte, length)
	for i := 0; i < length; i++ {
		key[i] = keyChars[rand.Intn(len(keyChars))]
	}
	return string(key)
}

func MessageWithRequestId(message string, id string) string {
	return fmt.Sprintf("%s (request id: %s)", message, id)
}

func GenerateKey() string {
	rand.Seed(time.Now().UnixNano())
	key := make([]byte, 48)
	for i := 0; i < 16; i++ {
		key[i] = keyChars[rand.Intn(len(keyChars))]
	}
	uuid_ := GetUUID()
	for i := 0; i < 32; i++ {
		c := uuid_[i]
		if i%2 == 0 && c >= 'a' && c <= 'z' {
			c = c - 'a' + 'A'
		}
		key[i+16] = c
	}
	return string(key)
}

func GetUUID() string {
	code := uuid.New().String()
	code = strings.Replace(code, "-", "", -1)
	return code
}

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func GetOrDefault(env string, defaultValue int) int {
	if env == "" || os.Getenv(env) == "" {
		return defaultValue
	}
	num, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		return defaultValue
	}
	return num
}
