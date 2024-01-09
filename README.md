# Go routines

Utilizando go routines juntamente com mutex do pacote sync, para resolver o problema de concorrência no processo de inserção no objeto do tipo map, entre agentes paralelos.

Como funciona?

Temos 9 tarefas e 4 trabalhadores em **paralelo**, quem realizar mais tarefas, vence.

**Observação**: Como está utilizando o Multex para evitar colisão, se torna mais complexo o desenvolvimento do código, uma alteranativa para evitar a utilização do Multex é utilizar channels com o recurso de select.

```bash
❯ go run main.go
task 0 -> Marcos
task 1 -> Japa
task 2 -> Japa
task 3 -> Japa
task 4 -> Erickson
task 5 -> Joab
task 6 -> Japa
task 7 -> Joab
task 8 -> Marcos
task 9 -> Japa
--------------------------------
Marcos -> 2
Japa -> 5
Erickson -> 1
Joab -> 2
--------------------------------
O infitete foi: Japa
```

## main.go

```go
/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description O algoritmo resolve problema de concorrência em inserção de dados em um mapa
 */

package main

import "routines/channels"

func main() {
	japa := channels.NewAgent("Japa")
	erickson := channels.NewAgent("Erickson")
	joab := channels.NewAgent("Joab")
	marcos := channels.NewAgent("Marcos")

	tasks := 9
	agents := []*channels.Agent{japa, erickson, joab, marcos}

	channels.StartRound(tasks, agents)
}
```

## mutex/worker.go

```go
/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description O algoritmo resolve problema de concorrência em inserção de dados em um mapa
 */

package channels

import (
	"fmt"
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

func StartRound(tasks int, agents []*Agent) {
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
```

<div>
  <img align="left" src="https://imgur.com/k8HFd0F.png" width=35 alt="Profile"/>
  <sub>Made with 💙 by <a href="https://github.com/venzel">Enéas Almeida</a></sub>
</div>
