package entity

type Cat struct {
}

func (p *Cat) Call(name string) string {
	return "喵喵"
}
