package wlclient

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLists(t *testing.T) {
	cl, err := New("client-id", "token")
	if err != nil {
		t.Fatal(err)
	}
	s := testServer()
	cl.Endpoint = s.URL
	lists, err := cl.Lists()
	if err != nil {
		t.Fatal(err)
	}
	if v, ex := len(lists), 1; v != ex {
		t.Fatalf("expected %d lists, found %d", ex, v)
	}
	if v, ex := lists[0].Title, "List 1"; ex != v {
		t.Errorf("expected first list to be %q, was %q", ex, v)
	}
}

func testServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(testServerHandler))
	return s
}

func testServerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("method=%s url=%s", r.Method, r.URL)
	switch r.Method + ":" + r.URL.String() {
	case "GET:/lists":
		serveJSON(w, []*List{{ID: 1, Title: "List 1"}})
	default:
		http.NotFound(w, r)
	}
}

func serveJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
