package hyper

type Counter struct {
	value int
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Incr() {
	c.value++
}

func (c *Counter) Reset() {
	c.value = 0
}
