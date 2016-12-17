package main

type Place struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Places []Place