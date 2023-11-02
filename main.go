package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Activity struct {
	Email               string `json:"email"`
	ActivityDescription string `json:"activityDescription"`
	Cause               string `json:"cause"`
	CreateDate          string `json:"createDate"`
	CategoryType        string `json:"categoryType"`
	ActivityType        string `json:"activityType"`
}

func main() {
	// Create or open the file for writing
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	file.WriteString("[")

	// Generate and write random activities to the file
	for i := 0; i < 10; i++ { // Change the number of activities as needed
		activity := generateRandomActivity()
		if i > 0 {
			file.WriteString(",") // Add a comma for separating JSON objects after the first one
		}
		if err := writeActivityToFile(file, activity); err != nil {
			fmt.Println("Error writing JSON to file:", err)
			return
		}
	}

	// Add "]" to the end of the file
	file.WriteString("]")

	fmt.Println("Random JSON data has been written to output.json")
}

// generateRandomActivity generates a random Activity struct
func generateRandomActivity() Activity {
	randomName := getRandomElement([]string{"john", "alice", "bob", "jane", "david"})
	randomActivity := getRandomElement([]string{"create", "update", "delete"})
	randomCategoryType := getRandomElement([]string{"master", "sub"})
	cause := getRandomElement([]string{"", "", "some cause"})
	email := generateRandomEmail(randomName)
	activity := generateActivityDescription(randomName, randomActivity)
	timestamp := generateRandomTimestamp("2023-11-01", "2023-11-03")

	return Activity{
		Email:               email,
		ActivityDescription: activity,
		Cause:               cause,
		CreateDate:          timestamp,
		CategoryType:        randomCategoryType,
		ActivityType:        randomActivity,
	}
}

// generateRandomString generates a random string of the specified length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// writeActivityToFile writes an Activity to a JSON file
func writeActivityToFile(file *os.File, activity Activity) error {
	encoder := json.NewEncoder(file)
	return encoder.Encode(activity)
}

func generateRandomEmail(name string) string {
	part2 := getRandomElement([]string{"gmail", "yahoo", "hotmail", "example", "testmail"})
	part3 := getRandomElement([]string{"com", "net", "org", "io", "mail"})

	return name + "@" + part2 + "." + part3
}

func generateActivityDescription(name, activity string) string {
	return name + " have " + activity + " something.."
}

func getRandomElement(words []string) string {
	if len(words) == 0 {
		return ""
	}
	return words[rand.Intn(len(words))]
}

func generateRandomTimestamp(startDate, endDate string) string {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return "" // Handle the error
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return "" // Handle the error
	}

	// Calculate the time difference between start and end dates
	diff := end.Sub(start)

	// Generate a random duration within the date range
	randomDuration := time.Duration(rand.Int63n(int64(diff)))

	// Add the random duration to the start time to get the random timestamp
	randomTime := start.Add(randomDuration)

	// Format the random time in the desired format
	return randomTime.Format(time.RFC3339)
}
