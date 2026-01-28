package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"ok","time":"%s"}`, time.Now().Format(time.RFC3339))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <title>Hello World</title>
    <style>
        body {
            font-family: sans-serif;
            text-align: center;
            padding: 50px;
            background: #f5f5f5;
        }
        h1 {
            color: #333;
        }
        img {
            max-width: 400px;
            border-radius: 10px;
            box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        }
        p {
            color: #666;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <h1>Hello, World!</h1>
    <img src="https://media.giphy.com/media/3o6Zt6ML6BklcajjsA/giphy.gif" alt="German Shorthaired Pointer waving hello">
    <p>A friendly German Shorthaired Pointer says hi!</p>
    <p style="font-size: 12px; color: #999;">Server time: %s</p>
</body>
</html>`, time.Now().Format("Monday, January 2, 2006 at 3:04 PM"))
	})

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
