package framework

type Product interface {
	Use(s string)
	createClone() Product
}

type Manager struct {
	showcase map[string]Product
}

func NewManager() *Manager {
	return &Manager{showcase: map[string]Product{}}
}

func (m Manager) Register(name string, proto Product) {
	m.showcase[name] = proto
}

func (m Manager) Create(protoname string) Product {
	var p Product = m.showcase[protoname]
	return p.createClone()
}
