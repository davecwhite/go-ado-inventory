package main

import (
	"encoding/json"
	"fmt"
)

func getUsers(organizationUrl string, authentication string) {

	endpoint := "graph/users"
	fileName := "graph-users.csv"
	response, err := apiCall(endpoint, apiURL(true, organizationUrl, endpoint), authentication)
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

func getProjects(organizationUrl string, authentication string) {

	endpoint := "projects"
	fileName := "projects.csv"
	response, err := apiCall(endpoint, apiURL(false, organizationUrl, endpoint), authentication)
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
