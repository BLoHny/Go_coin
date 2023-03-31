package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/BLoHny/Go_coin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml")) // If err != null

	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(port, nil))
}
