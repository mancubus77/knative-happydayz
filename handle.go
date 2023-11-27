package function

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Lang    string `json:"lang"`
	Message string `json:"message"`
}

type Simple struct {
	VarA int    `json:"vara"`
	VarB int    `json:"varb"`
	Oper string `json:"oper"`
}

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {

	fmt.Println("Received request")
	fmt.Println(prettyPrint(req)) // echo to local output
	// fmt.Fprint(res, prettyPrint(req)) // echo to caller
	payload := Response{
		Message: "I 💕 GoLang Serverless",
		Lang:    "GoLang Serverless",
	}
	jsonResponse, _ := json.Marshal(payload)

	if req.Method == "Post" {
		res.Header().Set("Content-Type", "application/json")
		jsonResponse, _ := json.Marshal(parsPost(req))
		res.Write(jsonResponse)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
	// fmt.Fprint(res, jsonResponse)
}
