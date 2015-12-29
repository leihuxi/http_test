package main

type UserInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Users []UserInfo

type UserRelationInfo struct {
	Id    int64  `json:"id"`
	Idrel int64  `json:"idrel"`
	State string `json:"state"`
	Type  string `json:"type"`
}

type UsersRelation []UserRelationInfo
