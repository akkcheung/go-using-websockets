This is Todo app using websockets

Features:
- add task
- complete task

Application design:
- Golang 
- Gorilla websockets package

Installation (local):

1. docker build -t go-using-websockets .
2. docker run --rm -p 5000:5000 go-using-websockets

Installation (deploy to Heroku):

1. heroku login 
2. heroku create
3. heroku stack:set container
4. git push heroku main
5. heroku open

Reference(s):

Using WebSockets in Golang[https://blog.logrocket.com/using-websockets-go/]


