package entity

type Dog struct {
}

func (p *Dog) Call(name string) string {
	return "汪汪"
}
