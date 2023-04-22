package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	client := &http.Client{}
	file, err := os.Open("usernames.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new file to write the available usernames
	outputFile, err := os.Create("available_usernames.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	startTime := time.Now()                                                                    // start the timer
	colors := []string{"\033[31m", "\033[33m", "\033[32m", "\033[36m", "\033[34m", "\033[35m"} // ANSI escape codes for rainbow colors
	colorIndex := 0
	for scanner.Scan() {
		wg.Add(1)
		go func(username string, colorIndex int) {
			defer wg.Done()
			var data = strings.NewReader(fmt.Sprintf(`[{"operationName":"UsernameValidator_User","variables":{"username":"%s"},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"fd1085cf8350e309b725cf8ca91cd90cac03909a3edeeedbd0872ac912f3d660"}}}]`, username))
			req, err := http.NewRequest("POST", "https://gql.twitch.tv/gql", data)
			if err != nil {
				log.Fatal(err)
			}
			req.Header.Set("Accept", "*/*")
			req.Header.Set("Accept-Language", "nl-NL")
			req.Header.Set("Client-Id", "kimne78kx3ncx6brgo4mv6wki5h1ko")
			req.Header.Set("Client-Integrity", "v4.public.eyJjbGllbnRfaWQiOiJraW1uZTc4a3gzbmN4NmJyZ280bXY2d2tpNWgxa28iLCJjbGllbnRfaXAiOiIyMTMuMTE4LjIxNi4yMTYiLCJkZXZpY2VfaWQiOiIxOEhxMUxJT1RhTWxNNG8ybnJ0OGpBWHhrSURUS2ZJRiIsImV4cCI6IjIwMjMtMDQtMjJUMjM6MTU6MjdaIiwiaWF0IjoiMjAyMy0wNC0yMlQwNzoxNToyN1oiLCJpc19iYWRfYm90IjoiZmFsc2UiLCJpc3MiOiJUd2l0Y2ggQ2xpZW50IEludGVncml0eSIsIm5iZiI6IjIwMjMtMDQtMjJUMDc6MTU6MjdaIiwidXNlcl9pZCI6IiJ9UASIvQKOqO2joFS-X1C9kjJxZTSe8-tXAIAqxEV0svfeF88taLCP4VpBGQ4yJnVt1QuA3KW7XwEGIiAektuFAA")
			req.Header.Set("Client-Session-Id", "3e0ee474af550dfa")
			req.Header.Set("Client-Version", "655b4ede-706b-4ab3-a6b4-7a6d86ccf8cb")
			req.Header.Set("Connection", "keep-alive")
			req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
			req.Header.Set("Origin", "https://www.twitch.tv")
			req.Header.Set("Referer", "https://www.twitch.tv/")
			req.Header.Set("Sec-Fetch-Dest", "empty")
			req.Header.Set("Sec-Fetch-Mode", "cors")
			req.Header.Set("Sec-Fetch-Site", "same-site")
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
			req.Header.Set("X-Device-Id", "18Hq1LIOTaMlM4o2nrt8jAXxkIDTKfIF")
			req.Header.Set("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
			req.Header.Set("sec-ch-ua-mobile", "?0")
			req.Header.Set("sec-ch-ua-platform", `"Windows"`)
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			bodyText, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			verib := "true"
			if strings.Contains(string(bodyText), verib) {
				fmt.Printf("%s[+] %s Available\033[0m\n", colors[colorIndex], username) // print the username in rainbow colors
				// Write the available username to the output file
				_, err := outputFile.WriteString(username + "\n")
				if err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Printf("%s[-]Taken\033[0m\n", colors[colorIndex])
			}
		}(scanner.Text(), colorIndex)
		colorIndex = (colorIndex + 1) % len(colors) // cycle through the rainbow colors
	}
	wg.Wait()
	elapsedTime := time.Since(startTime) // calculate the elapsed time
	fmt.Println("Elapsed time:", elapsedTime)
}
