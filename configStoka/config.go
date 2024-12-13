package configStoka

import (
	"fmt"
	"todo1/configBd"
)

func StrokaConectBd(load *configBd.ConfBdStruct) string {
	stroka := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		load.DBhost, load.DBport, load.DBuser, load.DBpassword, load.DBname)
	return stroka

}
