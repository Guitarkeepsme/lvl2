package cache

import (
	"fmt"
	"log"
	"sync"
	"time"

	"task11/domain"
	"task11/models"
)

// Данные кэшируем в мапу, значения которой --
// ссылка на событие из events
type Cache struct {
	mutex   sync.Mutex
	storage map[string][]*models.Event
}

// Функция для внесения данных в мапу кэша
func NewCache() *Cache {
	return &Cache{
		mutex:   sync.Mutex{},
		storage: map[string][]*models.Event{},
	}
}

// Метод, который создаёт новое событие
func (c *Cache) Create(e *models.Event) {
	// Закрываем мьюекс
	c.mutex.Lock()
	// Перед возвращением функции открываем мьютекс
	defer c.mutex.Unlock()

	c.storage[e.Date] = append(c.storage[e.Date], e)

}

// Метод, который обновляет событие
func (c *Cache) Update(e models.Event, newDate, newTime string) {
	// Проходим циклом по событиям в кэше
	for _, event := range c.storage[e.Date] {
		// Если находим событие с заданными датой и временем,
		// ...
		if event.Time == e.Time && event.Date == e.Date {
			UpdatedEvent := models.Event{}
			UpdatedEvent.UserID = e.UserID

			if newTime != "" {
				UpdatedEvent.Time = newTime
			}

			if newDate != "" {
				UpdatedEvent.Date = newDate
				c.Delete(e.Date, e.Time)
				newEvent := domain.NewEvent(UpdatedEvent.Date, UpdatedEvent.Time, UpdatedEvent.UserID)
				c.Create(newEvent)
			}
			return
		}
	}
}

// Метод для удаления событий
func (c *Cache) Delete(date, time string) {
	// Закрываем мьютекс
	c.mutex.Lock()
	// Перед возвращением функции открываем мьютекс
	defer c.mutex.Unlock()
	// Создаём переменную для очищенного стека событий
	event := []*models.Event{}
	// Проходим циклом по всем событиям
	for i := 0; i < len(c.storage[date]); i += 1 {
		// Если дата и время события не сходятся с данными величинами,
		// это значит, что перед нами не то, что трубется удалить,
		// поэтому добавляем событие в стек
		if c.storage[date][i].Time != time {
			event = append(event, c.storage[date][i])
		}
	}
	// Обновляем кэш
	c.storage[date] = event
}

// Методы для работы с событиями дня, недели и месяца

func (c *Cache) ReadDay(date string) ([]*models.Event, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.storage[date]
	fmt.Println(val)
	if ok {
		return val, true
	}
	log.Println("Ошибка в получении данных", date, ": событий не обнаружено")
	return []*models.Event{}, false
}

func (c *Cache) ReadWeek(date string) ([]models.Event, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := []models.Event{}

	today, errParseToday := time.Parse("2006-01-02", date)

	if errParseToday != nil {
		log.Println("Error at parsing date: ", errParseToday)
		return nil, false
	}

	for _, ev := range c.storage {

		for _, event := range ev {

			evDate, errParseEv := time.Parse("2006-01-02", event.Date)
			if errParseEv != nil {
				log.Println("Error at parsing date: ", errParseEv)
				return nil, false
			}

			if evDate.After(today) && evDate.Before(today.Add(time.Hour*168)) {
				result = append(result, *event)
			}
		}

	}

	return result, true
}

func (c *Cache) ReadMonth(date string) ([]models.Event, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := []models.Event{}

	today, errParseToday := time.Parse("2006-01-02", date)

	if errParseToday != nil {
		log.Println("Error at parsing date: ", errParseToday)
		return nil, false
	}

	for _, ev := range c.storage {

		for _, event := range ev {

			evDate, errParseEv := time.Parse("2006-01-02", event.Date)
			if errParseEv != nil {
				log.Println("Error at parsing date: ", errParseEv)
				return nil, false
			}

			if evDate.After(today) && evDate.Before(today.Add(time.Hour*24*30)) {
				result = append(result, *event)
			}
		}

	}

	return result, true
}
