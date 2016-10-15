package main;

import (
    "fmt"
    "net/http"
    "github.com/spf13/viper"
)


func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello world")
  privKey, _ := createKeys()
  fmt.Println(privKey)
}

func main() {
  loadConfig()
  err := viper.ReadInConfig()
  if (err != nil) { fmt.Println(err) }
  rpcPort := viper.GetInt("rpc.port")
  rpcHost := viper.GetString("rpc.host")
  httpPort := viper.GetInt("root.port")

  fmt.Println(fmt.Sprintf("sMeter running on port %d", httpPort))
  fmt.Println(fmt.Sprintf("Communicating with RPC %s:%d", rpcHost, rpcPort))

  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

func loadConfig() {
  // Get the configurations for setup:
  // - RPC server to talk to
  // - Local server to listen on
  viper.SetConfigType("json")
  viper.SetConfigName("setup")
  viper.AddConfigPath("$GOPATH/config/")
}
