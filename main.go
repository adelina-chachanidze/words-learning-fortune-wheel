package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Add these handler functions before your main() function

func teacherLoginHandler(w http.ResponseWriter, r *http.Request) {

}

func teacherDashboardHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/teacher-dashboard.html")
}

func createWheelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		wheelName := r.FormValue("wheel_name")
		wordsText := r.FormValue("words")

		wheelID := generateWheelID()
		words := strings.Split(wordsText, "\n")

		err1 := saveWheelToCSV(wheelID, wheelName)
		err2 := saveWordsToCSV(wheelID, words)

		if err1 != nil || err2 != nil {
			fmt.Fprintf(w, "Error saving wheel!")
			return
		}

		// Redirect back to dashboard
		http.Redirect(w, r, "/static/teacher-dashboard.html", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "static/create-wheel.html")
	}
}

func wheelManageHandler(w http.ResponseWriter, r *http.Request) {

}

func studentWheelHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/student-wheel.html")
}

// Add these helper functions to main.go

func generateWheelID() string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 6)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func saveWheelToCSV(wheelID, wheelName string) error {
	file, err := os.OpenFile("data/wheels.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write([]string{wheelID, wheelName, time.Now().Format("2006-01-02")})
}

func saveWordsToCSV(wheelID string, words []string) error {
	file, err := os.OpenFile("data/words.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, word := range words {
		if strings.TrimSpace(word) != "" {
			err := writer.Write([]string{wheelID, strings.TrimSpace(word), "0", "0"})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// API endpoint to get wheels data
func getWheelsAPI(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("data/wheels.csv")
	if err != nil {
		json.NewEncoder(w).Encode([]string{})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

// API endpoint to get words for a wheel
func getWordsAPI(w http.ResponseWriter, r *http.Request) {
	wheelID := strings.TrimPrefix(r.URL.Path, "/api/words/")

	file, err := os.Open("data/words.csv")
	if err != nil {
		json.NewEncoder(w).Encode([]string{})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	var words []string
	for _, record := range records {
		if len(record) >= 2 && record[0] == wheelID {
			words = append(words, record[1])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Teacher routes
	http.HandleFunc("/teacher/dashboard", teacherDashboardHandler)
	http.HandleFunc("/teacher/wheel/create", createWheelHandler)

	// API routes
	http.HandleFunc("/api/wheels", getWheelsAPI)
	http.HandleFunc("/api/words/", getWordsAPI)

	// Student routes
	http.HandleFunc("/wheel/", studentWheelHandler)

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
