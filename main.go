package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const OUTPUT_DIRECTORY = "output"

type UserResults struct {
	Count int     `json:"count"`
	Value []users `json:"value"`
}

type users struct {
	Descriptor    string
	DisplayName   string
	PrincipalName string
	MailAddress   string
	SubjectKind   string
	Domain        string
}

type ProjectResults struct {
	Count int       `json:"count"`
	Value []project `json:"value"`
}

type project struct {
	Id          string
	Name        string
	Description string
	State       string
	Visibility  string
	LastUpdate  string
	URL         string
}

func main() {

	organizationUrl := os.Getenv("GO_ADO_ORGANIZATION")
	authorization := os.Getenv("GO_ADO_AUTHORIZATION")

	err := os.Mkdir(OUTPUT_DIRECTORY, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	getUsers(organizationUrl, authorization)
	getProjects(organizationUrl, authorization)

}

func apiURL(isGraph bool, organizationUrl string, endpoint string) string {

	base := "https://dev.azure.com/"
	if isGraph {
		base = "https://vssps.dev.azure.com/"
	}

	return base + organizationUrl + "/_apis/" + endpoint + "?api-version=7.2-preview"

}
func apiCalll(url string, authorization string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Basic "+authorization)

	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	responseBody := ""

	for i := 0; scanner.Scan() && i < 5; i++ {
		responseBody += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return responseBody, nil
}

func writeToFile(fileName string, data string) {
	file, err := os.Create(OUTPUT_DIRECTORY + "/" + fileName)
	if err != nil {
		println(err)
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		println(err)
	}
}

func getUsers(organizationUrl string, authorization string) {

	endpoint := "graph/users"
	fileName := "graph-users.csv"
	response, err := apiCalll(apiURL(true, organizationUrl, endpoint), authorization)
	if err != nil {
		println(err)
	}

	userList := UserResults{}
	err = json.Unmarshal([]byte(response), &userList)
	if err != nil {
		println(err)
	}

	output := "User,Email,Subject Kind,Principal Name,Domain\n"

	for _, user := range userList.Value {
		output += fmt.Sprintf("%s,%s,%s,%s,%s\n", user.DisplayName, user.MailAddress, user.SubjectKind, user.PrincipalName, user.Domain)
	}

	writeToFile(fileName, output)

}

func getProjects(organizationUrl string, authorization string) {

	endpoint := "projects"
	fileName := "projects.csv"
	response, err := apiCalll(apiURL(false, organizationUrl, endpoint), authorization)
	if err != nil {
		println(err)
	}

	projectList := ProjectResults{}
	err = json.Unmarshal([]byte(response), &projectList)
	if err != nil {
		println(err)
	}

	output := "Id,Name,Description,State,Visibility,LastUpdate,URL\n"

	for _, project := range projectList.Value {
		output += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n", project.Id, project.Name, project.Description, project.State, project.Visibility, project.LastUpdate, project.URL)
	}

	writeToFile(fileName, output)

}
