package influx

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"os"
)

//const Host string = "http://ubuntu:8086"
//const token string = "9tZffY820hg2x_GBUXa5PA1rs4s3dwPCpjhP1hZ5PF92deX4mi-pKuWCK_ZD6bhla3-0BdWSbsAiy97VDXLwqA=="
const Bucket string = "sensors"
const Organization string = "MeiSo"

func New() influxdb2.Client {
	influxClient := influxdb2.NewClient(os.Getenv("INFLUX_HOST"), os.Getenv("INFLUX_TOKEN"))
	return influxClient
}
