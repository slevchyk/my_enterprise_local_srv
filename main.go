package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kardianos/service"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/api/v1"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

var obx *objectbox.ObjectBox
var cfg models.Config
var logger service.Logger

func init() {
	var err error

	cfg, err = getConfig()
	if err != nil {
		panic(err)
	}

	obx, err = objectbox.NewBuilder().Directory(cfg.DatabaseConfig.Path).Model(models.ObjectBoxModel()).Build()

	if err != nil {
		panic(err)
	}
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        cfg.ServiceConfig.Name,
		DisplayName: cfg.ServiceConfig.DisplayName,
		Description: cfg.ServiceConfig.Description,
	}

	as := &apiServer{}
	s, err := service.New(as, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}

}

func getConfig() (models.Config, error) {
	var cfg models.Config

	// fullexecpath, err := os.Executable()
	// if err != nil {
	// 	return cfg, err
	// }

	// dir, _ := filepath.Split(fullexecpath)

	dir := os.Getenv("MYENTPRS")
	if dir == "" {
		msg := "environment variable \"MYENTPRS\" not specified"
		log.Println(msg)
		return cfg, errors.New(msg)
	}

	path := filepath.Join(dir, "config.json")

	f, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(f, &cfg)
	if err != nil {
		log.Println(err)
		return cfg, err
	}

	return cfg, nil
}

type apiServer struct{}

func (as *apiServer) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go as.run()
	return nil
}

func (as *apiServer) run() {
	// Do work here
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/deleteall", deleteAllHandeler)

	http.HandleFunc("/api/v1/auth", authHandler)

	http.HandleFunc("/api/v1/appuser", appUserHandler)
	http.HandleFunc("/api/v1/goods", goodsHandler)
	http.HandleFunc("/api/v1/goodsgroup", goodsGroupHandler)
	http.HandleFunc("/api/v1/harvesttype", harvestTypeHandler)
	http.HandleFunc("/api/v1/storage", storageHandler)
	http.HandleFunc("/api/v1/person", personHandler)
	http.HandleFunc("/api/v1/serviceworker", serviceWorkerHandler)
	http.HandleFunc("/api/v1/subdivision", subdivisionHandler)
	http.HandleFunc("/api/v1/locality", localityHandler)
	http.HandleFunc("/api/v1/unit", unitHandler)
	http.HandleFunc("/api/v1/vehicle", vehicleHandler)
	http.HandleFunc("/api/v1/trailer", trailerHandler)
	http.HandleFunc("/api/v1/consignmentnotein", consignmentnoteinHandler)

	//app
	http.HandleFunc("/api/app/v1/consignmentnotein", appConsignmentnoteinHandler)
	http.HandleFunc("/api/app/v1/consignmentnotein/processed", appConsignmentnoteinProcessedHandler)
	http.HandleFunc("/api/app/v1/consignmentnotein/changed", appConsignmentnoteinChangedHandler)

	port := fmt.Sprintf(":%v", cfg.ServerConfig.Port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
	defer obx.Close()
}

func (as *apiServer) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	var message = "Hi there!"

	w.Write([]byte(message))
	w.WriteHeader(http.StatusOK)
}

func deleteAllHandeler(w http.ResponseWriter, r *http.Request) {
	BoxForAppUser := models.BoxForAppUser(obx)
	BoxForAppUser.RemoveAll()

	BoxForGoods := models.BoxForGoods(obx)
	BoxForGoods.RemoveAll()

	BoxForGoodsGroup := models.BoxForGoodsGroup(obx)
	BoxForGoodsGroup.RemoveAll()

	BoxForHarvestType := models.BoxForHarvestType(obx)
	BoxForHarvestType.RemoveAll()

	BoxForStorage := models.BoxForStorage(obx)
	BoxForStorage.RemoveAll()

	box := models.BoxForAppUser(obx)
	box.RemoveAll()

	BoxForPerson := models.BoxForPerson(obx)
	BoxForPerson.RemoveAll()

	BoxForSubdivision := models.BoxForSubdivision(obx)
	BoxForSubdivision.RemoveAll()

	BoxForUnit := models.BoxForUnit(obx)
	BoxForUnit.RemoveAll()

	BoxForVehicle := models.BoxForVehicle(obx)
	BoxForVehicle.RemoveAll()

	//документи

	BoxForConsignmentNoteIn := models.BoxForConsignmentNoteIn(obx)
	BoxForConsignmentNoteIn.RemoveAll()

	BoxForGoodsConsignmentNoteIn := models.BoxForGoodsConsignmentNoteIn(obx)
	BoxForGoodsConsignmentNoteIn.RemoveAll()

	w.WriteHeader(http.StatusOK)
}

func authHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)
	api.AppUserAuth(w, r)
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

func serviceWorkerHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.ServiceWorkerPost(w, r)
	} else if r.Method == http.MethodGet {
		api.ServiceWorkerGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func subdivisionHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.SubdivisionPost(w, r)
	} else if r.Method == http.MethodGet {
		api.SubdivisionGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func localityHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.LocalityPost(w, r)
	} else if r.Method == http.MethodGet {
		api.LocalityGet(w, r)
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

func trailerHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.TrailerPost(w, r)
	} else if r.Method == http.MethodGet {
		api.TrailerGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func consignmentnoteinHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.ConsignmentNoteInPost(w, r)
	} else if r.Method == http.MethodGet {
		api.ConsignmentNoteInGet(w, r)
	} else if r.Method == http.MethodDelete {
		api.ConsignmentNoteInDelete(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func appConsignmentnoteinHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.ConsignmentNoteInAppPost(w, r)
	} else if r.Method == http.MethodGet {
		api.ConsignmentNoteInAppGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func appConsignmentnoteinProcessedHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.ConsignmentNoteInAppProcessed(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}
func appConsignmentnoteinChangedHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodGet {
		api.ConsignmentNoteInAppChanged(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}
