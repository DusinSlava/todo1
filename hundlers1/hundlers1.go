package hundlers1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todo1/repozitory1"
	"todo1/tablePole"
)

func Add(repo repozitory1.Repository, w http.ResponseWriter, r *http.Request, ctx context.Context) {
	switch r.Method {
	case http.MethodPost:
		var task tablePole.Task
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Запрос составлен не верно", http.StatusBadRequest)
			return
		}
		err = repo.CreateTask(ctx, task)
		if err != nil {
			http.Error(w, "Не удалось добавить задачу", http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(task)
		if err != nil {
			http.Error(w, "Ошибка отправки ответа", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Задача успешно добавлена"))
		log.Println("Задача успешно добавлена:", task)
	case http.MethodGet:
		tasks, err := repo.GetTask(ctx)
		if err != nil {
			http.Error(w, "Ошибка получения задач", http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(tasks)
		if err != nil {
			http.Error(w, "Не удалось отправить ответ", http.StatusInternalServerError)
			log.Println("Ошибка отправки данных клиенту в hundler")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Задачи успешно Получены"))
		log.Println("Задачи успешно отправлены клиенту")

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "ID обязателен, укажите ID", http.StatusBadRequest)
		}
		taskId, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Некорректный ID задачи", http.StatusBadRequest)
			return
		}
		err = repo.DeleteTask(ctx, taskId)
		if err != nil {
			http.Error(w, "Задача не удалось удалить", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Задача удалена успешно"))
		log.Println("Задачи успешно удалена")

	case http.MethodPut:
		var tasks tablePole.Task
		err := json.NewDecoder(r.Body).Decode(&tasks)
		if err != nil {
			http.Error(w, "Запрос составлен не верно", http.StatusBadRequest)
			return
		}
		err = repo.PutTask(ctx, tasks)
		if err != nil {
			http.Error(w, "Не удалось обновить задачу", http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Задача обновлена успешно"))
		log.Println("Задачи уcпешно обновлена")

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
