package main

func main() {
	server := &Server{}
	if err := server.Run(); err != nil {
		panic(err)
	}
}
