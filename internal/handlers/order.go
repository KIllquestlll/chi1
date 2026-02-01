package handlers

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listen all order")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get by id")
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update By Id")
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete By Id")
}
