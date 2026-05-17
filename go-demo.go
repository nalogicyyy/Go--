package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Todo 待办事项结构体
type Todo struct {
	ID    int
	Task  string
	Done  bool
}

var todos []Todo
var nextID = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("=== 简易待办清单 ===")
	fmt.Println("支持命令：add 任务内容 / list / done 任务ID / delete 任务ID / exit")

	for {
		fmt.Print("\n请输入命令：")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "add":
			if len(parts) < 2 {
				fmt.Println("格式错误：add 任务内容")
				continue
			}
			task := strings.Join(parts[1:], " ")
			todos = append(todos, Todo{ID: nextID, Task: task, Done: false})
			fmt.Printf("已添加任务：%s（ID：%d）\n", task, nextID)
			nextID++
		case "list":
			if len(todos) == 0 {
				fmt.Println("暂无待办事项")
				continue
			}
			fmt.Println("ID\t状态\t任务内容")
			for _, t := range todos {
				status := "未完成"
				if t.Done {
					status = "已完成"
				}
				fmt.Printf("%d\t%s\t%s\n", t.ID, status, t.Task)
			}
		case "done":
			if len(parts) != 2 {
				fmt.Println("格式错误：done 任务ID")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("ID必须是数字")
				continue
			}
			found := false
			for i := range todos {
				if todos[i].ID == id {
					todos[i].Done = true
					fmt.Printf("已标记任务ID %d 为完成\n", id)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("未找到该任务ID")
			}
		case "delete":
			if len(parts) != 2 {
				fmt.Println("格式错误：delete 任务ID")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("ID必须是数字")
				continue
			}
			newTodos := []Todo{}
			found := false
			for _, t := range todos {
				if t.ID != id {
					newTodos = append(newTodos, t)
				} else {
					found = true
				}
			}
			if found {
				todos = newTodos
				fmt.Printf("已删除任务ID %d\n", id)
			} else {
				fmt.Println("未找到该任务ID")
			}
		case "exit":
			fmt.Println("程序退出")
			return
		default:
			fmt.Println("不支持的命令，请输入 add/list/done/delete/exit")
		}
	}
}