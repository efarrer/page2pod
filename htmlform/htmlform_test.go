package htmlform_test

import (
	"strings"
	"testing"

	"github.com/efarrer/page2pod/htmlform"
	"github.com/stretchr/testify/require"
)

func TestParse_ParsesValidData(t *testing.T) {
	payload := "username=one&podcast=emily&password=&title=mapping&content=one+two+three+four+five&add=Submit"
	expectedFormData := &htmlform.Request{
		Title:   "mapping",
		Podcast: "emily",
		Content: "one two three four five",
		Credentials: htmlform.Credentials{
			Username: "one",
			Password: "",
		},
	}

	res, err := htmlform.Parse(payload)
	require.NoError(t, err)
	require.Equal(t, expectedFormData, res)
}

func TestParse_ReturnsErrorForBadData(t *testing.T) {

	// Generate a 10 MB string which is too large and will fail
	var sb strings.Builder
	for i := 0; i != 1024; i++ {
		for j := 0; j != 1024; j++ {
			sb.WriteString("01234567890")
		}
	}

	_, err := htmlform.Parse(sb.String())
	require.Error(t, err)
}
