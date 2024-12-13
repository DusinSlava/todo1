package repozitory1

import (
	"context"

	"log"
	"todo1/tablePole"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateTask(ctx context.Context, task tablePole.Task) error
	GetTask(ctx context.Context) ([]tablePole.Task, error)
	DeleteTask(ctx context.Context, id int) error
	PutTask(ctx context.Context, task tablePole.Task) error
}
type bdRepository struct {
	pool *pgxpool.Pool
}

func NewBdRepository(pool *pgxpool.Pool) *bdRepository {
	return &bdRepository{pool: pool}
}
func (r *bdRepository) CreateTask(ctx context.Context, task tablePole.Task) error {
	query := "INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3) RETURNING id"
	return r.pool.QueryRow(ctx, query, task.Title, task.Description, task.Completed).Scan(&task.ID)
}
func (r *bdRepository) GetTask(ctx context.Context) ([]tablePole.Task, error) {
	var tasks []tablePole.Task
	query := "SELECT id, title, description, completed FROM tasks"
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		log.Println("Ошибка отправки get запроса", err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var temporaryTable tablePole.Task
		if err = rows.Scan(&temporaryTable.ID, &temporaryTable.Title, &temporaryTable.Description, &temporaryTable.Completed); err != nil {
			log.Println("Ошибка при Сканировании значений из таблицы")
			return nil, err
		}
		tasks = append(tasks, temporaryTable)
	}

	return tasks, nil
}
func (r *bdRepository) DeleteTask(ctx context.Context, id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	sendingRequest, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		log.Println("Ошибка в отправке запроса", err)
		return err
	}
	if sendingRequest.RowsAffected() == 0 {
		log.Println("Не найдена задача с ID=", id)
		return err
	}

	return nil
}
func (r *bdRepository) PutTask(ctx context.Context, tasks tablePole.Task) error {
	query := "UPDATE tasks SET title = $1, description = $2, completed = $3 WHERE id = $4"
	sendingRequest, err := r.pool.Exec(ctx, query, tasks.Title, tasks.Description, tasks.Completed, tasks.ID)
	if err != nil {
		log.Println("Ошибка в отправке запроса", err)
		return err
	}
	if sendingRequest.RowsAffected() == 0 {
		log.Println("Не найдена задача с ID=", tasks.ID)
		return err
	}

	return nil

}
