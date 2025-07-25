package imooc

type MyLesson struct {
	Name      string
	Id        int
	Testvalue string
}
type Mytest struct {
	Name string
}

func (l *MyLesson) LessonName() string {
	return l.Name
}
func (l *MyLesson) testLessonName() string {
	return l.Name
}
