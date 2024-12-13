package configBd

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfBdStruct struct {
	DBhost     string
	DBport     int
	DBuser     string
	DBpassword string
	DBname     string
}

func LoadConfBd(cbs *ConfBdStruct) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка не удалось загрузить данные из переменных окружения", err)
	}
	cbs.DBhost = os.Getenv("DB_HOST")
	cbs.DBuser = os.Getenv("DB_USER")
	cbs.DBpassword = os.Getenv("DB_PASSWORD")
	cbs.DBname = os.Getenv("DB_NAME")
	if cbs.DBhost == "" || cbs.DBuser == "" || cbs.DBpassword == "" || cbs.DBname == "" {
		log.Fatal("Не все параметры подключения к базе данных установлены!")
	}
	DBportstr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(DBportstr)
	cbs.DBport = port
	if err != nil || port < 0 || port >= 65535 {
		log.Fatal(err, "Ошибка порт указан не верно")
	}

}
