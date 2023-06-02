package serverhandler

import (
	"groupie_tracker/docs/datagather"
	"net/http"
	"os"
	"text/template"
)

type StatusData struct {
	StatusCode int
	StatusMsg  string
}

var (
	temp *template.Template
	D    = StatusData{}
)

func ErrorHandler(w http.ResponseWriter, D *StatusData) {
	w.WriteHeader(D.StatusCode)
	temp.ExecuteTemplate(w, "err.html", D)

	// Reset status code to enable reloading / return to homepage
	D.StatusCode = 0
}

// Future versions will need some kind of front-end logig to monitor
// incorrect ID calls, as all elements are pre-loaded on the front-end, running
// only HTML / CSS.

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// If an error has already been encountered / written in main.go before handlers are called
	if D.StatusCode != 0 {
		ErrorHandler(w, &D)
		return
	}
	// Initialise templates etc.
	g := datagather.NewData
	temp = template.Must(template.ParseGlob("docs/static/*.html"))

	// Error / status handling
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		// Check file integrity
		_, err1 := os.Stat("docs/datagather/datagather.go")
		_, err2 := os.Stat("docs/datagather/cleaner.go")
		_, err3 := os.Stat("docs/static/err.html")
		_, err4 := os.Stat("docs/static/errstylesheet.css")
		_, err5 := os.Stat("docs/static/index.html")
		_, err6 := os.Stat("docs/static/stylesheet.css")
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
			// Basic error handling in the event of incorrect file paths etc.
			D.StatusCode = 500
			D.StatusMsg = "Internal Server Error"
			ErrorHandler(w, &D)
			return
		}

		// Execute templates
		temp.ExecuteTemplate(w, "index.html", g)

		// Basic error handling in the case of incorrect URL
	} else if r.URL.Path != "/" {
		D.StatusCode = 404
		D.StatusMsg = "Page Not Found"
		ErrorHandler(w, &D)
		return
	}
}
