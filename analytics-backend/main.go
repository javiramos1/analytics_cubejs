package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

// Event Definition
type Event struct {
	EventID           string    `json:"event_id"`
	Type              string    `json:"event_type"`
	Path              string    `json:"page_path"`
	UserIP            string    `json:"user_ip"`
	UserAgent         string    `json:"user_agent"`
	UserOS            string    `json:"user_os"`
	IsMobile          bool      `json:"is_mobile"`
	UserCity          string    `json:"user_city"`
	UserCountry       string    `json:"user_country_code"`
	UserCountryName   string    `json:"user_country_name"`
	UserCountryRegion string    `json:"user_country_region"`
	UserContinent     string    `json:"user_continent"`
	UserTimezone      string    `json:"user_timezone"`
	UserLanguages     string    `json:"user_languages"`
	UserInEU          bool      `json:"user_in_eu"`
	ClickText         string    `json:"click_text"`
	ClickID           string    `json:"click_id"`
	ClickClass        string    `json:"click_class"`
	EventTime         time.Time `json:"event_time"`
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("content-type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func saveEvent(w http.ResponseWriter, r *http.Request) {

	var event Event
	e := json.NewDecoder(r.Body).Decode(&event)
	checkError(e)
	log.Printf("Event Request: %v", event)

	setupResponse(&w, r)

	insertDynStmt := `insert into analytics.events
	("event_id", "event_type","page_path", "user_ip", "user_agent", "user_os", 
	"user_city", "user_country_code", "user_country_name", "user_country_region", "user_continent", "user_timezone", "user_languages", "user_in_eu",
	"is_mobile", "click_text", "click_id","click_class", "event_time") 
	values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`
	log.Printf("Query: %v", insertDynStmt)

	_, e = db.Exec(insertDynStmt,
		event.EventID, event.Type, event.Path, event.UserIP, event.UserAgent, event.UserOS,
		event.UserCity, event.UserCountry, event.UserCountryName, event.UserCountryRegion, event.UserContinent, event.UserTimezone, event.UserLanguages, event.UserInEU,
		event.IsMobile, event.ClickText, event.ClickID, event.ClickClass, event.EventTime)

	if e != nil {
		log.Fatalf("Error saving event: %v", e)
	} else {
		json.NewEncoder(w).Encode(event)
	}

}

func main() {
	// concurrency level, use 2x number of vCPU
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	log.Printf("GOMAXPROCS %v", runtime.GOMAXPROCS(-1))

	e := godotenv.Load("etc/.env")
	if e != nil {
		e = godotenv.Load()
	}
	checkError(e)

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port,
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	log.Printf("connection: %s", psqlconn)

	db, e = sql.Open("postgres", psqlconn)
	checkError(e)

	defer db.Close()

	http.HandleFunc("/event", saveEvent)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
