package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	domain = "https://www.extra-life.org/api/"
)

func main() {
	var err error

	var teamData team

	qBytes, err := queryAPI("teams/"+"", "GET", []byte(""))
	if err != nil {
		log.Printf("There was an error getting the api information")
	}

	err = loadJSON(qBytes, &teamData)
	if err != nil {
		log.Printf("There was an error converting the data \n%s", err)
	}

	log.Printf("%s", strconv.Itoa(int(teamData.FundraisingGoal)))

	var teamParticipants teamUsers

	qBytes, err = queryAPI("teams/"+""+"/participants", "GET", []byte(""))
	if err != nil {
		log.Printf("There was an error getting the api information")
	}

	err = loadJSON(qBytes, &teamParticipants)
	if err != nil {
		log.Printf("There was an error converting the data \n%s", err)
	}

	log.Printf("%s", teamParticipants[1].DisplayName)

	var participant user

	qBytes, err = queryAPI("/participants/357888", "GET", []byte(""))
	if err != nil {
		log.Printf("There was an error getting the api information")
	}

	err = loadJSON(qBytes, &participant)
	if err != nil {
		log.Printf("There was an error converting the data \n%s", err)
	}

	log.Printf("%s", participant.DisplayName)
}

func queryAPI(endpoint string, request string, data []byte) ([]byte, error) {
	//var for response body byte
	var bodyBytes []byte
	//http get json request
	client := &http.Client{}
	req, _ := http.NewRequest(request, domain+endpoint, nil)

	//Sets request header for the http request
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	//send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 && resp.StatusCode != 204 {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
		return bodyBytes, errors.New(string(bodyBytes))
	}

	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}

	//set bodyBytes to the response body
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	//Close response thread
	defer resp.Body.Close()

	//return byte structure
	return bodyBytes, nil
}

func printJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func loadJSON(inBytes []byte, v interface{}) (err error) {
	err = json.Unmarshal(inBytes, &v)
	if err != nil {
		return err
	}

	return nil
}
