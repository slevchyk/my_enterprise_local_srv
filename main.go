package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kardianos/service"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/api/v1"
	"github.com/slevchyk/my_enterprise_local_srv/core"
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

	http.HandleFunc("/api/v1/check", basicAuth(checkHandler))
	http.HandleFunc("/api/v1/auth", basicAuth(authHandler))

	http.HandleFunc("/api/v1/appuser", basicAuth(appUserHandler))
	http.HandleFunc("/api/v1/appuser/cnirecipient", basicAuth(appUserCnirecipientHandler))
	http.HandleFunc("/api/v1/goods", basicAuth(goodsHandler))
	http.HandleFunc("/api/v1/goodsgroup", basicAuth(goodsGroupHandler))
	http.HandleFunc("/api/v1/harvesttype", basicAuth(harvestTypeHandler))
	http.HandleFunc("/api/v1/storage", basicAuth(storageHandler))
	http.HandleFunc("/api/v1/person", basicAuth(personHandler))
	http.HandleFunc("/api/v1/serviceworker", basicAuth(serviceWorkerHandler))
	http.HandleFunc("/api/v1/subdivision", basicAuth(subdivisionHandler))
	http.HandleFunc("/api/v1/locality", basicAuth(localityHandler))
	http.HandleFunc("/api/v1/unit", basicAuth(unitHandler))
	http.HandleFunc("/api/v1/vehicle", basicAuth(vehicleHandler))
	http.HandleFunc("/api/v1/trailer", basicAuth(trailerHandler))
	http.HandleFunc("/api/v1/consignmentnotein", basicAuth(consignmentnoteinHandler))

	//app
	http.HandleFunc("/api/app/v1/appuser", basicAuth(appUserAppHandler))
	http.HandleFunc("/api/app/v1/appuser/cnirecipient", basicAuth(appUserCnirecipientHandler))

	http.HandleFunc("/api/app/v1/consignmentnotein", basicAuth(appConsignmentnoteinHandler))
	http.HandleFunc("/api/app/v1/consignmentnotein/processed", basicAuth(appConsignmentnoteinProcessedHandler))
	http.HandleFunc("/api/app/v1/consignmentnotein/changed", basicAuth(appConsignmentnoteinChangedHandler))

	port := fmt.Sprintf(":%v", cfg.ServerConfig.Port)
	err := http.ListenAndServe(port, nil)
	// err := http.ListenAndServeTLS(port, cfg.ServerC6nfig.TlsCert, cfg.ServerConfig.TlsKey, nil)
	if err != nil {
		panic(err)
	}
	defer obx.Close()
}

func basicAuth(pass func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		isMain := "false"

		if len(pair) != 2 || !basicAppValidate(pair[0], pair[1]) {
			if len(pair) != 2 || !basicAdminValidate(pair[0], pair[1]) {
				http.Error(w, "authorization failed", http.StatusUnauthorized)
				return
			}
			isMain = "true"
		}

		form, _ := url.ParseQuery(r.URL.RawQuery)
		form.Del("is_main")
		form.Add("is_main", isMain)
		r.URL.RawQuery = form.Encode()

		pass(w, r)
	}
}

func basicAppValidate(username, password string) bool {
	if username == cfg.MobileAuth.User && password == cfg.MobileAuth.Password {
		return true
	}
	return false
}

func basicAdminValidate(username, password string) bool {
	if username == cfg.MainAuth.User && password == cfg.MainAuth.Password {
		return true
	}
	return false
}

func bearerAuth(pass func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Bearer" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		if !bearerValidate(auth[1], r) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		pass(w, r)
	}
}

func bearerValidate(token string, r *http.Request) bool {

	//формат токена не вірний
	pair := strings.SplitN(string(token), "|", 2)
	if len(pair) != 2 {
		return false
	}

	hash := core.EncodeHmac("password", pair[1])
	//підпис токена не відповідає вісту токена
	if pair[0] != hash {
		return false
	}

	var accessToken models.AccessToken

	err := core.DecodeFromBase64(&accessToken, pair[1])
	//структура токена пощкоджена
	if err != nil {
		return false
	}

	//час дії токена вийшов
	if accessToken.ExpiresAt.Before(time.Now()) {
		return false
	}

	form, _ := url.ParseQuery(r.URL.RawQuery)
	form.Del("auid")
	form.Add("auid", strconv.Itoa(int(accessToken.Id)))
	r.URL.RawQuery = form.Encode()

	return true
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

	fvIsMain := r.FormValue("is_main")
	if fvIsMain == "false" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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

	BoxForLocality := models.BoxForLocality(obx)
	BoxForLocality.RemoveAll()

	BoxForUnit := models.BoxForUnit(obx)
	BoxForUnit.RemoveAll()

	BoxForVehicle := models.BoxForVehicle(obx)
	BoxForVehicle.RemoveAll()

	BoxForTrailer := models.BoxForTrailer(obx)
	BoxForTrailer.RemoveAll()

	//документи

	BoxForConsignmentNoteIn := models.BoxForConsignmentNoteIn(obx)
	BoxForConsignmentNoteIn.RemoveAll()

	BoxForGoodsConsignmentNoteIn := models.BoxForGoodsConsignmentNoteIn(obx)
	BoxForGoodsConsignmentNoteIn.RemoveAll()

	w.WriteHeader(http.StatusOK)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func authHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)
	api.AppUserAuth(w, r)
}

func appUserHandler(w http.ResponseWriter, r *http.Request) {

	fvIsMain := r.FormValue("is_main")
	if fvIsMain == "false" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		api.AppUserPost(w, r)
	} else if r.Method == http.MethodGet {
		api.AppUserGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func appUserCnirecipientHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		api.AppUserCniRecipientPost(w, r)
	} else if r.Method == http.MethodGet {
		api.AppUserCniRecipientGet(w, r)
	} else if r.Method == http.MethodDelete {
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		api.AppUserCniRecipientDelete(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func goodsHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		api.GoodsPost(w, r)
	} else if r.Method == http.MethodGet {
		api.GoodsGet(w, r)
	} else if r.Method == http.MethodDelete {
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		api.GoodsDeleteAll(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func goodsGroupHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	if r.Method == http.MethodPost {
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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
		fvIsMain := r.FormValue("is_main")
		if fvIsMain == "false" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		api.TrailerPost(w, r)
	} else if r.Method == http.MethodGet {
		api.TrailerGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func consignmentnoteinHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)

	fvIsMain := r.FormValue("is_main")
	if fvIsMain == "false" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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

//APPLICATION APIs
func appUserAppHandler(w http.ResponseWriter, r *http.Request) {

	fvIsMain := r.FormValue("is_main")
	if fvIsMain == "true" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	api := api.NewApiV1(obx)

	if r.Method == http.MethodGet {
		api.AppUserAppGet(w, r)
	} else {
		http.Error(w, "method not specified", http.StatusBadRequest)
	}

}

func appConsignmentnoteinHandler(w http.ResponseWriter, r *http.Request) {

	api := api.NewApiV1(obx)
	api.SetAbn(r)

	// fvIsMain := r.FormValue("is_main")
	// if fvIsMain == "false" {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

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
