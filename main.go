package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	command := r.FormValue("command")
	username := r.FormValue("user_name")

	if command == "/command" {
		cmd := exec.Command("./script.sh", username)

		if err := cmd.Start(); err != nil {
			fmt.Printf("Failed to start script: %v", err)
			return
		}

		fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+": ran for %v \n", username)

		fmt.Fprint(w, "Command successful.")
	} else {
		fmt.Fprint(w, "Something went wrong.")
	}
}
