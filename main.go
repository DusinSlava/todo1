package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"todo1/bdConnect"
	"todo1/configBd"
	"todo1/hundlers1"
	"todo1/repozitory1"
)

func main() {
	confbd := &configBd.ConfBdStruct{}
	configBd.LoadConfBd(confbd)
	ctx := context.Background()
	bdPoll, err := bdConnect.ConectBD(confbd, ctx)
	if err != nil {
		log.Fatal(err, "Ошибка, пулл соединения не создался")
	}
	defer bdPoll.Close()
	repo := repozitory1.NewBdRepository(bdPoll)
	startServer(repo, ctx)

}
func startServer(repo repozitory1.Repository, ctx context.Context) {
	http.HandleFunc("/add-tasks", func(w http.ResponseWriter, r *http.Request) {
		hundlers1.Add(repo, w, r, ctx)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервиса", err)
	}
	fmt.Println("Сервер успешно запущена на порту 8080")
}
