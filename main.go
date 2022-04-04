package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
)

// var db *gorm.DB
var db *gorm.DB
var err error

// Product adl representasi dari sebuah produk
type Product struct {
	ID    int             `json:"id"`
	Code  string          `json:"code"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price" sql:"type:decimal(16,2)"`
}

// Result adl array dari produk
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

/**
function main
setting db format:
gorm.Open("mysql", "[USERNAME]:[PASSWORD]@/[DB_NAME]?charset=utf8&parseTime=True")
*/
func main() {
	db, err = gorm.Open("mysql", "root:Bee123456!@/go_learning?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("koneksi gagal", err)
	} else {
		log.Println("koneksi berhasil")
	}

	db.AutoMigrate(&Product{})

	handleRequests()

}

func handleRequests() {
	log.Println("start server dev at http://localhost:9999")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/products", createProduct).Methods("POST")
	myRouter.HandleFunc("/api/products", getProducts).Methods("GET")
	myRouter.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	myRouter.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	myRouter.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

/**
url homepage
*/
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

/**
create product
*/
func createProduct(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var product Product
	json.Unmarshal(payloads, &product)

	db.Create(&product)

	res := Result{Code: 200, Data: product, Message: "berhasil create produk"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

/**
get list product
*/
func getProducts(w http.ResponseWriter, r *http.Request) {

	product := []Product{}

	db.Find(&product)

	res := Result{Code: 200, Data: product, Message: "sukses get products"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

/**
get product by id
*/
func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var product Product
	db.First(&product, productID)

	res := Result{Code: 200, Data: product, Message: "sukses get product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

/**
update product
*/
func updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var productUpdates Product
	json.Unmarshal(payloads, &productUpdates)

	// db.Create(&product)
	var product Product
	db.First(&product, productID)
	db.Model(&product).Updates(productUpdates)

	res := Result{Code: 200, Data: product, Message: "berhasil update produk"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

/**
delete product
*/
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	model := mux.Vars(r)
	productID := model["id"]

	var product Product
	db.First(&product, productID)
	db.Delete(&product)

	res := Result{Code: 200, Message: "Product deleted successfully"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
