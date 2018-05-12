package main

import (
	"log"
	. "github.com/mlabouardy/dialogflow-go-client"
	apiai "github.com/mlabouardy/dialogflow-go-client/models"
)

func GetResponse(input string) apiai.Result {
	err, client := NewDialogFlowClient(apiai.Options{
		AccessToken: "d90d61abd22b48a7864210cf1c293366",
	})
	if err != nil {
		log.Fatal(err)
	}

	query := apiai.Query{
		Query: input,
	}
	resp, err := client.QueryFindRequest(query)
	if err != nil {
		log.Fatal(err)

	}
	return resp.Result
}
