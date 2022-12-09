package controllers

import (
	"finalTaskWan/models"
	"fmt"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	"strings"
)

func GetData() []models.Car {
	url := "https://storage.googleapis.com/kagglesdsdata/datasets/33080/1584391/CAR%20DETAILS%20FROM%20CAR%20DEKHO.csv?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=gcp-kaggle-com@kaggle-161607.iam.gserviceaccount.com/20221206/auto/storage/goog4_request&X-Goog-Date=20221206T171917Z&X-Goog-Expires=259200&X-Goog-SignedHeaders=host&X-Goog-Signature=66a37dd78b2c4452c02f3725d79eeda009ff71879bb199af06b2a52b9eaa99ca9b42e5944d0bdd6e878a204e0af6cd231fb2efde9671348f1c6180be4aa267ed476e61a77597339f1e2a5a7b3f08a6f756bb0b0f18caf109d4f8562b5c2a8854144e7ad71cbc24f5f80f04279a8c566b908ecfde3209f159a679312877a1c5b478ea9f1b04ee5d26d4f07dc975d1b2861b397100450e4b8813e3bf9f3aa04313f178a274d2df5d60a0e90047e6e22154806fd9fc3c90749a5708722a0d40e5e7aff5249f1651fa892411370e12e0bd94db1d0fc53c1bf07131e4ff7c6afa540cbcd126ac81aed6d43a2595490326a5d2b7fe4a8d9055c0cfd1034f4d06aadaaa"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rsp := string(body)
	listRsp := strings.Split(rsp, "\r\n")
	fmt.Println(listRsp)
	carList := []models.Car{}
	var temp []string
	for i := 1; i <= 10; i++ {
		temp = strings.Split(listRsp[i], ",")
		carList = append(carList, models.Car{
			Name:         temp[0],
			Year:         temp[1],
			SellingPrice: temp[2],
			KmDriven:     temp[3],
			Fuel:         temp[4],
			SellerType:   temp[5],
			Transmission: temp[5],
			Owner:        temp[7],
		})
	}
	fmt.Println(temp)
	return carList
}

func SaveData(cars []models.Car) error {
	// save in database
	for i := 1; i <= 20; i++ {
		err := models.DB.Model(&models.Car{}).Select("name", " year", "selling_price", "km_driven", "fuel", "seller_type", "transmission", "owner").Create(&cars).Error
		return err
	}
	fmt.Println(cars)
	return nil
}
