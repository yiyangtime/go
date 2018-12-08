package main

//user 在程序里定义一个用户类型
//type user struct {
//	name  string
//	email string
//}

// notify 使用值接收者实现了一个方法
//func (u user) notify() {
//	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
//}

// changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

// main 是应用程序的入口
func main() {
	// user类型的值可以用来调用 使用值接收者声明的方法
	bill := user{"Bill", "Bill@email.com"}
	bill.notify()

	// 指向user类型值的指针也可以用来调用 使用值接收者声明的方法
	lisa := &user{"lisa", "lisa@emain.com"}
	lisa.notify()

	// user类型的值可以用来调用 使用指针接收者声明的方法
	bill.changeEmail("Bill@newDomain.com")
	bill.notify()

	// 指向user类型值的指针可以用来调用 使用指针接收者声明的方法
	lisa.changeEmail("lisa@newDomain.com")
	lisa.notify()
}
