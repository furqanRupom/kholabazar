package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var config Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWTSecret   string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env variables :", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is Required !")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is Required !")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port is Required !")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port number should be a Number!")
	}
	JwtSecret := os.Getenv("JWT_SECRET")
	config = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JWTSecret: JwtSecret,
	}
}

func GetConfig() Config {
	loadConfig()
	return config
}
