package main

import (
	"net/http"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/api/v1"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

var obx *objectbox.ObjectBox

func init() {
	var err error

	obx, err = objectbox.NewBuilder().Model(models.ObjectBoxModel()).Build()
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/api/v1/appuser", appUserHandler)
	http.HandleFunc("/api/v1/goods", goodsHandler)
	http.HandleFunc("/api/v1/goodsgroup", goodsGroupHandler)
	http.HandleFunc("/api/v1/harvesttype", harvestTypeHandler)
	http.HandleFunc("/api/v1/storage", storageHandler)
	http.HandleFunc("/api/v1/person", personHandler)
	http.HandleFunc("/api/v1/subdivision", subdivisionHandler)
	http.HandleFunc("/api/v1/unit", unitHandler)
	http.HandleFunc("/api/v1/vehicle", vehicleHandler)
	http.HandleFunc("/test", testHandler)
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		panic(err)
	}
	defer obx.Close()
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	var message string

	message = "Hi there!"

	w.Write([]byte(message))
	w.WriteHeader(http.StatusOK)
}

func appUserHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.AppUserPost(w, r)
	} else if r.Method == http.MethodGet {
		api.AppUserGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func goodsHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.GoodsPost(w, r)
	} else if r.Method == http.MethodGet {
		api.GoodsGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func goodsGroupHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.GoodsGroupPost(w, r)
	} else if r.Method == http.MethodGet {
		api.GoodsGroupGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func harvestTypeHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.HarvestTypePost(w, r)
	} else if r.Method == http.MethodGet {
		api.HarvestTypeGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func storageHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.StoragePost(w, r)
	} else if r.Method == http.MethodGet {
		api.StorageGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func personHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.PersonPost(w, r)
	} else if r.Method == http.MethodGet {
		api.PersonGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func subdivisionHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.SubdivisionGet(w, r)
	} else if r.Method == http.MethodGet {
		api.SubdivisionGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func unitHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.UnitPost(w, r)
	} else if r.Method == http.MethodGet {
		api.UnitGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func vehicleHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.VehiclePost(w, r)
	} else if r.Method == http.MethodGet {
		api.VehicleGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}
