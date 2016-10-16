/**
 * Server configuration and route handlers
 */
package main;

import (
    "fmt"
    "log"
    "net/http"
    "github.com/spf13/viper"
)

/**
 * Handle setup route
 * 1: Create a private key and save it to disk
 */
func setupHandler(w http.ResponseWriter, r *http.Request) {
  // Create a private key and save to disk
  err := createKey()
  if (err != nil) { log.Fatal(err) }
}

/**
 * Initialize the http server and listen on configured port
 */
func server() {
  httpPort := viper.GetInt("root.port")
  http.HandleFunc("/", setupHandler)
  listen := fmt.Sprintf(":%d", httpPort)
  http.ListenAndServe(listen, nil)
}
