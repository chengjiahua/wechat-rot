package main

import (
	"bufio"
	"crypto/rand"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"
	"wechat-rot/common"

	"github.com/robfig/cron"
)

func Getup() {
	params := url.Values{}
	Url, err := url.Parse(common.RoomUrl)
	if err != nil {
		return
	}
	mun, _ := rand.Int(rand.Reader, big.NewInt(int64(len(common.MapGetup))))
	params.Set("msg", common.MapGetup[int(mun.Int64())])
	Url.RawQuery = params.Encode()
	resp, _ := http.Get(Url.String())
	resp.Body.Close()
}

func CaiHongPi() {
	params := url.Values{}
	Url, err := url.Parse(common.RoomUrl)
	if err != nil {
		return
	}
	mun, _ := rand.Int(rand.Reader, big.NewInt(int64(len(common.MapLoveWords))))
	params.Set("msg", common.MapLoveWords[int(mun.Int64())])
	Url.RawQuery = params.Encode()
	resp, _ := http.Get(Url.String())
	resp.Body.Close()
}

func TianGouDiary() {
	params := url.Values{}
	Url, err := url.Parse(common.RoomUrl)
	if err != nil {
		return
	}
	mun, _ := rand.Int(rand.Reader, big.NewInt(int64(len(common.MapTianGouDiary))))
	params.Set("msg", common.MapTianGouDiary[int(mun.Int64())])
	Url.RawQuery = params.Encode()
	resp, _ := http.Get(Url.String())
	resp.Body.Close()
}

func main() {

	// 加载MAP资源
	InitMap()

	// 定制corn任务
	c := cron.New()
	cronGetup := "0 30 8 * * ? "                                          //每天八点半
	cronCaiHongPi := "0 60 1,3,5,8,10,12,19,17,15,22 * * ? "              //每隔一小时
	cronTianGouDiary := "0 60 0,7,14,21,23,16,9,2,18,11,4,6,13,20 * * ? " //每隔一小时

	c.AddFunc(cronGetup, Getup)
	c.AddFunc(cronCaiHongPi, CaiHongPi)
	c.AddFunc(cronTianGouDiary, TianGouDiary)

	// test := "0/5 * * * * ? "
	// c.AddFunc(test, Getup)

	c.Start()

	select {}
}

func InitMap() {
	InitTiangouDiaryMap()
	InitCaiHongPiMap()
	InitGetupMap()
}

func InitTiangouDiaryMap() {
	file, _ := os.Open("./doc/舔狗日记.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		pureLine := strings.TrimSpace(line)
		common.MapTianGouDiary[i] = pureLine
		i++
		if err != nil {
			break
		}
	}
}

func InitCaiHongPiMap() {
	file, _ := os.Open("./doc/彩虹屁.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		pureLine := strings.TrimSpace(line)
		common.MapLoveWords[i] = pureLine
		i++
		if err != nil {
			break
		}
	}
}

func InitGetupMap() {
	file, _ := os.Open("./doc/叫起床.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		pureLine := strings.TrimSpace(line)
		common.MapGetup[i] = pureLine
		i++
		if err != nil {
			break
		}
	}
}
