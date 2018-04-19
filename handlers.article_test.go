package main

func TestShowIndexPageUnauthenticated() {
	r := getRouter(true)

	r.GET("/", showIndexPage)

}
