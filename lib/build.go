package sercon

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func WriteToJson(c *CueConfig, outputPath string) error {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	dirPath := filepath.Dir(outputPath)
	if dirPath != "." && dirPath != "/" {
		os.MkdirAll(dirPath, os.ModePerm)
	}
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}
