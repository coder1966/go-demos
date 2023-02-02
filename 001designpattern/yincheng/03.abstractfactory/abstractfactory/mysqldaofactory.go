package abstractfactory

import "fmt"

// mysql 针对两个接口的实现

type MySQLMainDAO struct{}

func (d *MySQLMainDAO) SaveOrderMain() {
	fmt.Println("RUN MySQL SaveOrderMain()")
}

type MySQLDetailDAO struct{}

func (d *MySQLDetailDAO) SaveOrderDetail() {
	fmt.Println("RUN MySQL SaveOrderDetail()")
}

// type OrderDetailDAO interface { // 订单报表 detail
// 	SaveOrderDetail() // 保存
// 	// DeleteOrderDetail() // 删除
// 	// SearchOrderDetail() // 搜索
// }

// // 抽象工厂 接口 ###### 完全抽象的接口
// type DAOFactory interface {
// 	CreatOrderMainDAO() OrderMainDAO
// 	CreatOrderDetailDAO() OrderDetailDAO
// }

// // 操作的抽象
// type PlusOperatorFactory struct{}

// // 操作类，依赖数据。所以引入数据
// type PlusOperator struct {
// 	*OperatorBase
// }

// // 对刚才数据做实际的操作
// func (o *PlusOperator) Result() int {
// 	return o.left + o.right
// }

// func (PlusOperatorFactory) Creat() Operator {
// 	return &PlusOperator{OperatorBase: &OperatorBase{}}
// }
