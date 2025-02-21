package zaal

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnvFile(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func LoadEnvVars(cfg *Config) error {
	val := reflect.ValueOf(cfg).Elem()
	return setFields(val)
}

func setFields(val reflect.Value) error {
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		fieldVal := val.Field(i)

		envTag := field.Tag.Get("env")
		if envTag != "" {
			tagName := strings.ToUpper(envTag)
			envValue := os.Getenv(tagName)
			if envValue != "" {
				switch fieldVal.Kind() {
				case reflect.String:
					fieldVal.SetString(envValue)
				case reflect.Int:
					intValue, err := strconv.Atoi(envValue)
					if err != nil {
						return fmt.Errorf("error converting env var %s to int: %w",
							envTag, err)
					}
					fieldVal.SetInt(int64(intValue))
				case reflect.Bool:
					boolValue, err := strconv.ParseBool(envValue)
					if err != nil {
						return fmt.Errorf("error converting env var %s to bool: %w",
							envTag, err)
					}
					fieldVal.SetBool(boolValue)
				}
			}
		} else if fieldVal.Kind() == reflect.Struct {
			// Recursively process nested structs
			if fieldVal.Type().String() != "time.Time" {
				setFields(fieldVal)
			}
		} else if fieldVal.Kind() == reflect.Ptr && !fieldVal.IsNil() &&
			fieldVal.Elem().Kind() == reflect.Struct {
			setFields(fieldVal.Elem())
		}
	}
	return nil
}
