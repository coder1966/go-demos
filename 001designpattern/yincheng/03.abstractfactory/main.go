package main

import "godemos/001designpattern/yincheng/03.abstractfactory/abstractfactory"

func main() {
	// 创造这个接口
	var fac abstractfactory.DAOFactory

	// 实例化---创建一个mysql的工厂
	fac = &abstractfactory.MySQLFactory{}

	// 调用了
	fac.CreatOrderDetailDAO().SaveOrderDetail()
	fac.CreatOrderMainDAO().SaveOrderMain()
}
