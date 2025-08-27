package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pricesTab() *fyne.Container {
	chart := app.getChart()
	chartContainer := container.NewVBox(chart)
	return chartContainer
}

func (app *Config) getChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))
	var img *canvas.Image

	err := app.downloadFile(apiURL, "gold.png")
	if err != nil {
		img = canvas.NewImageFromFile("./cmd/unreachable.png")
	} else {
		img = canvas.NewImageFromFile("./cmd/gold.png")
	}

	img.SetMinSize(fyne.Size{
		Width:  770,
		Height: 410,
	})

	img.FillMode = canvas.ImageFillOriginal
	return img
}

func (app *Config) downloadFile(url, fileName string) error {
	response, err := app.HTTPClient.Get(url)
	if err != nil {
		log.Println("Fail to get image from url")
		return err
	}

	if response.StatusCode == int(STATUS_OK) {
		b, err := io.ReadAll(response.Body)

		if err != nil {
			log.Println("Fail to read bytes from response body ", err.Error())
			return err
		}

		defer response.Body.Close()

		img, _, err := image.Decode(bytes.NewReader(b))

		if err != nil {
			log.Println("Fail to decode image bytes", err.Error())
			return err
		}

		out, err := os.Create(fmt.Sprintf("./cmd/%s", fileName))

		if err != nil {
			log.Println("Fail when creating image file", err.Error())
			return err
		}

		err = png.Encode(out, img)
		if err != nil {
			log.Println("Fail to encode png", err.Error())
			return err
		}
		return nil
	}
	return errors.New(fmt.Sprintf("Received wrong response code when dowloading image:%s", response.Status))

}
