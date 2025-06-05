package main

import (
	"fmt"
	"net/http"
)

// Add these handler functions before your main() function

func teacherLoginHandler(w http.ResponseWriter, r *http.Request) {

}

func teacherDashboardHandler(w http.ResponseWriter, r *http.Request) {

}

func createWheelHandler(w http.ResponseWriter, r *http.Request) {

}

func wheelManageHandler(w http.ResponseWriter, r *http.Request) {

}

func studentWheelHandler(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	// Serve static files (HTML, CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Home page - serve the HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Login page - coming soon!")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Users: teacher1, student1")
	})

	// Teacher routes
	http.HandleFunc("/teacher/login", teacherLoginHandler)
	http.HandleFunc("/teacher/dashboard", teacherDashboardHandler)
	http.HandleFunc("/teacher/wheel/create", createWheelHandler)
	http.HandleFunc("/teacher/wheel/", wheelManageHandler) // handles /wheel/123/edit, /wheel/123/delete

	// Student routes (public, no auth)
	http.HandleFunc("/wheel/", studentWheelHandler) // handles /wheel/abc123

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
