package main

import (
	"time"
	"net/http"
	// "io/ioutil"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"math/rand"
)

func curlLoop() {
	count := 0
	for {
		count ++
		url := "http://test.truckmanager.g7s.chinawayltd.com:84/app.php?method=truck.truck.truckList&token=fc1f7e6c9c6e5e9d959b1c5d752f6489"
		//url := "http://test.truckmanager.com/app.php?method=truck.truck.truckList&token=fc1f7e6c9c6e5e9d959b1c5d752f6489"
		client := &http.Client{}
		content := make(map[string]interface{})
		data := make(map[string]interface{})
		data["msgtype"] = rand.Float32()
		data["text"] = content
		dataBytes, _ := json.Marshal(data)

		for i := 0; i < 1; i++ {
			func() {
				request, err := http.NewRequest("POST", url, bytes.NewReader(dataBytes))
				if err != nil {
					panic(err)
				}
				request.Header.Set("Accept-Encoding", "gzip, deflate")

				resp, _ := client.Do(request)
				//fmt.Println(resp.Body)
				fmt.Println(resp.Body.Close())
			}()
		}

		println(count)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
