package main

import "micro-example/service/gateway/router"

func main()  {
	r := router.Router()
	_ = r.Run("127.0.0.1:5000")
}