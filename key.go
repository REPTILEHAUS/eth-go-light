/*
Everything related to creating, saving, and loading a private key
as well as recovering an Ethereum address from it
 */
package main;

import (
    "crypto/rand"
    "encoding/hex"
    "github.com/spf13/viper"
    "io/ioutil"
)

/**
 * Generate a private key and save it to disk
 *
 * @returns (error)
 */
func createKey() (error) {
  b, err := generateRandomBytes(viper.GetInt("key.nBytes"))
  if (err == nil) {
    err2 := keyToFile(b)
    if (err2 != nil) { return err2 }
  } else {
    return err
  }
  return nil
}

/**
 * Get the private key from a file
 *
 * @returns (string, error) - private key (hex), error
 */
func getKey() (string, error) {
  b, err := keyFromFile()
  key := hex.EncodeToString(b)
  return key, err
}

/**
 * Dump bytes to a file
 *
 * @param b {bytes} - arbitrary byte array
 * @returns (error)
 */
func keyToFile(b []byte) (error) {
  fpath := viper.GetString("key.path")
  err := ioutil.WriteFile(fpath, b, 0644)
  return err
}

/**
 * Read a file containing a byte array
 *
 * @returns ([]byte, error) - private key, error
 */
func keyFromFile() ([]byte, error) {
  fpath := viper.GetString("key.path")
  b, err := ioutil.ReadFile(fpath)
  return b, err
}

/**
 * Generate some random bytes
 *
 * @param n {int} - number of bytes to generate
 * @returns []bytes, error - byte array and error object
 */
func generateRandomBytes(n int) ([]byte, error) {
  b := make([]byte, n)
  _, err := rand.Read(b)
  if (err != nil) { return nil, err }
  return b, nil
}
