package inmemory

type MenuInfoById struct {
	Price int
	Name  string
}

var ListMenuInmemory = map[int]*MenuInfoById{
	0: {},
}
