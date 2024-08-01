package demo_one

type MyStruct struct {
	name string
}

func (c *MyStruct) Name() string {
	return c.name
}
