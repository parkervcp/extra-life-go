package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	var err error

	var apiData apiConf

	fileBytes, err := loadFile("config.json")
	if err != nil {
		log.Printf("there was an issue reading the config file \n%s", err)
	}

	err = loadJSON(fileBytes, &apiData)
	if err != nil {
		log.Printf("there was an issue loading the data from the config file \n%s", err)
	}

	var teamDono donations

	qBytes, err := queryAPI(apiData.APIURL, "teams/"+apiData.TeamID+"/donations", "GET", []byte(""))
	if err != nil {
		log.Printf("There was an error getting the api information")
	}

	err = loadJSON(qBytes, &teamDono)
	if err != nil {
		log.Printf("There was an error converting the data \n%s", err)
	}

	log.Printf("%s", strconv.Itoa(int(teamDono[0].Amount)))

}

func queryAPI(url string, endpoint string, request string, data []byte) ([]byte, error) {
	//var for response body byte
	var bodyBytes []byte
	//http get json request
	client := &http.Client{}
	req, _ := http.NewRequest(request, url+endpoint, nil)

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

	log.Printf("There are %s records", resp.Header.Get("Num-Records"))

	//return byte structure
	return bodyBytes, nil
}

func printJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func loadFile(file string) (fileBytes []byte, err error) {
	if !strings.HasSuffix(file, ".json") {
		return nil, fmt.Errorf("the file specified was not a json file")
	}

	// Open our jsonFile
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	fileBytes, _ = ioutil.ReadAll(jsonFile)

	return
}

func loadJSON(inBytes []byte, v interface{}) (err error) {
	err = json.Unmarshal(inBytes, &v)
	if err != nil {
		return err
	}

	return nil
}
