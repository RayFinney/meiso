package main

import (
	"meiso/devices"
	"meiso/influx"
	"meiso/router"
	"meiso/sensors"
)

const apiV1 = "/api/v1"

func main() {
	r := router.New()
	v1 := r.Group(apiV1)
	ic := influx.New()

	_, _, _ = sensors.Setup(v1, ic)
	_, _, _ = devices.Setup(v1)

	r.Logger.Fatal(r.Start(":8080"))
}
