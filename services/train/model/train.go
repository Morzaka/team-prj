package train

import "fmt"

//Train is Model
type Train struct {
	ID    int
	Route int
}

//Add is method
func (t Train) Add(id int, route int) {
	t.ID = id
	t.Route = route
}

//Print is method
func (t Train) Print() {
	fmt.Println(`Train Id:`, t.ID, `Route Id:`, t.Route)
}
