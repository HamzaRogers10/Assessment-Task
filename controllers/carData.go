package controllers

import (
	"finalTaskWan/models"
	"fmt"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	"strings"
)

// GetData Create a function that will call URL to get the data periodically
func GetData() []models.Car {
	//function to get the data generated from the URL
	url := "https://storage.googleapis.com/kagglesdsdata/datasets/33080/1584391/CAR DETAILS FROM CAR DEKHO.csv?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=gcp-kaggle-com@kaggle-161607.iam.gserviceaccount.com/20221220/auto/storage/goog4_request&X-Goog-Date=20221220T151224Z&X-Goog-Expires=259200&X-Goog-SignedHeaders=host&X-Goog-Signature=008e49bfa7f112ba970f80d63b0a5d119b5e7bc30a46f4928aeaf7fa132ecafa33a677e470f3024c0df018b89ae0405a2e922763b38f23924aee5a79687c67264151a7890cd0408af3dfeca1cb69e65e050bafbe2bd54c230052fb04e517274cc2d9a1ad11c4390675d29dac4f1c10f79c3ac34d76864211c7977517881e96ea99d025825add51587493afd4b662ff4b90fe6e67487e297551c5ccfc62b13b4f0d6d4647a6be04e41ffdf2b96071628ce3b6610c4983df301419639a6d7718f38f17bc51ef697131aca8d205372cc81e79e4f147382a0a4582fce8a9606e2102ea59582089b6e1572d21727f6ba45e8601a53afa0344c0c9f6f05e2e6e991dbe"
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
	response := string(body)
	//split the string into model
	listResponse := strings.Split(response, "\r\n")
	fmt.Println(listResponse)
	var carList []models.Car
	var temp []string
	for i := 1; i <= 30; i++ {
		temp = strings.Split(listResponse[i], ",")
		carList = append(carList, models.Car{
			Name:         temp[0],
			Year:         temp[1],
			SellingPrice: temp[2],
			KmDriven:     temp[3],
			Fuel:         temp[4],
			SellerType:   temp[5],
			Transmission: temp[6],
			Owner:        temp[7],
		})
	}
	fmt.Println(temp)
	return carList
}
