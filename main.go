package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	_ "github.com/Dmitry-Grizlov/swag-demo/docs"

	models "github.com/Dmitry-Grizlov/swag-demo/models"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var GlobalItemsList = []models.Item{
	{
		Id:    1,
		Value: "default",
	},
}

// @title			Project Name
// @version		0.0.1
// @host			localhost:3000
// @BasePath		/
// @contact.name	company_name
// @contact.url	http://www.link-to-support.io/support
// @contact.email	some.company.name@company.domain.com
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/list", list)
	r.HandleFunc("/list/{id}", one)
	r.HandleFunc("/update/{id}/{val}", update)
	r.HandleFunc("/add", add)
	r.HandleFunc("/delete/{id}", delete)
	r.HandleFunc("/misc", misc)
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), // The URL to fetch the Swagger JSON documentation.
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	log.Fatal(http.ListenAndServe(":3000", r))
}

// @title			list godoc
// @Summary		list items
// @Description	list all items from global var
// @Tags			items
// @Produce		json
// @Success		200	{object}	[]models.Item
// @Router			/list [get]“
func list(w http.ResponseWriter, r *http.Request) {
	obj := GlobalItemsList
	ret, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshallng object"))
	}

	w.Write(ret)
}

// @title			one godoc
// @Summary		single item
// @Description	returns item with input id from global var
// @Tags			items
// @Produce		json
// @Success		200	{object}	models.Item
// @Router			/list/{id} [get]“
// @Param			id	path	int	true	"Id of the item"
func one(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	if itemID == "" {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(itemID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting value"))
	}

	item := &models.Item{}
	for _, i := range GlobalItemsList {
		if id == i.Id {
			item = &i
			break
		}
	}

	if item == nil {
		http.NotFound(w, r)
		return
	}

	ret, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshallng object"))
	}

	w.Write(ret)
}

// @title			update godoc
// @Summary		updates item
// @Description	updates item from global var
// @Tags			items
// @Produce		json
// @Success		200	{string}	string
// @Router			/update/{id}/{val} [get]“
// @Param			id	path	int		true	"Id of the item"
// @Param			val	path	string	true	"value of the item"
func update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	if itemID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad id %"))
		return
	}

	val := vars["val"]
	if val == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad val"))
		return
	}

	id, err := strconv.Atoi(itemID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshallng object"))
	}

	notFound := true
	for i := range GlobalItemsList {
		if id == GlobalItemsList[i].Id {
			notFound = false
			GlobalItemsList[i].Value = val
			break
		}
	}

	if notFound {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Updated!"))
}

// @title			add godoc
// @Summary		adds item
// @Description	adds item to global var
// @Tags			items
// @Produce		json
// @Success		200	{string}	string
// @Router			/add [get]“
func add(w http.ResponseWriter, r *http.Request) {
	id := rand.Intn(100)

	GlobalItemsList = append(GlobalItemsList, models.Item{
		Id:    id,
		Value: fmt.Sprintf("val%d", id),
	})

	w.Write([]byte("Added"))
}

// @title			delete godoc
// @Summary		deletes item
// @Description	deletes item from global var with input id
// @Tags			items
// @Produce		json
// @Success		200	{string}	string
// @Router			/delete/{id} [get]“
// @Param			id	path	int	true	"Id of the item"
func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	if itemID == "" {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(itemID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshallng object"))
	}

	notFound := true
	tempList := []models.Item{}
	for _, i := range GlobalItemsList {
		if id == i.Id {
			notFound = false
			continue
		}
		tempList = append(tempList, i)
	}

	if notFound {
		http.NotFound(w, r)
		return
	}

	GlobalItemsList = tempList

	w.Write([]byte("Deleted"))
}

// @title			misc godoc
// @Summary		Miscellanious action
// @Description	Miscellanious action to be displayed in the separate method group
// @Tags			misc
// @Produce		json
// @Success		200	{string}	string
// @Router			/misc [get]“
func misc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Misc responded with success"))
}
