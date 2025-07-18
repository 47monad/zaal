package zaal_test

import (
	"os"
	"testing"

	"github.com/47monad/zaal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadEnvFile(t *testing.T) {
	t.Run("nonexistent_env/error", func(t *testing.T) {
		err := zaal.LoadEnvFile("nonexistent.env")
		assert.Error(t, err)
	})

	// Create a temporary .env file for testing
	t.Run("load_env/ok", func(t *testing.T) {
		content := `
		TEST_ENV=test
		`
		tmpFile, err := os.CreateTemp("", "test*.env")
		require.NoError(t, err)
		defer os.Remove(tmpFile.Name())

		_, err = tmpFile.WriteString(content)
		require.NoError(t, err)
		tmpFile.Close()

		err = zaal.LoadEnvFile(tmpFile.Name())
		assert.NoError(t, err)

		assert.Equal(t, "test", os.Getenv("TEST_ENV"))

		os.Unsetenv("TEST_ENV")
	})
}

func TestLoadEnvVars(t *testing.T) {
	// Setup helper function to reset env vars after each test
	resetEnvVars := func() {
		os.Unsetenv("ENV")
		os.Unsetenv("MODE")
		os.Unsetenv("HOST")
		os.Unsetenv("LOG_LEVEL")
		os.Unsetenv("MONGODB_URI")
		os.Unsetenv("MONGODB_USERNAME")
		os.Unsetenv("MONGODB_PASSWORD")
		os.Unsetenv("MONGODB_DBNAME")
		os.Unsetenv("RABBITMQ_URI")
		os.Unsetenv("MAIN_GRPC_PORT")
		os.Unsetenv("MAIN_GRPC_CLIENT_ADDRESS")
		os.Unsetenv("MAIN_HTTP_PORT")
	}

	t.Run("load_basic_vars/ok", func(t *testing.T) {
		defer resetEnvVars()

		os.Setenv("ENV", "test")
		os.Setenv("MODE", "debug")
		os.Setenv("HOST", "localhost")
		os.Setenv("LOG_LEVEL", "info")

		cfg := &zaal.Config{
			Name:    "test-app",
			Title:   "Test App",
			Version: "1.0.0",
			Logging: zaal.LoggingConfig{},
		}

		err := zaal.LoadEnvVars(cfg)
		require.NoError(t, err)

		assert.Equal(t, "test", cfg.Env)
		assert.Equal(t, "debug", cfg.Mode)
		assert.Equal(t, "localhost", cfg.Host)
		assert.Equal(t, "info", cfg.Logging.Level)
	})

	t.Run("load_nested_vars/ok", func(t *testing.T) {
		defer resetEnvVars()

		os.Setenv("LOG_LEVEL", "debug")
		os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
		os.Setenv("MONGODB_USERNAME", "testuser")
		os.Setenv("MONGODB_PASSWORD", "testpass")
		os.Setenv("MONGODB_DBNAME", "testdb")
		os.Setenv("POSTGRES_URI", "postgres://localhost:2134")

		cfg := &zaal.Config{
			Name:     "test-app",
			Title:    "Test App",
			Version:  "1.0.0",
			Logging:  zaal.LoggingConfig{},
			Mongodb:  &zaal.MongodbConfig{},
			Postgres: &zaal.PostgresConfig{},
		}

		err := zaal.LoadEnvVars(cfg)
		require.NoError(t, err)

		assert.Equal(t, "debug", cfg.Logging.Level)
		assert.Equal(t, "mongodb://localhost:27017", cfg.Mongodb.URI)
		assert.Equal(t, "testuser", cfg.Mongodb.Username)
		assert.Equal(t, "testpass", cfg.Mongodb.Password)
		assert.Equal(t, "testdb", cfg.Mongodb.DBName)
		assert.Equal(t, "postgres://localhost:2134", cfg.Postgres.URI)
	})

	t.Run("load_numeric_vars/ok", func(t *testing.T) {
		defer resetEnvVars()

		os.Setenv("MAIN_GRPC_PORT", "50051")
		os.Setenv("MAIN_HTTP_PORT", "8080")

		cfg := &zaal.Config{
			Name:    "test-app",
			Title:   "Test App",
			Version: "1.0.0",
			GRPC:    &zaal.GRPCConfig{Servers: map[string]zaal.GRPCServerConfig{"main": {}}},
			HTTP:    &zaal.HTTPConfig{Servers: map[string]zaal.HTTPServerConfig{"main": {}}},
		}

		err := zaal.LoadEnvVars(cfg)
		require.NoError(t, err)

		assert.Equal(t, 50051, cfg.GRPC.Servers["main"].Port)
		assert.Equal(t, 8080, cfg.HTTP.Servers["main"].Port)
	})

	t.Run("invalid_numeric_var/error", func(t *testing.T) {
		defer resetEnvVars()

		os.Setenv("MAIN_HTTP_PORT", "not-a-number")

		cfg := &zaal.Config{
			HTTP: &zaal.HTTPConfig{Servers: map[string]zaal.HTTPServerConfig{"main": {}}},
		}

		err := zaal.LoadEnvVars(cfg)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error converting env var")
	})

	// TODO: Test Boolean vars. currently no boolean env var exist

	t.Run("map_fields/ok", func(t *testing.T) {
		defer resetEnvVars()

		os.Setenv("SERVICE1_GRPC_CLIENT_ADDRESS", "localhost:50051")

		cfg := &zaal.Config{
			GRPC: &zaal.GRPCConfig{
				Clients: map[string]zaal.GRPCClientConfig{
					"service1": {},
				},
			},
		}

		err := zaal.LoadEnvVars(cfg)
		require.NoError(t, err)
		assert.Equal(t, "localhost:50051", cfg.GRPC.Clients["service1"].Address)
	})

	t.Run("full/ok", func(t *testing.T) {
		defer resetEnvVars()

		os.Setenv("ENV", "production")
		os.Setenv("MODE", "release")
		os.Setenv("HOST", "0.0.0.0")
		os.Setenv("LOG_LEVEL", "info")
		os.Setenv("MONGODB_URI", "mongodb://mongo:27017")
		os.Setenv("MONGODB_USERNAME", "produser")
		os.Setenv("MONGODB_PASSWORD", "prodpass")
		os.Setenv("MONGODB_DBNAME", "proddb")
		os.Setenv("RABBITMQ_URI", "amqp://guest:guest@rabbitmq:5672/")
		os.Setenv("MAIN_GRPC_PORT", "5000")
		os.Setenv("MAIN_HTTP_PORT", "8000")

		cfg := &zaal.Config{
			Name:    "prod-app",
			Title:   "Production App",
			Version: "1.0.0",
			Logging: zaal.LoggingConfig{},
			Mongodb: &zaal.MongodbConfig{},
			RabbiMQ: &zaal.RabbitMQConfig{},
			GRPC:    &zaal.GRPCConfig{Servers: map[string]zaal.GRPCServerConfig{"main": {}}},
			HTTP:    &zaal.HTTPConfig{Servers: map[string]zaal.HTTPServerConfig{"main": {}}},
		}

		err := zaal.LoadEnvVars(cfg)
		require.NoError(t, err)

		assert.Equal(t, "production", cfg.Env)
		assert.Equal(t, "release", cfg.Mode)
		assert.Equal(t, "0.0.0.0", cfg.Host)
		assert.Equal(t, "info", cfg.Logging.Level)
		assert.Equal(t, "mongodb://mongo:27017", cfg.Mongodb.URI)
		assert.Equal(t, "produser", cfg.Mongodb.Username)
		assert.Equal(t, "prodpass", cfg.Mongodb.Password)
		assert.Equal(t, "proddb", cfg.Mongodb.DBName)
		assert.Equal(t, "amqp://guest:guest@rabbitmq:5672/", cfg.RabbiMQ.URI)
		assert.Equal(t, 5000, cfg.GRPC.Servers["main"].Port)
		assert.Equal(t, 8000, cfg.HTTP.Servers["main"].Port)
	})

	t.Run("nil_pointer_in_config/ok", func(t *testing.T) {
		defer resetEnvVars()

		os.Setenv("MONGODB_URI", "mongodb://localhost:27017")

		cfg := &zaal.Config{
			Name:    "test-app",
			Title:   "Test App",
			Version: "1.0.0",
			// Mongodb is nil
		}

		err := zaal.LoadEnvVars(cfg)
		assert.NoError(t, err)
	})
}
