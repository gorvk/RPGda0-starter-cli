package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gorvk/RPGda0-starter-cli/types"
)

func main() {
	fmt.Println("Cloning starter project...")

	if err := os.RemoveAll("./RPGda0-starter"); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("git", "clone", "--depth", "1", "https://github.com/gorvk/RPGda0-starter.git")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll("./RPGda0-starter/.git"); err != nil {
		log.Fatal(err)
	}

	var userInput types.UserInput = types.UserInput{}
	userInput.CLIENTURL = "http://localhost:5173"
	userInput.APIHOST = "http://localhost"
	userInput.APIPORT = "9090"
	userInput.APIURL = userInput.APIHOST + ":" + userInput.APIPORT
	userInput.DBPORT = "5432"
	userInput.DBHOST = "starter_db"
	userInput.DBNAME = "starter_db"
	userInput.DBCONNECTIONSTRING = "host=${DB_HOST} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable"

	fmt.Print("Enter auth0 domain: ")
	fmt.Scanln(&userInput.AUTH0DOMAIN)
	fmt.Print("Enter auth0 client id: ")
	fmt.Scanln(&userInput.AUTH0CLIENTID)
	fmt.Print("Enter auth0 client secret: ")
	fmt.Scanln(&userInput.AUTH0CLIENTSECRET)
	fmt.Print("Enter db username: ")
	fmt.Scanln(&userInput.DBUSER)
	fmt.Print("Enter db password: ")
	fmt.Scanln(&userInput.DBPASSWORD)

	apiEnv, err := os.Create("RPGda0-starter/.env")
	if err != nil {
		log.Fatal(err)
	}

	webEnv, err := os.Create("RPGda0-starter/web/.env")
	if err != nil {
		log.Fatal(err)
	}

	commonEnvValues := fmt.Sprintf("AUTH0_DOMAIN=\"%v\"\nAUTH0_CLIENT_ID=\"%v\"", userInput.AUTH0DOMAIN, userInput.AUTH0CLIENTID)

	webEnvValues := commonEnvValues + fmt.Sprintf("\nAPI_URL=\"%v\"", userInput.APIURL)
	_, err = webEnv.WriteString(webEnvValues)
	if err != nil {
		log.Fatal(err)
	}

	apiEnvValues := fmt.Sprintf("CLIENT_URL=\"%v\"\nAPI_HOST=\"%v\"\nAPI_PORT=\"%v\"\nDB_PORT=\"%v\"\nDB_HOST=\"%v\"\nDB_USER=\"%v\"\nDB_PASSWORD=\"%v\"\nDB_NAME=\"%v\"\nDB_CONNECTION_STRING=\"%v\"\nAUTH0_CLIENT_SECRET=\"%v\"",
		userInput.CLIENTURL,
		userInput.APIHOST,
		userInput.APIPORT,
		userInput.DBPORT,
		userInput.DBHOST,
		userInput.DBUSER,
		userInput.DBPASSWORD,
		userInput.DBNAME,
		userInput.DBCONNECTIONSTRING,
		userInput.AUTH0CLIENTSECRET,
	)
	apiEnvValues = apiEnvValues + "\n" + commonEnvValues
	_, err = apiEnv.WriteString(apiEnvValues)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Project created successfully !")
	fmt.Println("Press enter to exit")
	fmt.Scanln()
	fmt.Println("Thank you !")
}
