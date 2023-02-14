package config

import (
	"strconv"

	envlib "github.com/khai93/stella/lib/env"
)

type Configuration struct {
	MemoryLimits int // Memory Limit for each execution in bytes
	Timeout      int // Maimum execution time before timing out (seconds)
	Port         int
	Workers      int
	Redis        RedisConfiguration
}

type RedisConfiguration struct {
	Address  string
	Password string
	DB       int
}

var config *Configuration

func Get() (*Configuration, error) {
	if config != nil {
		return config, nil
	}

	var (
		memoryLimits = envlib.Getenv("STELLA_MEMORY_LIMITS", "1009288000")
		timeout      = envlib.Getenv("STELLA_TIMEOUT", "5")
		workers      = envlib.Getenv("STELLA_WORKERS", "4")
		port         = envlib.Getenv("STELLA_PORT", "4000")
		redisAddr    = envlib.Getenv("REDIS_ADDRESS", "redis:6379")
		redisPass    = envlib.Getenv("REDIS_PASSWORD", "")
		redisDB      = envlib.Getenv("REDIS_DB", "0")
	)

	memLimitsInt, err := strconv.Atoi(memoryLimits)
	if err != nil {
		panic(err)
	}

	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		panic(err)
	}

	workersInt, err := strconv.Atoi(workers)
	if (err) != nil {
		panic(err)
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	redisDBInt, err := strconv.Atoi(redisDB)
	if err != nil {
		panic(err)
	}

	config = &Configuration{
		MemoryLimits: memLimitsInt,
		Timeout:      timeoutInt,
		Port:         portInt,
		Workers:      workersInt,
		Redis: RedisConfiguration{
			Address:  redisAddr,
			Password: redisPass,
			DB:       redisDBInt,
		},
	}

	return config, nil
}
