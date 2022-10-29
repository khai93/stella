package config

import (
	"strconv"

	envlib "github.com/khai93/stella/lib/env"
)

type Configuration struct {
	MemoryLimits int // Memory Limit for each execution in bytes
	Timeout      int // Maimum execution time before timing out (seconds)
	Port         int
}

var config *Configuration

func Get() (*Configuration, error) {
	if config != nil {
		return config, nil
	}

	var (
		memoryLimits = envlib.Getenv("STELLA_MEMORY_LIMITS", "104857600")
		timeout      = envlib.Getenv("STELLA_TIMEOUT", "5")
		port         = envlib.Getenv("STELLA_PORT", "4000")
	)

	memLimitsInt, err := strconv.Atoi(memoryLimits)
	if err != nil {
		panic(err)
	}

	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		panic(err)
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	config = &Configuration{
		MemoryLimits: memLimitsInt,
		Timeout:      timeoutInt,
		Port:         portInt,
	}

	return config, nil
}
