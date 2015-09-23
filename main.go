package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)



func main(){
	
	client := &http.Client{}
	
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.pivotaltracker.com/services/v5/projects/%s/stories/101619634", os.Getenv("PROJECT_ID")), nil)
	
	if err != nil {
		fmt.Println(fmt.Sprintf("There was an error %s", err))
		return
	}
	req.Header.Add("X-TrackerToken", os.Getenv("TOKEN"))
	
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println(fmt.Sprintf("There was an error %s", err))
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
	fmt.Println(string(body))
}