package main

type Visitor interface {
	Visit(VisitorFunc) error
}

type VisitorFunc func() error

type VisitorList []Visitor

func (l VisitorList) Visit(fn VisitorFunc) error {
	return nil
}

type Visitor1 struct {
}

func (v Visitor1) Visit(fn VisitorFunc) error {
	return nil
}

func main() {
	var visitor Visitor
	var visitors []Visitor

	visitor = Visitor1{}
	visitors = append(visitors, visitor)
	_ = VisitorList(visitors)

}
