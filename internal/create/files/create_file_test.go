package files

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
    tests := []struct {
        content            string
        path               string
        removeEscapeValues bool
        expectedContent    string
    }{
        {
            content:           "Hello\\World\\",
            path:              "test.txt",
            removeEscapeValues: true,
            expectedContent:   "HelloWorld",
        },
		{
            content:           "Hello World",
            path:              "test2.txt",
            removeEscapeValues: true,
            expectedContent:   "Hello World",
        },
        {
            content:           "Hello World",
            path:              "test3.txt",
            removeEscapeValues: false,
            expectedContent:   "Hello World",
        },
		{
            content:           "Hello\\World\\",
            path:              "test4.txt",
            removeEscapeValues: false,
            expectedContent:   "Hello\\World\\",
        },
    }

    for _, tc := range tests {
		CreateFile(tc.content, tc.path, tc.removeEscapeValues)

		data, err := os.ReadFile(tc.path)
		if err != nil {
			t.Fatal(err)
		}

		if string(data) != tc.expectedContent {
			t.Errorf("Expected content to be '%s', got '%s'", tc.expectedContent, string(data))
		}

		err = os.Remove(tc.path)
		if err != nil {
			t.Fatal(err)
		}
    }
}