package defs

type UserCredential struct {
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}

type VideoInfo struct {
	Id string
	AuthodId string
	Name string
	DisplayCtime string
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}