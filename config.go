package zaal

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

type RabbitMQConfig struct {
	URI string `json:"uri" env:"rabbitmq_uri"`
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

type Config struct {
	Name       string            `json:"name"`
	Title      string            `json:"title"`
	Version    string            `json:"version"`
	Env        string            `json:"env" env:"env"`
	Mode       string            `json:"mode" env:"mode"`
	Host       string            `json:"host" env:"host"`
	Logging    LoggingConfig     `json:"logging"`
	Mongodb    *MongodbConfig    `json:"mongodb,omitempty"`
	RabbiMQ    *RabbitMQConfig   `json:"rabbitmq,omitempty"`
	Prometheus *PrometheusConfig `json:"prometheus,omitempty"`
	GRPC       *GRPCConfig       `json:"grpc,omitempty"`
	HTTP       *HTTPConfig       `json:"http,omitempty"`
}
