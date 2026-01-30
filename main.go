package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"

	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	// 1. Setup Config using Viper
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// 2. Setup Database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// 3. Dependency Injection - Products
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// 4. Dependency Injection - Categories
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// 5. Routes
	// Product Routes
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	// Category Routes
	http.HandleFunc("/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/categories/", categoryHandler.HandleCategoryByID)

	// Health Check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//fmt.Fprintf(w, "API Running")
		// return json
		fmt.Fprint(w, `{"status": "OK", "message": "API Running"}`)
	})

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
