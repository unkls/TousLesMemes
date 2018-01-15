package main

import (
  "touslesmemes/ZAPI/actions"
)

func main() {

  a := new(actions.App)

  a.InitializeDatabase()
  a.InitializeRoutes()
  a.Run()

}
