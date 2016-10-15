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
 */
func setupHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello world")
  // Create a private key and save to disk
  err := createKey()
  if (err != nil) { log.Fatal(err) }

  // Test: grab private key and print string
  
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
