package main

import (
	"fmt"
	"gopkg.in/pg.v3"
)

func CreateUser(db *pg.DB, userInfo *UserInfo) error {
	_, err := db.QueryOne(userInfo, `INSERT INTO users (name,type) VALUES (?name, ?type) RETURNING id`, userInfo)
	return err
}

func CreateUserRel(db *pg.DB, userInfo *UserRelationInfo) error {
	fmt.Println(userInfo.Id, userInfo.Idrel, userInfo.State, userInfo.Type)
	_, err := db.ExecOne(`INSERT INTO user_rel (id, idrel, state, type) VALUES (?id, ?idrel, ?state, ?type)`, userInfo)
	return err
}

func UpdateUserRel(db *pg.DB, userInfo *UserRelationInfo) error {
	_, err := db.ExecOne(`update user_rel set state = 'matched' where idrel = ?id and id = ?idrel`, userInfo)
	return err
}

func GetUserInfo(db *pg.DB, id int64) (*UserInfo, error) {
	var userInfo UserInfo
	_, err := db.QueryOne(&userInfo, `
	SELECT * FROM users WHERE id = ?
	`, id)
	return &userInfo, err
}

func GetUserInfoByName(db *pg.DB, name string) (*UserInfo, error) {
	var userInfo UserInfo
	_, err := db.QueryOne(&userInfo, `
	SELECT * FROM users WHERE name = ?
	`, name)
	return &userInfo, err
}

func GetUsers(db *pg.DB) (Users, error) {
	var users Users
	_, err := db.Query(&users, `
	SELECT * FROM users
	`)
	return users, err
}

func GetUsersRel(db *pg.DB, id int64) (UsersRelation, error) {
	var usersRel UsersRelation
	_, err := db.Query(&usersRel, `
	select * from user_rel where id = ?
	`, id)
	return usersRel, err
}

func GetUsersRelSingle(db *pg.DB, id int64, idrel int64) (UserRelationInfo, error) {
	var usersRel UserRelationInfo
	fmt.Println(id, idrel)
	_, err := db.QueryOne(&usersRel, `
	select * from user_rel where id = ? and idrel = ?
	`, id, idrel)
	return usersRel, err
}

var db *pg.DB

func init() {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()
}
