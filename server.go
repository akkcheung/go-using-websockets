package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var todoList []string
var completeTodoList []string

func getCmd(input string) string {
	inputArr := strings.Split(input, " ")

	return inputArr[0]
}

func getMessage(input string) string {
	inputArr := strings.Split(input, " ")
	var result string
	for i := 1; i < len(inputArr); i++ {
		result += inputArr[i]
	}

	return result
}

func updateTodoList(input string) {
	tmpList := todoList
	todoList = []string{}

	//tmpCompleteList := completeTodoList
	//completeTodoList = []string{}

	for _, val := range tmpList {
		if val == input {
			completeTodoList = append(completeTodoList, val)
			continue
		}

		todoList = append(todoList, val)
	}
}

func main() {
	//fmt.Println("vim-go")

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read failed:", err)
				break
			}

			input := string(message)
			cmd := getCmd(input)
			msg := getMessage(input)

			if cmd == "add" {
				todoList = append(todoList, msg)
			} else if cmd == "done" {
				updateTodoList(msg)
			}

			//output := "Current Todos: \n"
			output := ""

			if len(todoList) > 0 {
				output += "<b><span uk-icon='icon: check'></span>Current Todos:</b>"
			}

			output += "<ol>"

			for _, todo := range todoList {
				//output += "\n - " + todo + "\n"
				output += "<li><input type='checkbox' value='" + todo + "' onclick='handleClick(this);'>" + todo + "</li>"
			}

			//output += "\n--------------------------------------"
			output += "</ol>"

			if len(completeTodoList) > 0 {
				output += "<b>Complete Todos:</b>"
			}

			output += "<ul>"

			for _, todo := range completeTodoList {
				//output += "\n - " + todo + "\n"
				output += "<li><s>" + todo + "</s></li>"
			}

			output += "</ul>"

			message = []byte(output)
			err = conn.WriteMessage(mt, message)

			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Setting up the server!")
		http.ServeFile(w, r, "websockets.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.ListenAndServe(":"+port, nil)
}
