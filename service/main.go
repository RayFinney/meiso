package main

import (
	"fmt"
	"meiso/db"
	"meiso/devices"
	"meiso/influx"
	"meiso/plants"
	"meiso/router"
	"meiso/sensors"
	"os"
)

const apiV1 = "/api/v1"

func main() {
	r := router.New()
	v1 := r.Group(apiV1)
	ic := influx.New()
	d := db.New()

	_, _, _ = sensors.Setup(v1, ic)
	_, _, _ = devices.Setup(v1)
	_, _, _ = plants.Setup(v1, d)

	r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
