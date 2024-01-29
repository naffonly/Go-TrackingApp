package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"trackingApp/features/geoapify/model"
)

type Response struct {
	//Features   interface{} `json:"features"`
	Properties interface{} `json:"properties"`
	Type       string      `json:"type"`
}

type Geoapinterface struct {
	Db  *gorm.DB
	Key string
}

func NewGeoapify(db *gorm.DB, key string) GeoapifyInterface {
	return &Geoapinterface{Db: db, Key: key}
}

type GeoapifyInterface interface {
	Insert() gin.HandlerFunc
}

func (handler *Geoapinterface) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := handler.Key

		var payload model.GeoDTO
		errs := c.Bind(&payload)
		if errs != nil {
			c.JSON(http.StatusBadRequest, errs.Error())
			c.Abort()
			return
		}

		pickUp := fmt.Sprintf("%s,%s", payload.PickUp.Lat, payload.PickUp.Lon)
		dropOff := fmt.Sprintf("%s,%s", payload.DropOff.Lat, payload.DropOff.Lon)

		url := fmt.Sprintf("https://api.geoapify.com/v1/routing?waypoints=%s|%s&mode=%s&apiKey=%s", pickUp, dropOff, payload.Type, key)
		method := "GET"
		fmt.Println(url)
		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		var data Response
		err = json.Unmarshal(body, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, data)

	}
}
