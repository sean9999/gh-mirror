package main

type Repo struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	Url    string `json:"url"`
	SshUrl string `json:"sshUrl"`
}
