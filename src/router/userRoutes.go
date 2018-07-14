package router

import (
  "github.com/gorilla/mux"
  "../model/user"
)

var userRoutes = Routes{
  Route{
      "getUsers",
      "GET",
      "/users",
      user.GetUsers,
    },
  Route{
    "createUser",
    "POST",
    "/users",
    user.CreateUser,
    },
  Route{
    "getUser",
    "GET",
    "/users/{userID}",
    user.GetUser,
    },
}

func addUserRoutes(r *mux.Router){
  for _, route := range userRoutes {
      r.
          Methods(route.method).
          Path(route.pattern).
          Name(route.name).
          Handler(route.handlerFunc)
        }
}
