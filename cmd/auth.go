// entry point and global flags
package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Credentials struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	AuthToken string `json:"authToken"`
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with the MDB Platform",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Prompt for username/password (use a library like 'survey')
		mockAuthCredentials := Credentials{
			Username: "email@gmail.com",
			Password: "-----BEGIN PGP MESSAGE-----\nVersion: BCPG v1.67\n\nhQIMA1RBGJD4NAj2AQ/9HLVBumHkL7pLFjKVki6Tdde/X2OkK5hzdi6ZJdGTSSIm\n6XGAL3HHQHwiMv6r7CwJXC3ktx+9WEwjbx9Ltwl5i35OAq5RvrBxZ1yopa50OcH3\nIy7hA7sXuQnrczzMwWgtLBsj32Y3OWQQLMFHkNEWgZDUGdC+xxauGeofaHHvEsY8\nQG+51Ux69/+94oczMhGDKKh0MbbxF9Oo2gYNxtjhtUFKh4YyF1CNWBUV9aU6qkSA\nI3DwCMiuVdrw/9tXRBNSO5Aqw3y6gjhvOCJ94R4M5pVgeLMwHT+f/WTB6APZqxdn\nUcesh9AkHVJZVUAxNFKPQ5BbFP/z6TqlzsBQo2Vus0vsb5pMO0noJ6JlO2IKsYOX\nvj/j4dfDVEMEEGkReB5f6P/Ybzs9CoaZ0ZvUa7+1GmAbn8qolgOWl0bo7ZGSt0sL\nQI1el2ICbfPwQA/V+3/rAJryUaMvrIrJgH86vxpWtxz3T6LV6i8OmTOQwLTPflL2\nfVGM+EmqMMdoy77AykWqRnZUckBhu6F9S/9WTw/dJcMrYFeepBZHRlrrASb66Usw\nA/UIUjrYdMEU//dq1l/lbfpQ4Tdgw5ctIQUP8PdSh9uNiUayVKJGk/+DhMslMzjG\npS2tsTer9A5jzS8Q6YpBZmTMXVvI4++H4BHb7mlx3A5gKp5UIQhCGfvCDmAN9bjS\nRwGkoIGrRD5CuEPapzRkvLa23E05Hm4g+qYOlnUdY1uA8cWZ8brn6Slkh1hsp3EQ\n5l8G29EQi45L9OKR/MQdTBQEgaJhbASs\n=CXAb\n-----END PGP MESSAGE-----\n",
		}

		mdbURL := "http://localhost:8082"
		url := mdbURL + "/mdb/v1/auth/login"
		log.Println("url: " + url)
		jsonData, err := json.Marshal(mockAuthCredentials)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Logging in..", mockAuthCredentials.Username)

		response, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
		defer response.Body.Close()

		bodyBytes, err := io.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		log.Println("Response Body: " + string(bodyBytes))
		log.Println("Response Status: " + response.Status)
		scanner := bufio.NewScanner(response.Body)
		// scanner.Scan()
		var responseData Credentials
		for i := 0; scanner.Scan() && i < 50; i++ {
			fmt.Println(scanner.Text())
		}
		json.Unmarshal(bodyBytes, &responseData)
		authToken := responseData.AuthToken
		fmt.Println("responseData", authToken)
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		viper.Set("auth.token", authToken)
		viper.WriteConfig()

		fmt.Println("Successfully authenticated!")
		// fmt.Println(viper.ReadInConfig())
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
