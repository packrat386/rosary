package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func configAddr() string {
	if val, ok := os.LookupEnv("ROSARY_ADDR"); ok {
		return val
	}

	return "localhost:6724"
}

func configBeat() time.Duration {
	if val, ok := os.LookupEnv("ROSARY_BEAT_MS"); ok {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Fatalf("FATAL ERROR invalid integer ROSARY_BEAT_MS: %s -> %s", val, err.Error())
		}

		return time.Duration(i) * time.Millisecond
	}

	return 400 * time.Millisecond
}

func configPollPeriod() time.Duration {
	if val, ok := os.LookupEnv("ROSARY_POLL_PERIOD_MS"); ok {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Fatalf("FATAL ERROR invalid integer ROSARY_POLL_PERIOD_MS: %s -> %s", val, err.Error())
		}

		return time.Duration(i) * time.Millisecond
	}

	return 100 * time.Millisecond
}
