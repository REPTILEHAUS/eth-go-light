/*
Everything related to creating, saving, and loading a private key
as well as recovering an Ethereum address from it
 */
package main;

import (
    "crypto/rand"
    "encoding/base64"
    "github.com/spf13/viper"
)

/**
 * Generate a private key and save it to disk
 */
func createKeys() (string, error){
  b, err := generateRandomBytes(viper.GetInt("key.nBytes"))
  if (err != nil) {
    privKey := base64.URLEncoding.EncodeToString(b)
    return privKey, err
  } else {
    return "", err
  }
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
