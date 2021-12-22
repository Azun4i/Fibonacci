package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gobootstrap/handler"
	"net/http"
	"strconv"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger
	Data   handler.Data
}

func newServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		Data:   handler.NewData(),
	}
	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.handlehello())
	s.router.HandleFunc("/fibonacci", s.handleGetFib())
}

func (s *Server) handlehello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	}
}

func (s *Server) handleGetFib() http.HandlerFunc {
	type treq struct {
		X string `json:"x"`
		Y string `json:"y"`
	}
	type Data struct {
		NumFib  []int64 `json:"x"`
		FibPost []int   `json:"y"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		treq := &treq{}
		if err := json.NewDecoder(r.Body).Decode(treq); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		x, _ := strconv.Atoi(treq.X)
		y, _ := strconv.Atoi(treq.Y)
		req := handler.NewData()
		req, err := s.Data.Fib(x, y)
		//req, err := handler.Fib(x, y)
		tmp := Data{}
		tmp.NumFib = req.NumFib
		tmp.FibPost = req.FibPost
		if err != nil {
			s.error(w, r, http.StatusNoContent, err)
		}
		w.Header().Set("Content-Type", "application/json")
		s.respond(w, r, http.StatusOK, tmp)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {

	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
