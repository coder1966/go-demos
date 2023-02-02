package abstractfactory

// 抽象的工厂
type MySQLFactory struct{}

func (f *MySQLFactory) CreatOrderMainDAO() OrderMainDAO {
	return &MySQLMainDAO{}
}
func (f *MySQLFactory) CreatOrderDetailDAO() OrderDetailDAO {
	return &MySQLDetailDAO{}
}
