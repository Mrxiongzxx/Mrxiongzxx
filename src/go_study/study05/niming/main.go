package niming

import "fmt"


	//可以像字段成员那样访问匿名字段⽅法，编译器负责查找。
	type User struct {
		id int
		name string
	}
	type Manager struct {
		User
	}
	func (self *User) ToString() string { // receiver = &(Manager.User)
		return fmt.Sprintf("User: %p, %v", self, self)
	}
	func main() {
		m := Manager{User{1, "Tom"}}
		fmt.Printf("Manager: %p\n", &m)
		fmt.Println(m.ToString())  //Manager 会找到 User的 ToString方法

		//输出：
		//Manager: 0x2102281b0
		//User : 0x2102281b0, &{1 Tom}
	}

