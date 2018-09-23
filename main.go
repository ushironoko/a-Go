package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"time"

	"github.com/mattn/go-shellwords"
)

//Setting setting.json struct
type Setting struct {
	StartHour     int    `json:"startHour"`
	RecTime       string `json:"recTime"`
	OutputDir     string `json:"outputDir"`
	ConnectServer string `json:"connectServer"`
}

func runCmdStr(cmdstr string) error {

	log.Println(cmdstr)

	c, err := shellwords.Parse(cmdstr)
	failOnError(err)

	switch len(c) {
	case 0:
		return nil

	case 1:
		err := exec.Command(c[0]).Run()
		failOnError(err)

	default:
		err := exec.Command(c[0], c[1:]...).Run()
		failOnError(err)

	}

	return nil
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	bytes, err := ioutil.ReadFile("setting.json")
	failOnError(err)

	var setting []Setting
	if err := json.Unmarshal(bytes, &setting); err != nil {
		failOnError(err)
	}

	t := time.Now()
	hour := t.Hour()

	fileName := t.Format("2006_01_02_15_04_05") + ".flv"

	for _, schedule := range setting {
		if schedule.StartHour == hour {
			cmdstr := "rtmpdump --live --rtmp " + schedule.ConnectServer + " --timeout 60 -B " + schedule.RecTime + " -o " + schedule.OutputDir + fileName

			runCmdStr(cmdstr)

			break
		}
	}
}
