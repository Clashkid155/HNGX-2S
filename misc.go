package main

import (
	"net/http"
	"regexp"
)

type Response struct {
	Id      int    `json:"id,omitempty"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func jsonContent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func isDigit(value string) bool {
	/*data, err := strconv.Atoi(value.(string))

	if err != nil {
		log.Println("This is not an int")
	}
	fmt.Println(data)*/
	isNum := regexp.MustCompile(`^\d+$`)

	return isNum.MatchString(value)
}
