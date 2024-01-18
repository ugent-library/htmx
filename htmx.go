package htmx

import (
	"net/http"
)

// Request headers

func Request(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func Boosted(r *http.Request) bool {
	return r.Header.Get("HX-Boosted") == "true"
}

func CurrentURL(r *http.Request) string {
	return r.Header.Get("HX-Current-URL")
}

func HistoryRestoreRequest(r *http.Request) bool {
	return r.Header.Get("HX-History-Restore-Request") == "true"
}

func Prompt(r *http.Request) string {
	return r.Header.Get("HX-Prompt")
}

func TargetID(r *http.Request) string {
	return r.Header.Get("HX-Target")
}

func TriggerID(r *http.Request) string {
	return r.Header.Get("HX-Trigger")
}

func TriggerName(r *http.Request) string {
	return r.Header.Get("HX-Trigger-Name")
}

// Response headers

func Location(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Location", v)
}

func PushURL(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Push-Url", v)
}

func Redirect(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Redirect", v)
}

func Refresh(w http.ResponseWriter) {
	w.Header().Set("HX-Refresh", "true")
}

func ReplaceURL(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Replace-Url", v)
}

func Reswap(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Reswap", v)
}

func Retarget(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Reswap", v)
}

func Reselect(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Reswap", v)
}

func Trigger(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Trigger", v)
}

func TriggerAfterSettle(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Trigger-After-Settle", v)
}

func TriggerAfterSwap(w http.ResponseWriter, v string) {
	w.Header().Set("HX-Trigger-After-Swap", v)
}
