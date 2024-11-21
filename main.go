package main

func main() {
	DBInstance = NewDB()
	ListenAndServe(":3007", &RedisHandler{})
}
