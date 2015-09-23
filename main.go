package main

import (
	"net/http"
	"os"
	"fmt"
)



func main(){
	http.Get("https://www.pivotaltracker.com/services/v5/projects/$PROJECT_ID/stories/101619634")
	fmt.Println(os.Getenv("TOKEN"))
}