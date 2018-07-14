package user

import (
  "time"
  "net/http"
  "fmt"
  "log"
  "gopkg.in/mgo.v2"
  "github.com/gorilla/mux"
  "gopkg.in/mgo.v2/bson"
  "../../../config"
)

type User struct {
    Id              string    `json: "id"`
    Username        string    `json: "username"`
    Password        string    `json: password`
    FirstName       string    `json: firstName`
    LastName        string    `json: lastName`
    Gender          string    `json: gender`
    BirthDate       time.Time `json: birthDate`
  }

func GetUsers(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, conf.GetDbUrl())
}

func CreateUser(w http.ResponseWriter, r *http.Request){
  session, err := mgo.Dial(conf.GetDbUrl())
    if err != nil {
            panic(err)
    }
    defer session.Close()

   c := session.DB(conf.GetDbName()).C(conf.GetDbUsersCollecion())
   err = c.Insert(&User{Id: "2", Username: "GustavoAdrianGimenez", Password: "1234", FirstName: "Gustavo", LastName: "Gimenez", Gender: "Male", BirthDate: time.Now()})
   if err != nil {
           log.Fatal(err)
   }

}

func GetUser(w http.ResponseWriter, r *http.Request){
  session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
            panic(err)
    }
    defer session.Close()

    ps := mux.Vars(r)

    userID := ps["userID"]
    c := session.DB("local").C("users")
    result := User{}
    err = c.Find(bson.M{"id": userID}).One(&result)
    if err != nil {
            log.Fatal(err)
    }

    fmt.Fprintln(w, result)
}
