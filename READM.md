# ToDo

---
1. **[Для разработчиков](#для-разработчика)**
2. **[Установка Task](#установка-task)**
3. **[Список доступных команд Taskfile](#taskfileyaml)**
---

# **Для разработчика**
```bash
git clone https://github.com/SpectrumDev-tech/todo.git
cd todo
go mod tidy

// Run with Go
go run cmd/api/main.go

// Or Task
task run 
```

---
## Установка Task
* ### **Ubuntu**
```bash
sudo snap install task --classic
```
* ### **Archlinux**
```bash
sudo pacman -S go-task
```
> Но могут понадобится дополнительные пакеты
> 
> Лучше почитать официальную документацию [Archlinux and Taskfile](https://archlinux.org/packages/extra/x86_64/go-task/)

* ### **Windows**
```shell
winget install Task.Task
```
> Так же могут возникнуть проблемы с переменными окружения
> 
> Возможно потребуется вручную добавить бинарник в переменную окружения

---
## Taskfile.yaml
|   Команда    |   Описание   |      Пример      |
|:------------:|---|---|
|       `default`       |  выводит список доступных команд <br/>`task --list`          |   `task`   |
|       `run`       |  запускает проект локально <br/>`go run cmd/api/main.go --config path/to/config.yaml`       |   `task run`|
