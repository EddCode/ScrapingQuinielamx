package main

func main() {
	serve := InitServer(":9000")
	serve.Run()
}
