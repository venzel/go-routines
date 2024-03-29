/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description O algoritmo resolve problema de concorrência em inserção de dados em um mapa
 */

package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Agent struct {
	Name string
}

type Task struct {
	Mp map[int]*Agent
	Mt sync.Mutex
}

func NewTask() *Task {
	return &Task{
		Mp: make(map[int]*Agent),
	}
}

func NewAgent(name string) *Agent {
	return &Agent{Name: name}
}

func (tsk *Task) write(index int, agent *Agent, wg *sync.WaitGroup) {
	tsk.Mt.Lock()
	tsk.Mp[index] = agent
	tsk.Mt.Unlock()
	wg.Done()
}

func (tsk *Task) worker(agent *Agent, ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		tsk.write(i, agent, wg)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(3)))
	}

	close(ch)
}

func (tsk *Task) publish(ch chan int, tasks int) {
	for i := 0; i < tasks; i++ {
		ch <- i
	}
}

func (tsk *Task) ranking(tasks int, agents int) {
	mp := map[*Agent]int{}

	for i := 0; i < tasks; i++ {
		fmt.Println("task", i, "->", tsk.Mp[i].Name)
		mp[tsk.Mp[i]]++
	}

	var count int = 0
	var agent Agent

	fmt.Println("--------------------------------")

	for k, v := range mp {
		fmt.Println(k.Name, "->", v)

		if v > count {
			count = v
			agent = *k
		}
	}

	fmt.Println("--------------------------------")

	fmt.Println("O infitete foi:", agent.Name)
}

func Execute() {
	japa := NewAgent("Japa")
	erickson := NewAgent("Erickson")
	joab := NewAgent("Joab")
	marcos := NewAgent("Marcos")

	tasks := 10
	agents := []*Agent{japa, erickson, joab, marcos}

	tsk := NewTask()

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(tasks)

	for i := range agents {
		go tsk.worker(agents[i], ch, wg)
	}

	tsk.publish(ch, tasks)

	wg.Wait()

	tsk.ranking(tasks, len(agents))
}
