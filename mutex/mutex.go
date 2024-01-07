/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description O algoritmo resolve problema de concorrência em inserção de dados em um mapa
 */

package mutex

import (
	"sync"
	"time"
)

type Agent struct {
	Name string
	Age  uint8
}

type Task struct {
	mt sync.Mutex
	mp map[int]*Agent
}

func NewTask() *Task {
	return &Task{
		mp: make(map[int]*Agent),
	}
}

func (t *Task) Write(index int, agent *Agent, wg *sync.WaitGroup) {
	t.mt.Lock()
	t.mp[index] = agent
	t.mt.Unlock()
	wg.Done()
}

func (t *Task) Read(key int) *Agent {
	t.mt.Lock()
	defer t.mt.Unlock()
	return t.mp[key]
}

func (t *Task) Add(agent *Agent, positions int, wg *sync.WaitGroup) {
	wg.Add(positions)
	for i := 0; i < positions; i++ {
		time.Sleep(time.Millisecond.Abs() * 2)
		go t.Write(i, agent, wg)
	}
	wg.Done()
}

func (t *Task) Execute() {
	task := NewTask()

	wg := sync.WaitGroup{}
	wg.Add(2)

	tiago := &Agent{Name: "Erickson", Age: 30}
	japa := &Agent{Name: "Japa", Age: 30}

	positions := 10

	go task.Add(tiago, positions, &wg)
	go task.Add(japa, positions, &wg)

	wg.Wait()

	for i := 0; i < positions; i++ {
		println(i, task.Read(i).Name)
	}

	countJapa := 0
	countTiago := 0

	for _, v := range task.mp {
		if v.Name == "Japa" {
			countJapa++
		} else {
			countTiago++
		}
	}

	println("--------------------------------")
	println("Japa:", countJapa)
	println("Erickson:", countTiago)
}
