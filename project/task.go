package project

import (
	"errors"

	"github.com/google/uuid"
)

/*
Создание задач
*/

/*
Реализуйте структуру Task, которая будет содержать уникальный идентификатор, заголовок, описание и статус задачи (активная или закрытая).
*/
type Status string

const (
	StatusActive Status = "ACTIVE"
	StatusClosed Status = "CLOSED"
)

type Task struct {
	UUID        uuid.UUID
	TaskName    string
	Description string
	TaskStatus  Status
}

/*
Напишите функцию NewTask, которая будет создавать новую задачу. Задача не может быть создана с пустым заголовком или описанием. Задача всегда создается с активным статусом.
*/
func NewTask(UUID uuid.UUID, taskName, description string) (*Task, error) {
	if taskName == "" {
		return nil, errors.New("Передан пустой taskName")
	}

	if description == "" {
		return nil, errors.New("Передан пустой description")
	}

	return &Task{
		UUID:        UUID,
		TaskName:    taskName,
		Description: description,
		TaskStatus:  StatusActive,
	}, nil
}

/*
Обеспечьте возможность обновления описания задачи только для активных задач с помощью метода UpdateDescription. Не забудьте проверить то,
что проверяли в функции-конструкторе.
*/
func (t *Task) UpdateDescription(description string) error {
	if description == "" {
		return errors.New("Передан пустой description")
	}
	if t.TaskStatus != StatusActive {
		return errors.New("нельзя обновить описание закрытой задачи")
	}

	t.Description = description
	return nil
}

/*
Закрытие задач
Реализуйте метод Close для структуры Task, который будет изменять статус задачи на закрытый, если статус уже закрытый, то необходимо вернуть ошибку.
*/
func (t *Task) Close() error {
	if t.TaskStatus == StatusClosed {
		return errors.New("Задача уже закрыта")
	}
	t.TaskStatus = StatusClosed
	return nil
}
