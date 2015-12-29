package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"UsersIndex",
		"GET",
		"/users",
		UsersIndex,
	},
	Route{
		"UsersCreate",
		"POST",
		"/users",
		UsersCreate,
	},
	Route{
		"UsersRelationIndex",
		"GET",
		"/users/{id}/relationships",
		UsersRelationIndex,
	},
	Route{
		"UsersRelationCreate",
		"PUT",
		"/users/{id}/relationships/{idrel}",
		UsersRelationCreate,
	},
}
