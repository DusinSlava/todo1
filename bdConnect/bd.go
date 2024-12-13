package bdConnect

import (
	"context"
	"fmt"
	"log"
	"todo1/configBd"
	"todo1/configStoka"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConectBD(bsc *configBd.ConfBdStruct, ctx context.Context) (*pgxpool.Pool, error) {
	stroka := configStoka.StrokaConectBd(bsc)
	parsStroki, err := pgxpool.ParseConfig(stroka)

	if err != nil {
		log.Fatal(err, "Ошибка на уровне подключения базы данных, данные не получислось распарсить!!!")
	}
	parsStroki.MaxConns = 10       // Максимальное количество соединений в пуле
	parsStroki.MinConns = 1        // Минимальное количество соединений в пуле
	parsStroki.MaxConnLifetime = 0 // Максимальное время жизни соединений
	parsStroki.MaxConnIdleTime = 0
	poll, err := pgxpool.NewWithConfig(ctx, parsStroki)
	if err != nil {
		return nil, fmt.Errorf("ошибка на уровне подключения бд, пулл соединение не создалось:%v", err)
	}
	err = poll.Ping(ctx)
	if err != nil {
		poll.Close()
		return nil, fmt.Errorf("не удалось подключиться к базе данных :%v", err)
	}
	log.Println("успешное подключение к базе данных")
	return poll, nil

}
