package integration

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"tc-demo/config"
	"tc-demo/containers"
	"tc-demo/database"
	"tc-demo/internal/app/handler"
	"tc-demo/internal/app/model"
	"tc-demo/internal/app/repository"
	"tc-demo/internal/app/router"
	"tc-demo/internal/app/service"
	"testing"
)

const (
	MigrationPath = "file://../db/migrations"
	FixturePath   = "../db/fixtures/product.yml"
)

var (
	db     *sql.DB
	ctx    context.Context
	server *httptest.Server
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	cfg := config.NewDatabaseConfig()

	postgresC, err := containers.NewPostgresContainer(ctx, cfg)

	if err != nil {
		log.Fatalf("could not start database container: %v", err)
	}
	defer postgresC.Terminate(ctx)

	err = cfg.Update(ctx, postgresC)

	if err != nil {
		log.Fatalf("could not fetch test container configuration: %v", err)
	}

	connString := cfg.DSN()

	log.Println("✅", "Connection String:", connString)

	db, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("could not create database connection string: %v", err)
	}
	defer db.Close()

	// Wait for the database to be ready
	err = database.WaitAndPingDBWithRetry(db, 30)

	if err != nil {
		log.Fatalf("could not connect to the container database: %v", err)
	}

	err = database.MigrateUP(db, MigrationPath)

	if err != nil {
		log.Fatalf("could not run migration: %v", err)
	}

	err = database.LoadProductFixtures(db, FixturePath)

	if err != nil {
		log.Fatalf("could not load fixtures: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	productRepository := repository.NewProductRepository(gormDB)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	router := router.NewRouter(productHandler)
	server = httptest.NewServer(router)
	defer server.Close()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestGetProducts(t *testing.T) {

	resp, err := http.Get(server.URL + "/v1/products/")
	if err != nil {
		t.Fatalf("Failed to send GET request to /v1/products/ endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	var products []model.Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		t.Fatalf("Failed to decode JSON response products: %v", err)
	}
	expectedProductCount := 10
	if len(products) != expectedProductCount {
		t.Fatalf(`Expected %d products, but got %d`, expectedProductCount, len(products))
	}
}

func TestGetProduct(t *testing.T) {
	resp, err := http.Get(server.URL + "/v1/products/1")
	if err != nil {
		t.Fatalf("Failed to send GET request to /v1/products/1 endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	var product model.Product
	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		t.Fatalf("Failed to decode JSON response products: %v", err)
	}
	expectedProductID := uint(1)
	if product.ID != expectedProductID {
		t.Fatalf("Expected product ID to be %d but got %d", expectedProductID, product.ID)
	}
}

func TestCreateProduct(t *testing.T) {
	product := &model.Product{
		Name:        "ZEBRONICS Zeb-Bro",
		Description: "Ear Wired Earphones with Mic, 3.5mm Audio Jack, 10mm Drivers",
		Price:       129.99,
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		t.Fatalf("Failed to marshall product to JSON: %v", err)
	}

	// Send a POST request to create the product
	resp, err := http.Post(server.URL+"/v1/products/", "application/json", bytes.NewBuffer(productJSON))
	if err != nil {
		t.Fatalf("Failed to send POST request to /v1/products/ endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d but got %d", http.StatusCreated, resp.StatusCode)
	}

	var createdProduct model.Product
	err = json.NewDecoder(resp.Body).Decode(&createdProduct)

	if err != nil {
		t.Fatalf("Failed to decode JSON response products: %v", err)
	}

	expectedProductPrice := 129.99

	if createdProduct.Price != expectedProductPrice {
		t.Fatalf("Expected product ID to be %f but got %f", expectedProductPrice, product.Price)
	}
}

func TestUpdateProduct(t *testing.T) {
	product := &model.Product{
		Price: 399.99,
	}

	jsonBody, err := json.Marshal(product)
	if err != nil {
		t.Fatalf("Failed to marshall product to JSON: %v", err)
	}

	req, err := http.NewRequest("PUT", server.URL+"/v1/products/1", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create PUT request object for /v1/products/1 endpoint: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send PUT request to /v1/products/1 endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	var updatedProduct model.Product
	err = json.NewDecoder(resp.Body).Decode(&updatedProduct)

	if err != nil {
		t.Fatalf("Failed to decode JSON response products: %v", err)
	}

	expectedProductPrice := 399.99

	if updatedProduct.Price != expectedProductPrice {
		t.Fatalf("Expected product ID to be %f but got %f", expectedProductPrice, product.Price)
	}
}

func TestDeleteProduct(t *testing.T) {
	req, err := http.NewRequest("DELETE", server.URL+"/v1/products/1", nil)

	if err != nil {
		t.Fatalf("Failed to create DELETE request object for /v1/products/1 endpoint: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to send DELETE request to /v1/products/1 endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("Expected status code %d but got %d", http.StatusNoContent, resp.StatusCode)
	}
}
