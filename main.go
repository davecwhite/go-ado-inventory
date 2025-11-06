package main

import (
	"os"
)

const OUTPUT_DIRECTORY = "output"

func main() {

	tenant := os.Getenv("ADO_TENANT_ID")
	clientId := os.Getenv("ADO_CLIENT_ID")
	clientSecret := os.Getenv("ADO_CLIENT_SECRET")
	organization := os.Getenv("GO_ADO_ORGANIZATION")

	err := os.Mkdir(OUTPUT_DIRECTORY, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	adoToken := getADOToken(tenant, clientId, clientSecret)
	getUsers(organization, adoToken)
	getProjects(organization, adoToken)

}
