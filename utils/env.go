package utils

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

// 获取环境变量值，并设置默认值
func GetEnv[T any](key string, defaultValue T) T {
	valueStr, ok := os.LookupEnv(key)
	if ok {
		parsedValue, err := parseStringToT[T](valueStr)
		if err == nil {
			return parsedValue
		}
		// 解析失败，继续执行设置默认值的逻辑
	}

	// 设置默认值到环境变量并返回
	setDefaultEnv(key, defaultValue)
	return defaultValue
}

func parseStringToT[T any](s string) (T, error) {
	var t T
	rt := reflect.TypeOf(t)
	rv := reflect.ValueOf(&t).Elem()

	switch rt.Kind() {
	case reflect.String:
		rv.SetString(s)
		return t, nil

	case reflect.Bool:
		val, err := strconv.ParseBool(s)
		if err != nil {
			return t, err
		}
		rv.SetBool(val)
		return t, nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bitSize := 64
		switch rt.Kind() {
		case reflect.Int8:
			bitSize = 8
		case reflect.Int16:
			bitSize = 16
		case reflect.Int32:
			bitSize = 32
		}
		val, err := strconv.ParseInt(s, 10, bitSize)
		if err != nil {
			return t, err
		}
		rv.SetInt(val)
		return t, nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		bitSize := 64
		switch rt.Kind() {
		case reflect.Uint8:
			bitSize = 8
		case reflect.Uint16:
			bitSize = 16
		case reflect.Uint32:
			bitSize = 32
		}
		val, err := strconv.ParseUint(s, 10, bitSize)
		if err != nil {
			return t, err
		}
		rv.SetUint(val)
		return t, nil

	case reflect.Float32, reflect.Float64:
		bitSize := 64
		if rt.Kind() == reflect.Float32 {
			bitSize = 32
		}
		val, err := strconv.ParseFloat(s, bitSize)
		if err != nil {
			return t, err
		}
		rv.SetFloat(val)
		return t, nil

	default:
		// 处理 time.Duration 类型
		if rt.PkgPath() == "time" && rt.Name() == "Duration" {
			duration, err := time.ParseDuration(s)
			if err != nil {
				return t, err
			}
			rv.Set(reflect.ValueOf(duration))
			return t, nil
		}

		return t, fmt.Errorf("unsupported type: %s", rt.String())
	}
}

func setDefaultEnv[T any](key string, value T) {
	valueStr := fmt.Sprintf("%v", value)
	os.Setenv(key, valueStr)
}
