package inmemory

type MenuInfoByName struct {
	Price int
	Id    int
}

var MenuIdNameMap = map[int]string{
	0: "",
}

var ListMenuInmemory = map[string]*MenuInfoByName{
	"": {},
}
