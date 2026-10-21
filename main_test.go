package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLLSegments(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/segments/starred":
			w.Write([]byte(`[{"id":1},{"id":2}]`))
		case "/segments/1/local_legend": // athlete 42 is the LL, 10 vs 7 efforts
			w.Write([]byte(`[{"local_legend":{"athlete_id":42,"mayor_effort_count":7},"your_efforts":{"total_efforts":"10"}}]`))
		case "/segments/2/local_legend": // someone else is LL, 9 vs 4 efforts
			w.Write([]byte(`[{"local_legend":{"athlete_id":99,"mayor_effort_count":9},"your_efforts":{"total_efforts":"4"}}]`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	stravaAPI = srv.URL

	got := getLLSegments("token", 42)
	want := []SegmentData{{1, true, 3}, {2, false, 5}}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("segment %d: got %v, want %v", i, got[i], want[i])
		}
	}
}
