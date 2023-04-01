package middlewares

import "fmt"

func ServeMessage() {
	fmt.Println("   ┌───────────────────────────────────────────┐\n   │                                           │\n   │   Serving!                                │\n   │                                           │\n   │   - Golang: http://localhost:8080         │\n   │   - React:  http://localhost:3000         │\n   │                                           │\n   │   Don't forget to launch React app!       │\n   │                                           │\n   └───────────────────────────────────────────┘")
}
