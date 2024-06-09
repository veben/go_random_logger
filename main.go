package main

import (
	"fmt"
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
		message := randomSentence()

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

func randomSentence() string {
	var subjects = []string{
		"The cat", "A dog", "The bird", "A person", "An alien", "The elephant", "A mouse",
		"The lion", "A tiger", "The monkey", "A fish", "The bear", "An owl", "The rabbit",
		"A kangaroo", "The horse", "A squirrel", "The fox", "A snake", "The spider",
		"A dolphin", "The penguin", "A whale", "The frog", "A wolf", "The deer", "A bat",
		"The giraffe", "A rhino", "The hippo", "A crocodile", "The eagle", "A parrot",
		"The panda", "A flamingo", "The hedgehog", "A tortoise", "The lizard", "A beaver",
		"The otter", "A peacock", "The camel", "A lobster", "The seal", "A seagull",
		"The ant", "A beetle", "The jellyfish", "A starfish", "The shrimp", "A crab",
	}
	var verbs = []string{
		"eats", "jumps over", "runs to", "looks at", "destroys", "chases", "catches",
		"hides from", "watches", "climbs", "flies over", "dives into", "swims in",
		"dances around", "sings to", "hunts", "crawls under", "barks at", "sneaks up on",
		"rests near", "grazes on", "plays with", "leaps over", "rolls in", "sniffs",
		"chews on", "nibbles at", "pecks at", "glides across", "slithers through",
		"lurks in", "perches on", "explores", "discovers", "ventures into", "pounces on",
		"grabs", "nudges", "sprints to", "ambles toward", "creeps under", "scuttles over",
		"hovers above", "dives down", "emerges from", "gallops past", "trots beside",
		"twirls around", "swirls in", "weaves through", "flutters around",
	}
	var objects = []string{
		"a mouse", "the fence", "a tree", "the moon", "a car", "the river", "a banana",
		"the mountain", "a cave", "the ocean", "a flower", "the grass", "a nest",
		"the clouds", "a shadow", "the stars", "a bush", "the rock", "a hill", "the pond",
		"a branch", "the meadow", "a leaf", "the waterfall", "a stream", "the forest",
		"a log", "the field", "a berry", "the hive", "a burrow", "the trail", "a cliff",
		"the horizon", "a puddle", "the ice", "a snowflake", "the desert", "a dune",
		"the shore", "a pebble", "the sand", "a wave", "the tide", "a shell", "the coral",
		"a reef", "the lagoon", "a cave", "the jungle", "a canopy", "the vines", "a boulder",
	}
	var adverbs = []string{
		"quickly", "slowly", "gracefully", "angrily", "happily", "quietly", "loudly",
		"carefully", "recklessly", "sneakily", "boldly", "elegantly", "furiously",
		"peacefully", "suddenly", "deliberately", "mysteriously", "gently", "vigorously",
		"silently", "briskly", "joyfully", "reluctantly", "curiously", "intently",
		"absentmindedly", "casually", "earnestly", "fretfully", "gleefully", "heartily",
		"hurriedly", "lazily", "menacingly", "playfully", "thoughtfully", "wildly",
		"zealously", "tactfully", "nonchalantly", "cheerfully", "calmly", "diligently",
		"frantically", "patiently", "relentlessly", "shyly", "tenderly", "warily",
		"enthusiastically", "timidly", "gleamingly", "bustlingly", "voraciously",
	}

	// Select random elements from each slice
	subject := subjects[rand.Intn(len(subjects))]
	verb := verbs[rand.Intn(len(verbs))]
	object := objects[rand.Intn(len(objects))]
	adverb := adverbs[rand.Intn(len(adverbs))]

	// Construct the sentence
	sentence := fmt.Sprintf("%s %s %s %s.", subject, verb, object, adverb)
	return sentence
}