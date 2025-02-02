package utils_test

import (
	"log"
	"math"
	"os"
	"testing"
	"time"

	"github.com/rushairer/sso/utils"
)

func TestEnv(t *testing.T) {
	// 示例 1: 字符串类型
	os.Setenv("STRING_ENV", "hello")
	strVal := utils.GetEnv("STRING_ENV", "default")
	log.Println("STRING_ENV:", strVal) // 输出: hello

	// 示例 2: 整数类型
	os.Unsetenv("INT_ENV")
	intVal := utils.GetEnv("INT_ENV", 8080)
	log.Println("INT_ENV:", intVal) // 输出: 8080，并设置环境变量为 "8080"

	// 示例 3: 布尔类型
	os.Setenv("BOOL_ENV", "true")
	boolVal := utils.GetEnv("BOOL_ENV", false)
	log.Println("BOOL_ENV:", boolVal) // 输出: true

	// 示例 4: 浮点数类型
	os.Setenv("FLOAT_ENV", "3.14")
	floatVal := utils.GetEnv("FLOAT_ENV", 2.718)
	log.Println("FLOAT_ENV:", floatVal) // 输出: 3.14

	// 示例 5: 时间间隔类型
	os.Setenv("DURATION_ENV", "1h30m")
	durationVal := utils.GetEnv("DURATION_ENV", time.Hour)
	log.Println("DURATION_ENV:", durationVal) // 输出: 1h30m0s

	maxInt64Val := utils.GetEnv("MAX_INT64", math.MaxInt64)
	log.Println("MAX_INT64:", maxInt64Val)
}
