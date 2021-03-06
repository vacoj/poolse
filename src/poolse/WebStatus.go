package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func statusWeb(rw http.ResponseWriter, req *http.Request) {
	StatusMu.Lock()
	defer StatusMu.Unlock()

	result := false
	var blob []byte

	req.ParseForm()
	ppid := req.Form.Get("id")
	id, err := strconv.Atoi(ppid)

	if err == nil && id >= 0 && id < len(STATUS.Targets) && len(STATUS.Targets) > 0 {
		blob, _ = json.Marshal(&STATUS.Targets[id])
		result = STATUS.Targets[id].OK
	} else {
		blob, _ = json.Marshal(&STATUS)
		result = STATUS.isOk()
	}

	if SETTINGS.Service.ShowHTTPLog {
		logRequest(req, result)
	}

	io.WriteString(rw, string(blob))
}

func statusSimpleWeb(rw http.ResponseWriter, req *http.Request) {
	result := false

	req.ParseForm()
	ppid := req.Form.Get("id")
	id, err := strconv.Atoi(ppid)

	if err == nil && id >= 0 && id < len(STATUS.Targets) && len(STATUS.Targets) > 0 {
		result = STATUS.checkStatusByID(id)
	} else {
		result = STATUS.checkStatus()
	}

	if SETTINGS.Service.ShowHTTPLog {
		logRequest(req, result)
	}

	if result {
		rw.WriteHeader(http.StatusOK)
	} else {
		http.Error(rw, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	}
}

func statusSimple2Web(rw http.ResponseWriter, req *http.Request) {
	result := false

	req.ParseForm()
	ppid := req.Form.Get("id")
	id, err := strconv.Atoi(ppid)

	if err == nil && id >= 0 && id < len(STATUS.Targets) && len(STATUS.Targets) > 0 {
		result = STATUS.checkStatusByID(id)
	} else {
		result = STATUS.checkStatus()
	}

	if SETTINGS.Service.ShowHTTPLog {
		logRequest(req, result)
	}

	if result {
		rw.WriteHeader(http.StatusOK)
	} else {
		var err = errors.New("intentionally erroring for sake of not returning anything to the caller")
		panic(err)
	}
}
