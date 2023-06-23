package controller

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

type UserResponse struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}

func CheckUserAuthToken(tokenString string) bool {
	userResponse := callAuthApi(tokenString)
	if userResponse.Id == 0 || userResponse.Email == "" || userResponse.Username == "" {
		return false
	}
	
	return true
}

func callAuthApi(tokenString string) UserResponse {
	host := os.Getenv("AUTH_API_HOST")
	url := fmt.Sprintf("%s/user_check", host)
	request, _ := http.NewRequest("GET", url, nil)
    request.Header.Set("Authorization", tokenString)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var userResponse UserResponse
	json.Unmarshal(body, &userResponse)
	
	return userResponse
}
