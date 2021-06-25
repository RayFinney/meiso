package main

import (
	"context"
	"encoding/json"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"io/ioutil"
	"log"
	"meiso/models"
	"net/http"
	"os"
	"time"
)

var influxClient influxdb2.Client

const influxHost string = "http://ubuntu:8086"
const influxToken string = "9tZffY820hg2x_GBUXa5PA1rs4s3dwPCpjhP1hZ5PF92deX4mi-pKuWCK_ZD6bhla3-0BdWSbsAiy97VDXLwqA=="
const influxBucket string = "sensors"
const influxOrg string = "MeiSo"

func handleStats(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var stats models.SensorStats
	err := json.Unmarshal(reqBody, &stats)
	if err != nil {
		log.Println(err)
	}
	log.Println(stats)

	writeAPI := influxClient.WriteAPIBlocking(influxOrg, influxBucket)
	// Create point using fluent style
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "lux").
		AddField("current", stats.Lux).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("current", stats.Temp).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "humidity").
		AddField("current", stats.Humidity).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "feltTemperature").
		AddField("current", stats.FTemp).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)

	w.WriteHeader(http.StatusCreated)
}

func handleStartup(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := fmt.Sprintf("%s - Startup", time.Now().String())
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getLogs(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("logs.log")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(dat)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	// Create a new client using an InfluxDB server base URL and an authentication token
	influxClient = influxdb2.NewClient(influxHost, influxToken)
	defer influxClient.Close()
	// Use blocking write client for writes to desired bucket

	http.HandleFunc("/stats", handleStats)
	http.HandleFunc("/startup", handleStartup)
	http.HandleFunc("/logs", getLogs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
