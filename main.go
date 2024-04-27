package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

const logFilePath = "/var/log/random.log"

func main() {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	// Open the log file for writing
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Set log output to the file
	log.SetOutput(logFile)

	for {
		level := randomLogLevel()
		message := randomLogMessage()

		switch level {
		case "INFO":
			log.Printf("[INFO] %s\n", message)
		case "WARNING":
			log.Printf("[WARNING] %s\n", message)
		case "ERROR":
			log.Printf("[ERROR] %s\n", message)
		}

		time.Sleep(time.Duration(r.Intn(3)+1) * time.Second)
	}
}

func randomLogLevel() string {
	levels := []string{"INFO", "WARNING", "ERROR"}
	return levels[rand.Intn(len(levels))]
}

func randomLogMessage() string {
	messages := []string{
		"This is a random log message",
		"Another random log message",
		"Random log message number three",
	}
	return messages[rand.Intn(len(messages))]
}