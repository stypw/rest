package router

import (
	"errors"
	"net/http"
	JSON "rest/json"
)

func (r *Router) first(w http.ResponseWriter, req *http.Request) (int, JSON.Value, error) {
	body, err := JSON.FromStream(req.Body)
	if err != nil {
		return 1, JSON.Null(), errors.New("where can not be empty")
	}
	where, y := body["where"]
	if !y {
		return 1, JSON.Null(), errors.New("where can not be empty")
	}

	whereObj, y := where.(JSON.Object)
	if !y {
		return 1, JSON.Null(), errors.New("where can not be empty")
	}
	var oo JSON.Object = nil
	order, y := body["order"]
	if y {
		orderObj, y := order.(JSON.Object)
		if y {
			oo = orderObj
		}
	}

	array, err := r.orm.First(whereObj, oo)
	if err != nil {
		return 1, JSON.Null(), err
	}
	return 0, array, nil
}
