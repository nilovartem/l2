package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/nilovartem/l2/develop/dev11/cache"
	"github.com/nilovartem/l2/develop/dev11/event"
)

type Server struct {
	mux   *http.ServeMux
	cache *cache.Cache
}

func New() *Server {
	var server Server
	server.cache = &cache.Cache{
		Events: make(map[int][]event.Event),
		Mutex:  sync.RWMutex{},
	}

	server.mux = http.NewServeMux()

	server.mux.HandleFunc("/create_event", server.middleware(server.createEventHandler))
	server.mux.HandleFunc("/update_event", server.middleware(server.updateEventHandler))
	server.mux.HandleFunc("/delete_event", server.middleware(server.deleteEventHandler))

	server.mux.HandleFunc("/events_for_day", server.middleware(server.eventsForDayHandler))
	server.mux.HandleFunc("/events_for_week", server.middleware(server.eventsForWeekHandler))
	server.mux.HandleFunc("/events_for_month", server.middleware(server.eventsForMonthHandler))

	return &server
}

func (server *Server) Run() error {
	return http.ListenAndServe(":8080", server.mux)
}

func (server *Server) middleware(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request:\nMethod: %s\nURI: %s\nBody: %s\n", r.Method, r.RequestURI, r.Body)
		hf(w, r)
		log.Printf("Response: %v\n\n", w)
	}
}

func (server *Server) createEventHandler(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	if err := e.Decode(r.Body); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}

	if err := server.cache.CreateEvent(&e); err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Event is created", []event.Event{e}, http.StatusOK)
}

func (server *Server) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	if err := e.Decode(r.Body); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	if err := server.cache.UpdateEvent(&e); err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Event is updated", []event.Event{e}, http.StatusOK)
}

func (server *Server) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	if err := e.Decode(r.Body); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	deletedEvent, err := server.cache.DeleteEvent(e.UserID, e.EventID)
	if err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Event has been deleted", []event.Event{*deletedEvent}, http.StatusOK)
}

const layout = "2006-01-02"

func (server *Server) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	date, err := time.Parse(layout, r.URL.Query().Get("date"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	events, err := server.cache.GetEventsForDay(userID, date)
	if err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Events foud", events, http.StatusOK)
}

func (server *Server) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	date, err := time.Parse(layout, r.URL.Query().Get("date"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	events, err := server.cache.GetEventsForWeek(userID, date)
	if err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Events foud", events, http.StatusOK)
}

func (server *Server) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	date, err := time.Parse(layout, r.URL.Query().Get("date"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	var events []event.Event
	if events, err = server.cache.GetEventsForMonth(userID, date); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	respondWithResult(w, "Events foud", events, http.StatusOK)
}

func respondWithError(w http.ResponseWriter, e error, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{e.Error()}

	respond(errorResponse, w, status)

}

func respondWithResult(w http.ResponseWriter, r string, e []event.Event, status int) {
	resultResponse := struct {
		Result string        `json:"result"`
		Events []event.Event `json:"events"`
	}{r, e}

	respond(resultResponse, w, status)
}

func respond(response interface{}, w http.ResponseWriter, status int) {
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
