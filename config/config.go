package config

import (
		"fmt"
		"net/http"
		"time"

	_ "github.com/lib/pq" // Driver de conexión con Postgres
)

func StartServer(port string, router http.Handler) error {
	server := &http.Server {
		Handler: router,
		Addr: port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	fmt.Printf("Starting server on port %s ..\n", port)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("Error while starting up server: %v", err)
	}
	
	return nil
}
