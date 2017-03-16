package main

import (
	"net/http"
)

func startWeb() {
	routes := map[string]func(http.ResponseWriter, *http.Request){
		// return status
		"/status":         statusWeb,
		"/status/simple":  statusSimpleWeb,
		"/status/simple2": statusSimple2Web,

		// turn on/off application (for benefit of monitor or LB)
		"/toggle/on":  toggleOnWeb,
		"/toggle/adminon":  toggleAdminOnWeb,
		"/toggle/adminoff":  toggleAdminOffWeb,
		"/toggle/off": toggleOffWeb,
		"/toggle":     toggleWeb,

		"/fakesmoke":    fakeSmoke,
		"/fakehealth":   fakeHealth,
		"/fakeexpected": fakeExpected,

		"/settings":        settingsWeb,
		"/settings/reload": settingsReloadWeb,

		// show landing page?
		"/": statusWeb,
	}

	// register routes
	for k, v := range routes {
		http.HandleFunc(k, v)
	}

	// In case we want static content
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// Host web server
	panic(http.ListenAndServe(":"+SETTINGS.Service.HTTPPort, nil))
}
