package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	r := mux.NewRouter()

	// GET\\s+/.*TMS.*\\?terminalID=.*
	sub_paramdownload_generic := r.PathPrefix("/TMS").Subrouter()
	sub_paramdownload_generic.Path("/ParamDownload.aspx").Methods("GET").HandlerFunc(TMSHandler_valuelink).Queries("terminalID",
		"{terminalID:[A-Z]+[0-9]+}")

	// GET /TMS/ParamDownload.aspx?terminalID=FB2088853701&type=partial|full|keys
	sub_paramdownload_type := r.PathPrefix("/TMS").Subrouter()
	sub_paramdownload_type.Path("/ParamDownload.aspx").Methods("GET").HandlerFunc(TMSHandler).Queries("terminalID",
		"{terminalID:[A-Z]+[0-9]+}", "type", "{type:(?:full|partial|keys)}")

	// GET /TMS/ParamDownload.aspx?terminalid=TM2709313801&sequenceNumber=481
	sub_paramdownload_seq := r.PathPrefix("/TMS").Subrouter()
	sub_paramdownload_seq.Path("/ParamDownload.aspx").Methods("GET").HandlerFunc(TMSHandler).Queries("terminalid",
		"{terminalid:[A-Z]+[0-9]+}", "sequenceNumber", "{sequenceNumber:(?:[0-9]+)}")

	// GET /TMS/eventReport.aspx?EOD=GW2364846901&DATE=20201107&TIME=074508&STATUS=0000&encode=1
	sub_eventreport_eod := r.PathPrefix("/TMS").Subrouter()
	sub_eventreport_eod.Path("/eventReport.aspx").Methods("GET").HandlerFunc(TMSHandler).Queries("EOD",
		"{EOD:[A-Z]+[0-9]+}", "DATE", "{DATE:(?:[0-9]+)?}", "TIME", "{TIME:(?:[0-9]+)?}",
		"STATUS", "{STATUS:(?:[0-9]+)?}", "encode", "{encode:[0-9]+}")

	// GET /TMS/eventReport.aspx?Version=0100&TID=FS2085143601&EID=0002&Date=20200914&Time=163017&Details=System
	sub_eventreport_version := r.PathPrefix("/TMS").Subrouter()
	sub_eventreport_version.Path("/eventReport.aspx").Methods("GET").HandlerFunc(TMSHandler).Queries("Version",
		"{Version:[0-9]+}", "TID", "{TID:(?:[A-Z]+[0-9]+)?}",
		"EID", "{EID:(?:[0-9]+)?}", "Date", "{Date:(?:[0-9]+)?}",
		"Time", "{Time:(?:[0-9]+)?}", "Details", "{Details:[a-zA-Z0-9]+}")

	go func() {
		err33913 := http.ListenAndServeTLS(":33913", "conf/test/server.crt", "encrypt/server.key", r)
		if err33913 != nil {
			log.Fatal("TD TMS redirector: port 33913: ", err33913)
		}
		wg.Done()
	}()
	go func() {
		err33914 := http.ListenAndServeTLS(":33914", "conf/test/server.crt", "encrypt/server.key", r)
		if err33914 != nil {
			log.Fatal("TD TMS redirector: port 33914: ", err33914)
		}
		wg.Done()
	}()
	go func() {
		err34913 := http.ListenAndServeTLS(":34913", "conf/test/server.crt", "encrypt/server.key", r)
		if err34913 != nil {
			log.Fatal("TD TMS redirector: port 34913: ", err34913)
		}
		wg.Done()
	}()
	go func() {
		err43913 := http.ListenAndServeTLS(":43913", "conf/test/server.crt", "encrypt/server.key", r)
		if err43913 != nil {
			log.Fatal("TD TMS redirector: port 43913: ", err43913)
		}
		wg.Done()
	}()
	go func() {
		err33971 := http.ListenAndServeTLS(":33971", "conf/test/server.crt", "encrypt/server.key", r)
		if err33971 != nil {
			log.Fatal("TD TMS redirector: port 33971: ", err33971)
		}
		wg.Done()
	}()
	go func() {
		err33972 := http.ListenAndServeTLS(":33972", "conf/test/server.crt", "encrypt/server.key", r)
		if err33972 != nil {
			log.Fatal("TD TMS redirector: port 33972: ", err33972)
		}
		wg.Done()
	}()
	go func() {
		err61914 := http.ListenAndServeTLS(":61914", "conf/test/server.crt", "encrypt/server.key", r)
		if err61914 != nil {
			log.Fatal("Valuelink redirector: port 61914: ", err61914)
		}
		wg.Done()
	}()
	wg.Wait()
}

func TMSHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://127.0.0.1:11443", http.StatusMovedPermanently)
}

func TMSHandler_valuelink(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://127.0.0.1:11553", http.StatusMovedPermanently)
}
