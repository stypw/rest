package tl

import (
	"crypto/sha512"
	"encoding/base64"
)

func ToSha512String(input string) string {
	if input == "" {
		return ""
	}
	var e = sha512.Sum512_256([]byte(input))
	return base64.URLEncoding.EncodeToString(e[:])
}

// type httpRequest http.Request

// func (req *httpRequest) Query(k string) string {
// 	matchString := `(?:\A|&)\s*` + k + `\s*=\s*([^&]+)\s*(?:&|\z)`
// 	matchs := regexp.MustCompile(matchString).FindStringSubmatch(req.URL.RawQuery)
// 	if len(matchs) > 1 {
// 		return matchs[1]
// 	}
// 	return ""
// }

// func createSha512(w http.ResponseWriter, req *http.Request) {
// 	r := (*httpRequest)(req)
// 	q := r.Query("sha512")
// 	if q == "" {
// 		w.Write([]byte("error:queryString{sha512} can not empty"))
// 		return
// 	}

// 	w.Write([]byte(ToSha512String(q)))
// }
