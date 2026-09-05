package main

// TODO: Define a Queue struct that holds items internally.
type Queue struct{
	items []string
}
// TODO: Implement Push(item string) to add an item to the back.
func (q *Queue) Push(item string) {
	q.items = append(q.items,item)

}

// TODO: Implement Pop() (string, bool) to remove and return the front item.
func (q *Queue) Pop() (string,bool){
	n:= len(q.items)
	if n == 0{
		return "",false
	}
	front := q.items[0]
	q.items = q.items[1:n]
	return front,true
}
// TODO: Implement Len() int to return the current queue size
func (q *Queue) Len() int{
	return len(q.items)
}
// TODO: Implement Peek() (string, bool) to return the front item without removing it.
func (q *Queue) Peek() (string,bool){
	if len(q.items) == 0{
		return "",false
	}
	return q.items[0],true
}
func main() {}
