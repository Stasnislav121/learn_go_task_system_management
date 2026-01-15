package project

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

/*
Создание проекта
*/

/*
Реализуйте структуру Project, которая будет содержать уникальный идентификатор, название проекта и список задач.
*/
type Project struct {
	UUID    uuid.UUID
	Name    string
	TasksId []uuid.UUID
	Tasks   map[uuid.UUID]Task
}

/*
Напишите функцию New, которая будет создавать новый проект. Проект не может быть создан с пустым названием.
*/

func New(id uuid.UUID, name string) (*Project, error) {
	if name == "" {
		return nil, errors.New("Передан пустой name")
	}
	return &Project{
		UUID:    id,
		Name:    name,
		TasksId: make([]uuid.UUID, 0),
		Tasks:   make(map[uuid.UUID]Task),
	}, nil
}

/*
Добавление и обновление задач
Реализуйте метод AddTask для структуры Project, который будет добавлять задачи в проект (если ID задачи повторяется, нужно выбросить ошибку).
*/
func (p *Project) AddTask(task Task) error {
	if _, exist := p.Tasks[task.UUID]; exist {
		return errors.New("Задача уже существует")
	}
	p.TasksId = append(p.TasksId, task.UUID)
	p.Tasks[task.UUID] = task
	return nil
}

/*
Вывод информации
Реализуйте метод PrintInfo для структуры Project, который будет выводить информацию о проекте и всех его задачах.
*/
func (p Project) PrintInfo() {
	fmt.Printf("ID проекта: %s,\n Название проекта: %s\n\n", p.UUID, p.Name)
	fmt.Println("Задачи:")
	for _, id := range p.TasksId {
		task := p.Tasks[id]

		fmt.Printf("ID задачи: %s,\n Название задачи: %s,\n Описание задачи: %s,\n Статус задачи: %s\n", id, task.TaskName, task.Description, task.TaskStatus)
		fmt.Println()
	}
}

/*
Реализуйте метод UpdateTask, который будет обновлять информацию о задаче в проекте. Если в проекте нет переданной задачи, нужно вернуть ошибку.
*/
func (p *Project) UpdateTask(task Task) error {
	if _, exist := p.Tasks[task.UUID]; !exist {
		return errors.New("Задача не существует")
	}
	p.Tasks[task.UUID] = task
	return nil
}

/*
Фильтрация задач
Реализуйте метод FilterTasksByStatus, который будет возвращать список задач по заданному статусу (активные или закрытые).
*/
func (p *Project) FilterTasksByStatus(status Status) []Task {
	l := []Task{}

	for _, id := range p.TasksId {
		task := p.Tasks[id]
		if task.TaskStatus == status {
			l = append(l, task)
		}
	}

	return l
}
