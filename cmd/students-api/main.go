package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config"
)

func main() {
	/*
		TO-DO
			1. load config
			2. database setup
			3. setup router
			4. setup server
	*/

	// ------------------------------------------- 1. load config----------------------------------------
	
	cfg:= config.MustLoad()
	
	// ------------------------------------------- 3. setup router----------------------------------------

	router := http.NewServeMux()
	router.HandleFunc("GET /",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Welcome to students API"))
	})

	server := http.Server{
		Addr         : cfg.Addr,
		Handler : router,
	}

	fmt.Printf("Server started: %s \n",cfg.Addr)
	slog.Info("Server start succes")

	done := make(chan os.Signal,1);
	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM);

	go func(){
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal("failed to start server")
	}
	
	}()
	<- done

	slog.Info("Shutting down the server: ")

	// server.Shutdown()		// This is inbuilt but can wait infinitely

	ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)		// context.background is an empty context
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil{
		slog.Error("failed to shut-down server",slog.String("error",err.Error()))
	}

	slog.Info("Shut-down sucessfully")


}