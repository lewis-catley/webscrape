package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lewis-catley/webscrape/urls/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
	urls    *mongodb.URLModel
}

type GetUrl struct {
	ID        string            `json:"id"`
	Url       string            `json:"url"`
	UrlsFound map[string]string `json:"urlsFound"`
	Count     int               `json:"count"`
}

type PostUrl struct {
	Url string `json:"url"`
}

func main() {
	app := application{
		infoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errLog:  log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	app.infoLog.Printf("Initialising urls API service")
	app.infoLog.Printf("Reatrieving all command line args")

	// Command line flags
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")
	mongoURI := flag.String("mongoURI", "mongodb://localhost:27017", "Database hostname url")
	mongoDatabase := flag.String("mongoDatabase", "urls", "Database name")
	enableCredentials := flag.Bool("enableCredentials", false, "Enable the use of credentials for mongo connection")
	flag.Parse()

	// Create mongo config
	mc := options.Client().ApplyURI(*mongoURI)
	if *enableCredentials {
		mc.Auth = &options.Credential{
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
		}
	}

	// Establish mongo connection
	client, err := mongo.NewClient(mc)
	if err != nil {
		app.errLog.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		app.errLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			app.errLog.Fatal(err)
		}
	}()

	app.infoLog.Printf("Database connection established")
	app.urls = &mongodb.URLModel{
		C: client.Database(*mongoDatabase).Collection("urls"),
	}

	// Create the server
	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     app.errLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.infoLog.Printf("Starting URLS server on %s", serverURI)
	app.errLog.Fatal(srv.ListenAndServe())
}
