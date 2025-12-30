package config

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfig() (config Config, err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Warn("Error loading .env file, loading default config")
	}
	kvs := os.Environ()
	mp := map[string]string{}
	for _, kv := range kvs {
		sp := strings.SplitN(kv, "=", 2)
		mp[sp[0]] = sp[1]
	}
	dataBytes, err := json.Marshal(mp)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataBytes, &config)
	if err != nil {
		return
	}
	loadDefaults(&config)
	logLevel := logrus.InfoLevel
	if config.APP_DEBUG == "true" {
		logLevel = logrus.DebugLevel
	}
	logrus.SetLevel(logLevel)
	logrus.Info("Successfully loaded configurations.")
	return
}

func loadDefaults(v interface{}) {
	val := reflect.ValueOf(v).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if !field.IsZero() {
			continue // Skip if the field is already initialized
		}

		defaultTag := fieldType.Tag.Get("default")
		if defaultTag == "" {
			continue // Skip if there's no default tag
		}
		// Set the field's value based on its type
		switch field.Kind() {
		case reflect.String:
			field.SetString(defaultTag)
		case reflect.Int:
			if intValue, err := strconv.Atoi(defaultTag); err == nil {
				field.SetInt(int64(intValue))
			}
		case reflect.Bool:
			if boolValue, err := strconv.ParseBool(defaultTag); err == nil {
				field.SetBool(boolValue)
			}
		default:
			// You can handle other types here if needed
		}
	}
}
