package main

import "www/src/rooter"

func main() {
	rt := rooter.Init()

	// サーバー起動
	rt.Logger.Fatal(rt.Start(":8080"))
}
