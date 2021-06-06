package responses

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	serviceErr "github.com/gopher-translator-service/internal/errors"
)

type fakeResponseWriter struct {
	writeAnError  bool
	wroteStatCode int
	wroteBytes    *bytes.Buffer
}

func (r *fakeResponseWriter) Header() http.Header {
	m := make(map[string][]string)
	return m
}
func (r *fakeResponseWriter) Write(d []byte) (int, error) {
	if r.writeAnError {
		err := errors.New("Some fake error")
		return 0, err
	}
	_, err := r.wroteBytes.Write(d)
	if err != nil {
		return 0, err
	}
	return len(d), nil
}
func (r *fakeResponseWriter) WriteHeader(statusCode int) {
	r.wroteStatCode = statusCode
}

func TestJsonResponse(t *testing.T) {
	cases := []struct {
		rw       http.ResponseWriter
		json     string
		expected string
	}{
		{
			&fakeResponseWriter{wroteBytes: bytes.NewBuffer([]byte(""))},
			`{"data":"testing"}`,
			`"{\"data\":\"testing\"}"`,
		},
		{
			&fakeResponseWriter{wroteBytes: bytes.NewBuffer([]byte(""))},
			``,
			`""`,
		},
		{
			&fakeResponseWriter{wroteBytes: bytes.NewBuffer([]byte(""))},
			`{}`,
			`"{}"`,
		},
		{
			&fakeResponseWriter{wroteBytes: bytes.NewBuffer([]byte(""))},
			`{ "data": [ {"a": "1"}, {"d": 1}, {"b": "3uak"}] }`,
			`"{ \"data\": [ {\"a\": \"1\"}, {\"d\": 1}, {\"b\": \"3uak\"}] }"`,
		},
		{
			&fakeResponseWriter{writeAnError: true},
			`never mind this`,
			`never mind this`,
		},
	}

	for _, tc := range cases {
		err := JsonResponse(tc.rw, tc.json)
		frw := tc.rw.(*fakeResponseWriter)
		if (err != nil) && !frw.writeAnError {
			t.Fatalf("Failed to write %s", tc.json)
		}
		if err != nil {
			continue
		}

		got := frw.wroteBytes.String()
		if got != tc.expected {
			t.Fatalf("JsonResponse failed expected %s got %s", tc.expected, got)
		}
	}
}

func TestErrorResponse(t *testing.T) {
	cases := []struct {
		rw       http.ResponseWriter
		err     error
	}{
		{
			&fakeResponseWriter{wroteBytes: bytes.NewBuffer([]byte(""))},
			serviceErr.NewHttpErr(400, "some error", errors.New("some error")),
		},
		{
			&fakeResponseWriter{wroteBytes: bytes.NewBuffer([]byte(""))},
			serviceErr.NewHttpErr(404, "different description", errors.New("some other error")),
		},
		{
			&fakeResponseWriter{wroteBytes: bytes.NewBuffer([]byte(""))},
			errors.New("not http error"),
		},
		{
			&fakeResponseWriter{wroteBytes: bytes.NewBuffer([]byte(""))},
			nil,
		},
	}

	for _, tc := range cases {
		err := ErrorResponse(tc.rw, tc.err, nil)
		frw := tc.rw.(*fakeResponseWriter)
		if err != nil {
			t.Fatalf("Failed to write %v", tc.err)
		}

		got := frw.wroteBytes.String()
		expected, err := json.Marshal(tc.err)
		if err != nil {
			t.Fatalf("Failed to marshal %v", tc.err)
		}

		httpErr, ok := tc.err.(*serviceErr.HttpErr)
		if ok {
			if frw.wroteStatCode != httpErr.Status {
				t.Fatalf("JsonResponse wrote wrong status code expected %d got %d", httpErr.Status, frw.wroteStatCode)
			}
		}

		if got == "" && tc.err == nil {
			continue
		}

		if got != string(expected) {
			t.Fatalf("JsonResponse failed expected %s got %s", tc.err.Error(), string(expected))
		}
	}
}