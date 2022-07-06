package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "strings"
)

func request(url string) {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(body)
}

func input(text string) string {
    fmt.Print(text)

    // var then variable name then variable type
    var first string

    // Taking input from user
    fmt.Scanln(&first)

    return first
}

func parseResponse(text string) {
    trophy := strings.Split(text, `"SkillRating":`)[1]
    trophy = strings.Split(trophy, ",")[0]
    crown := strings.Split(text, `"Crowns":`)[1]
    crown = strings.Split(crown, ",")[0]
    username := strings.Split(text, `"Username":`)[1]
    username = strings.Split(username, ",")[0]

    fmt.Printf("\rUsername : %s | Trophy : %s | Crown : %s", username, trophy, crown)
}

func requestWithHeader(url string, token string) {
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)

    req.Header.Set("Host", "kitkabackend.eastus.cloudapp.azure.com:5010")
    req.Header.Set("use_response_compression", "true")
    req.Header.Set("authorization", token)

    res, _ := client.Do(req)
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
    }

    if res.StatusCode == 200 {
        parseResponse(string(body))
    }
    if res.StatusCode == 401 {
      fmt.Printf("\n[401] Failed Maybe Auth Key Expired Or Wrong?")
    }
	if res.StatusCode == 403 {
		fmt.Printf("\n[403] Your Account Has Ben Banned From Server")
	}
	if res.StatusCode == 501 {
		fmt.Printf("\n[501] Failed")
	}
}

func main() {
	fmt.Printf("Simple Stumble Duplicate Crown & Tropy")
	fmt.Printf("\nkontoll!")
	fmt.Printf("\nmain murni ngebosenin dek dek")
	fmt.Printf("\nHaha stumble memek kntl anjing burik")

	fmt.Printf("\n\n\n╔═╗╔╦╗╦ ╦╔╦╗╔╗ ╦  ╔═╗  ╔╦╗╦ ╦╔═╗╦  ╦╔═╗╔═╗╔╦╗╔═╗\n")
	fmt.Printf("╚═╗ ║ ║ ║║║║╠╩╗║  ║╣    ║║║ ║╠═╝║  ║║  ╠═╣ ║ ║╣ \n")
	fmt.Printf("╚═╝ ╩ ╚═╝╩ ╩╚═╝╩═╝╚═╝  ═╩╝╚═╝╩  ╩═╝╩╚═╝╩ ╩ ╩ ╚═╝")

	round := input("\n\n 0. Eleminated \n1. Round 1\n2. Round 2\n3. Round(Winner) 3\n Input Your Round: ")
    token := input("Enter your authorization: ")
    thread := input("Threads :")
    thrd, _ := strconv.Atoi(thread)
    url := "http://kitkabackend.eastus.cloudapp.azure.com:5010/round/finishv2/" + round

    for i := 0; i < thrd; i++ {
        go func() {
            for {
                requestWithHeader(url, token)
            }
        }()
    }
    for {
        requestWithHeader(url, token)
    }
}
