package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalCalls = promauto.NewCounter(prometheus.CounterOpts{
	Name: "square_endpoint_total_calls",
	Help: "The total number of times square endpoint was called",
})

func squareHandler(w http.ResponseWriter, r *http.Request) {
	totalCalls.Inc()

	vars := mux.Vars(r)
	num, err := strconv.ParseInt(vars["number"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("%d", num*num)))
}
