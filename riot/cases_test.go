package riot_test

import (
	"net/http"
)

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "6 numbers",
		input:       "123456",
		expected:    true,
	},
	{
		description: "6 character input containing letters",
		input:       "123fws",
		expected:    false,
	},
	{
		description: "7 numbers",
		input:       "1234567",
		expected:    false,
	},
}

var testCaseshttp = []struct {
	description string
	httpWriter  int
	expectedErr error
	testURL     string
	BIN         string
	endpoint    string
	authHeader  string
}{
	{
		description: "Server responds with StatusOK 200",
		httpWriter:  http.StatusAccepted,
		expectedErr: nil,
		testURL:     "http://randomurl.com",
		BIN:         "123456",
		endpoint:    "endpoint",
		authHeader:  "header",
	},
	{
		description: "Malformed URL should return error",
		httpWriter:  http.StatusAccepted,
		expectedErr: nil,
		testURL:     "http://randomurl.c om",
		BIN:         "123456",
		endpoint:    "endpoint",
		authHeader:  "header",
	},
}
