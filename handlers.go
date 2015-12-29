package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/pg.v3"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	//users := Users{
	//UserInfo{Id: 21341231231, Name: "Bob", Type: "user"},
	//UserInfo{Id: 21341231231, Name: "Samantha", Type: "user"},
	//}
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	users, err := GetUsers(db)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}
func UsersRelationIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVal, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	userRel, err := GetUsersRel(db, idVal)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userRel); err != nil {
		panic(err)
	}
}

func UsersRelationCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idVal, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}
	idRelVal, err := strconv.ParseInt(vars["idrel"], 10, 64)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	userRel := UserRelationInfo{
		Id:    idVal,
		Idrel: idRelVal,
		Type:  "relationship",
	}

	if err := json.Unmarshal(body, &userRel); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	fmt.Println("connect ok")
	_, err = GetUsersRelSingle(db, idVal, idRelVal)
	if err == pg.ErrNoRows {
		_, err = GetUsersRelSingle(db, idRelVal, idVal)
		fmt.Println(err)
		if err == pg.ErrNoRows {
			fmt.Println("ErrNoRows test")
			err := CreateUserRel(db, &userRel)
			if err != nil {
				panic(err)
			}
		} else if err == nil {
			err := UpdateUserRel(db, &userRel)
			if err != nil {
				panic(err)
			}
			userRel.State = "matched"
		} else {
			panic(err)
		}
	} else if err == pg.ErrMultiRows {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userRel); err != nil {
		panic(err)
	}

}

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	var user UserInfo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	user = UserInfo{
		Type: "user",
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()
	userOld, err := GetUserInfoByName(db, user.Name)
	if err != pg.ErrNoRows {
		user.Id = userOld.Id
		user.Type = userOld.Type
	} else {
		err = CreateUser(db, &user)
		if err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}
