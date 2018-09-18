package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"time"

	"github.com/mattn/go-shellwords"
)

//Setting : 録画設定json用構造体
type Setting struct {
	StartHour     int    `json:"startHour"`
	RecTime       string `json:"recTime"`
	OutputDir     string `json:"outputDir"`
	ConnectServer string `json:"connectServer"`
}

func runCmdStr(cmdstr string) error {

	log.Println(cmdstr)

	//文字列をコマンドとオプションにスライス
	c, err := shellwords.Parse(cmdstr)
	failOnError(err)

	switch len(c) {
	case 0:
		//空のコマンドの場合
		return nil
	case 1:
		//オプションなしコマンド
		err := exec.Command(c[0]).Run()
		failOnError(err)

	default:
		//コマンド+オプション
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
	//setting.json 読み込み
	bytes, err := ioutil.ReadFile("setting.json")
	failOnError(err)

	//json デコード
	var setting []Setting
	if err := json.Unmarshal(bytes, &setting); err != nil {
		failOnError(err)
	}

	//現在時刻取得
	t := time.Now()
	hour := t.Hour()

	//ファイル名生成
	fileName := t.Format("2006_01_02_15_04_05") + ".flv"

	//スケジュールを読み込んで録画スタート
	for _, schedule := range setting {
		if schedule.StartHour == hour {
			//コマンド
			cmdstr := "rtmpdump --live --rtmp " + schedule.ConnectServer + " --timeout 60 -B " + schedule.RecTime + " -o " + schedule.OutputDir + fileName

			//実行
			runCmdStr(cmdstr)

			break
		}
	}
}
