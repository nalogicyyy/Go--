package main

import "fmt"

// 联系人结构体
type Contact struct {
	id     int
	name   string
	phone  string
	group  string
	score  int
}

// 全局存储
var contactList []Contact
var groupMap map[string][]Contact

// 添加联系人，返回提示信息
func addContact(c Contact) (bool, string) {
	// 查重
	for _, item := range contactList {
		if item.id == c.id {
			return false, "学号重复，添加失败"
		}
	}
	contactList = append(contactList, c)
	// 同步分组map
	groupMap[c.group] = append(groupMap[c.group], c)
	return true, "添加成功"
}

// 根据ID删除
func delContact(id int) bool {
	for idx, item := range contactList {
		if item.id == id {
			contactList = append(contactList[:idx], contactList[idx+1:]...)
			return true
		}
	}
	return false
}

// 修改联系方式
func updatePhone(id int, newPhone string) bool {
	for i := range contactList {
		if contactList[i].id == id {
			contactList[i].phone = newPhone
			return true
		}
	}
	return false
}

// 分数筛选
func filterByScore(limit int) []Contact {
	var res []Contact
	for _, item := range contactList {
		if item.score >= limit {
			res = append(res, item)
		}
	}
	return res
}

// 遍历打印
func showAll(list []Contact) {
	if len(list) == 0 {
		fmt.Println("暂无数据")
		return
	}
	for _, v := range list {
		fmt.Printf("编号:%d 姓名:%s 电话:%s 分组:%s 评分:%d\n",
			v.id, v.name, v.phone, v.group, v.score)
	}
}

func main() {
	groupMap = make(map[string][]Contact)
	var opt int

	for {
		fmt.Println("\n=====通讯录管理系统=====")
		fmt.Println("1.新增联系人 2.删除联系人 3.修改电话")
		fmt.Println("4.全部查看 5.分数筛选 6.分组查询 0.退出")
		fmt.Print("请输入操作选项：")
		fmt.Scan(&opt)

		switch opt {
		case 1:
			var c Contact
			fmt.Print("输入编号、姓名、电话、分组、评分：")
			fmt.Scan(&c.id, &c.name, &c.phone, &c.group, &c.score)
			_, msg := addContact(c)
			fmt.Println(msg)
		case 2:
			var delId int
			fmt.Print("输入删除编号：")
			fmt.Scan(&delId)
			if delContact(delId) {
				fmt.Println("删除完成")
			} else {
				fmt.Println("未找到该联系人")
			}
		case 3:
			var editId int
			var newTel string
			fmt.Print("输入修改编号与新号码：")
			fmt.Scan(&editId, &newTel)
			if updatePhone(editId, newTel) {
				fmt.Println("修改成功")
			} else {
				fmt.Println("编号不存在")
			}
		case 4:
			showAll(contactList)
		case 5:
			var scoreLine int
			fmt.Print("输入最低筛选分数：")
			fmt.Scan(&scoreLine)
			filterData := filterByScore(scoreLine)
			showAll(filterData)
		case 6:
			var gName string
			fmt.Print("查询分组名称：")
			fmt.Scan(&gName)
			showAll(groupMap[gName])
		case 0:
			fmt.Println("程序退出")
			return
		default:
			fmt.Println("输入选项无效")
		}
	}
}