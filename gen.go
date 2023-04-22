package main

import (
	"bufio"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	MaxUsernames = 1000000 // Maximum number of usernames to generate
	BatchSize    = 1000    // Number of usernames to write to the file in each batch
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a new file to write the usernames to
	file, err := os.Create("usernames.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered writer to write the usernames to the file in batches
	writer := bufio.NewWriter(file)

	// Use a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Generate and write usernames to the file concurrently
	for {
		if countUsernames() >= MaxUsernames {
			break
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < BatchSize; j++ {
				username := generateUsername()
				_, err := writer.WriteString(username + "\n")
				if err != nil {
					panic(err)
				}
			}
		}()

		// Flush the buffer every 1000 usernames to improve performance
		if rand.Intn(1000) == 0 {
			err = writer.Flush()
			if err != nil {
				panic(err)
			}
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Flush the buffer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

func generateUsername() string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 4)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func countUsernames() int {
	file, err := os.Open("usernames.txt")
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}
