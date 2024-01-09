/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description O algoritmo resolve problema de concorrência em inserção de dados em um mapa
 */

package channels

import "time"

type Agent struct {
	Name string
}

func NewAgent(name string) *Agent {
	return &Agent{
		Name: name,
	}
}

type Worker struct {
	ch chan Agent
	ag *Agent
}

func NewWorker(ch chan Agent, ag *Agent) *Worker {
	return &Worker{
		ch: ch,
		ag: ag,
	}
}

func (w *Worker) Run(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond * 100)
		w.ch <- *w.ag
	}
}

func Execute() {
	ch1 := make(chan Agent)
	ch2 := make(chan Agent)

	tiago := NewAgent("Tiago")
	marcos := NewAgent("Marcos")

	wk1 := NewWorker(ch1, tiago)
	wk2 := NewWorker(ch2, marcos)

	tasksMax := 20

	go wk1.Run(5)
	go wk2.Run(20)

	for i := 0; i < tasksMax; i++ {
		select {
		case agent := <-ch1:
			println("Executed by Tiago: ", agent.Name)
		case agent := <-ch2:
			println("Executed by Marcos: ", agent.Name)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}

	close(ch1)
	close(ch2)
}
