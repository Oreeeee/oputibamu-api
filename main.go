package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"oputibamu"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	timetableUrl := os.Getenv("TIMETABLE_URL")
	elektronikApi := os.Getenv("ELEKTRONIK_API")
	var elektronikMode bool
	if elektronikApi == "0" {
		elektronikMode = false
	} else {
		elektronikMode = true
	}
	vo := oputibamu.VOScraper{timetableUrl, elektronikMode, elektronikApi}
	fmt.Println(vo.GetClasses())
}
