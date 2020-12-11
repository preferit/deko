package deko

func NewSpecification() *Specification {
	return &Specification{}
}

type Specification struct{}

// Requirements
func (me *Specification) Requirements() []*Requirement {
	return nil
}

type Requirement struct{}
