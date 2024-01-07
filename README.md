# Go routines

Utilizando go routines juntamente com mutex do pacote sync, para resolver o problema de concorrência no processo de inserção no objeto do tipo map, entre agentes paralelos.

```bash
❯ go run main.go
0 Erickson
1 Japa
2 Japa
3 Erickson
4 Japa
5 Japa
6 Japa
7 Japa
8 Japa
9 Japa
--------------------------------
Japa: 8
Erickson: 2
```

<div>
  <img align="left" src="https://imgur.com/k8HFd0F.png" width=35 alt="Profile"/>
  <sub>Made with 💙 by <a href="https://github.com/venzel">Enéas Almeida</a></sub>
</div>
