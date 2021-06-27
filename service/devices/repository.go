package devices

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const logFile = "logs.log"

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r *Repository) StoreStartupLogs(deviceId string) error {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	newLine := fmt.Sprintf("%s (%s) - Startup", deviceId, time.Now().String())
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *Repository) GetLogs() (string, error) {
	dat, err := ioutil.ReadFile(logFile)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
