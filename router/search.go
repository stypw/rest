package router

import (
	"errors"
	"net/http"
	JSON "rest/json"
)

func (r *Router) page(w http.ResponseWriter, req *http.Request) (int, JSON.Value, error) {

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

	page := 0
	size := 10
	p, y := body["page"]
	if y {
		po, y := p.(*JSON.Number_T)
		if y {
			page = int(po.Value)
		}
	}

	s, y := body["size"]
	if y {
		sz, y := s.(*JSON.Number_T)
		if y {
			size = int(sz.Value)
		}
	}

	if size < 1 {
		size = 1
	}
	if page < 0 {
		page = 0
	}

	array, err := r.orm.Page(whereObj, oo, page, size)
	if err != nil {
		return 1, JSON.Null(), err
	}
	return 0, array, nil

}
