package sercon

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	cueload "cuelang.org/go/cue/load"
)

type LoggingConfig struct {
	Level string `json:"level" env:"log_level"`
}

type MongodbOptions struct {
	ReplicaSet string `json:"replicaSet"`
}

type MongodbConfig struct {
	URI      string         `json:"uri" env:"mongodb_uri"`
	Username string         `json:"username" env:"mongodb_username"`
	Password string         `json:"password" env:"mongodb_password"`
	DbName   string         `json:"dbName" env:"mongodb_dbname"`
	Hosts    []string       `json:"hosts"`
	Options  MongodbOptions `json:"options"`
}

type PrometheusConfig struct {
	GRPCMetrics bool `json:"grpcMetrics"`
}

type GRPCFeatures struct {
	Reflection  bool `json:"reflection"`
	HealthCheck bool `json:"healthCheck"`
	Logging     bool `json:"logging"`
}

type GRPCConfig struct {
	Port     int          `json:"port" env:"grpc_port"`
	Features GRPCFeatures `json:"features"`
}

type HTTPConfig struct {
	Port int `json:"port" env:"http_port"`
}

type CueConfig struct {
	Name       string            `json:"name"`
	Version    string            `json:"version"`
	Env        string            `json:"env" env:"env"`
	Mode       string            `json:"mode" env:"mode"`
	Host       string            `json:"host" env:"host"`
	Logging    LoggingConfig     `json:"logging"`
	Mongodb    *MongodbConfig    `json:"mongodb,omitempty"`
	Prometheus *PrometheusConfig `json:"prometheus,omitempty"`
	GRPC       *GRPCConfig       `json:"grpc,omitempty"`
	HTTP       *HTTPConfig       `json:"http,omitempty"`
}

func BuildCue() {
	cuectx := cuecontext.New()
	ins := cueload.Instances([]string{"./cue/."}, &cueload.Config{
		ModuleRoot: "./cue",
	})
	res := cuectx.BuildInstance(ins[0])
	if res.Err() != nil {
		panic(res.Err())
	}

	serviceVal := res.LookupPath(cue.ParsePath("service"))
	if !serviceVal.Exists() {
		panic("service field not found")
	}

	var dec CueConfig
	if err := serviceVal.Decode(&dec); err != nil {
		panic(err)
	}

	if err := LoadEnvVars(&dec); err != nil {
		panic(err)
	}

	js, err := json.MarshalIndent(dec, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", string(js))
}

func LoadEnvVars(cfg *CueConfig) error {
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
