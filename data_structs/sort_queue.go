package data_structs

type Queue struct {
	queue  []interface{}
	length int
}

func (q *Queue) Put(elem interface{}) {
	q.queue = append(q.queue, elem)
	q.length++
}

func (q *Queue) Get() interface{} {
	if 0 == q.length {
		return nil
	}

	elem := q.queue[0]
	q.queue = q.queue[1:]
	q.length--

	return elem
}

func (q *Queue) IsEmpty() bool {
	return 0 == q.length
}
