package conf

import (
  "encoding/json"
  "os"
  "fmt"
  "log"
)

type Configuration  struct {
  DB_URL          string
  DB_NAME         string
  DB_COLLECTION   string
}

func NewConfig(environment string) Configuration {

  config := Configuration{}

  file, err := os.Open("./config/config." + environment + ".json")

  if err != nil {
    log.Fatal("Error: ", err)
  }
  defer file.Close()
  decoder := json.NewDecoder(file)
  err = decoder.Decode(&config)

  if err != nil {
    fmt.Println("error:", err)
  }
  os.Setenv("DB_URL", config.DB_URL)
  os.Setenv("DB_NAME", config.DB_NAME)
  os.Setenv("DB_COLLECTION", config.DB_COLLECTION)

  return config

}

func GetDbUrl() string{
  return os.Getenv("DB_URL")
}

func GetDbName() string{
  return os.Getenv("DB_NAME")
}

func GetDbUsersCollecion() string{
  return os.Getenv("DB_COLLECTION")
}
