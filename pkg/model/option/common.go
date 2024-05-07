package optionmodel

type CommonStruct struct {
	ID string
}

func (c CommonStruct) HasID() bool {
	return c.ID == ""
}
