package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"labora-api/services"
	"labora-api/models"
	"log"
	"strconv"
	"net/http"
)



// Función para obtener todos los elementos
func GetItems(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	items, err := services.GetItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al obtener los items"))
		return
	}

	pageParam := r.URL.Query().Get("page")
	itemsPerPageParam := r.URL.Query().Get("itemsPerPage")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}

	itemsPerPage, err := strconv.Atoi(itemsPerPageParam)
	if err != nil {
		itemsPerPage = 10
	}

	start := (page - 1) * itemsPerPage
	end := start + itemsPerPage
	if end > len(items) {
		end = len(items)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// Función para obtener un elemento específico
func GetItem(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idItem := params["id"]
	item, err := services.GetItem(idItem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al obtener el item"))
		return
	}

	if item == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Objeto con id %s no encontrado", idItem)))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func SearchItem(w http.ResponseWriter, r *http.Request)  {

}

// Función para crear un nuevo elemento
func CreateItem(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var newItem models.Item

	err = json.Unmarshal(requestBody, &newItem)
	if err != nil {
		fmt.Println(err)
		return
	}

	createdItemID, err := services.CreateItem(newItem)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error al crear el item")))
			return
	}

	responseBody, err := json.Marshal(fmt.Sprintf("Id inserido: %d", createdItemID))
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}

// Función para actualizar un elemento existente
func UpdateItem(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var itemUpdated models.Item

	err = json.Unmarshal(requestBody, &itemUpdated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	idItem, err := strconv.Atoi(params["id"])
	if err != nil {
		err = errors.New("Error al convertir id a entero")
		w.Write([]byte(fmt.Sprint(err)))
	}
	itemUpdated.ID = idItem

	rowsAffected, err := services.UpdateItem(itemUpdated)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("No fue posible actualizar el item solicitado")))
	} else if rowsAffected == 0 {
		err = errors.New("El Objeto no existe")
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("(%d) Objeto(s) con id: %d actualizado(s) correctamente",rowsAffected, itemUpdated.ID)))
	}
}

// Función para eliminar un elemento
func DeleteItem(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idItem:= params["id"]

	rowsAffected,err := services.DeleteItem(idItem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if rowsAffected == 0 {
	  err = errors.New("El Objeto no existe")
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("El Objeto con id: %s fue eliminado correctamente", idItem)))
	}
}

func GetDetails(w http.ResponseWriter, r *http.Request)  {

}
