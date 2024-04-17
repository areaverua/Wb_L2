package main

import (
	"net/http"

	"github.com/areaverua/wbL2/develop/dev11/handlers"
)

func main() {
	ch := handlers.NewCalendarHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", ch.CreateEventRequest)
	mux.HandleFunc("/update_event", ch.UpdateEventRequest)
	mux.HandleFunc("/delete_event", ch.DeleteEventRequest)
	mux.HandleFunc("/events_for_day", ch.EventsForDayRequest)
	mux.HandleFunc("/events_for_week", ch.EventsForWeekRequest)
	mux.HandleFunc("/events_for_month", ch.EventsForMonthRequest)

	wrappedMux := handlers.NewLogger(mux)

	http.ListenAndServe("localhost:8080", wrappedMux)
}
