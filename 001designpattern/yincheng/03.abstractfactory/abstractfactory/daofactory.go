package abstractfactory

// 订单 main
// 订单报表 detail

// 根据实际类型，才能生成工厂对象。多了一层抽象。

type OrderMainDAO interface { // 订单 main
	SaveOrderMain() // 保存
	// DeleteOrderMain() // 删除
	// SearchOrderMain() // 搜索
}

type OrderDetailDAO interface { // 订单报表 detail
	SaveOrderDetail() // 保存
	// DeleteOrderDetail() // 删除
	// SearchOrderDetail() // 搜索
}

// 抽象工厂 接口 ###### 完全抽象的接口
type DAOFactory interface {
	CreatOrderMainDAO() OrderMainDAO
	CreatOrderDetailDAO() OrderDetailDAO
}
