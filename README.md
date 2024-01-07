# Go routines

Utilizando go routines juntamente com mutex do pacote sync, para resolver o problema de concorrência no processo de inserção no objeto do tipo map, entre agentes paralelos.

Como funciona?

Temos 9 tarefas e 4 trabalhadores em **paralelo**, quem realizar mais tarefas, vence.

```bash
task 0 -> Marcos
task 1 -> Japa
task 2 -> Erickson
task 3 -> Joab
task 4 -> Joab
task 5 -> Japa
task 6 -> Marcos
task 7 -> Erickson
task 8 -> Marcos
--------------------------------
Erickson -> 2
Joab -> 2
Marcos -> 3
Japa -> 2
--------------------------------
O infitete foi: Marcos
```

<div>
  <img align="left" src="https://imgur.com/k8HFd0F.png" width=35 alt="Profile"/>
  <sub>Made with 💙 by <a href="https://github.com/venzel">Enéas Almeida</a></sub>
</div>
