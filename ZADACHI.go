package main

import (
	"fmt"
	"time"
)

type CustomWaitGroup struct {
	semaphore chan struct{} // подсчет количества задач
	done      chan struct{} // уведомление о завершении задач
}

func (wg *CustomWaitGroup) Add() {
	wg.semaphore <- struct{}{} // добавление задачи в канал
}

func (wg *CustomWaitGroup) Done() {
	<-wg.semaphore // извлечение задачи из канала
	if len(wg.semaphore) == 0 {
		close(wg.done) // сигнал о пустом канале
	}
}

func (wg *CustomWaitGroup) Wait() {
	<-wg.done
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		semaphore: make(chan struct{}, 100),
		done:      make(chan struct{}),
	}
}

func main() {
	wg := NewCustomWaitGroup()

	for i := 0; i < 5; i++ {
		wg.Add()

		go func(i int) {
			defer wg.Done()
			fmt.Printf("Task %d started\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Task %d finished\n", i)
		}(i)
	}

	wg.Wait() // Ждём завершения всех задач
	fmt.Println("All tasks completed")
}
