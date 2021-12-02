package util

import (
	"encoding/base64"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"time"
)

func Base64toFile(base64code string) string {
	decodeData, err := base64.StdEncoding.DecodeString(base64code)
	if err != nil {
		panic(err)
	}
	imgName := "img-" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".png"
	f, err := os.OpenFile(imgName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := time.Now().Second()
	f.Write(decodeData)
	e := time.Now().Second()
	log.Info().Msgf("Took %d seconds generating file.", e-s)
	return imgName
}

func SimpleFormPost(client *resty.Client, body []byte, url string) (*resty.Response, error) {
	return client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded").SetBody(body).Post(url)
}

func LogError(err error) {
	log.Err(err)
}
