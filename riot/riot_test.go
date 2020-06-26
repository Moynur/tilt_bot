package riot_test

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"tiltbot/riot"

	"github.com/stretchr/testify/assert"
)

func TestErrorCodes(t *testing.T) {
	for _, tc := range testCaseshttp {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(tc.httpWriter) // test table
		}))
		defer ts.Close()
		testURL := ts.URL
		_, err := riot.GetSummonerByName(testURL, tc.Name, tc.authHeader)
		if err != tc.expectedErr {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v\n", tc.description, tc.expectedErr, err)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestSearchOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got ‘%s’", r.Method)
		}

	}))
	defer ts.Close()
	URL := ts.URL
	_, err := SearchPost(URL, "123456", "verify", "key")
	if err != nil {
		t.Errorf("Search returned an error: %s", err)
	}
}

func TestNoKey(t *testing.T) {
	var emptyString string
	_, err := riot.GetInfo("1", "2", emptyString)
	expectedErr := errors.New("No Auth Headesr")
	if !assert.Equal(t, err, expectedErr, "Both errors should be the same") {
		log.Fatalf("Test failed \ngot %#v \nexpected %#v", err, expectedErr)
	}

}
