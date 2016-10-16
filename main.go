package main;

import (
    "fmt"
    "github.com/spf13/viper"
)

/**
 * Main function. Load the configuration and setup the web server.
 */
func main() {
  // Setup the config file
  loadConfig()
  err := viper.ReadInConfig()
  if (err != nil) { fmt.Println(err) }
  rpcPort := viper.GetInt("rpc.port")
  rpcHost := viper.GetString("rpc.host")
  fmt.Println(fmt.Sprintf("Communicating with RPC %s:%d", rpcHost, rpcPort))

  // Setup the server
  // fmt.Println(fmt.Sprintf("Server listening on port %d", viper.GetInt("root.port")))
  // server()

  // err = createKey()
  // if (err != nil) { log.Fatal(err) }
  priv, err := keyFromFile()
  privateToAddress(priv)

}

func loadConfig() {
  // Get the configurations for setup:
  // - RPC server to talk to
  // - Local server to listen on
  viper.SetConfigType("json")
  viper.SetConfigName("setup")
  viper.AddConfigPath("$GOPATH/config/")
}
