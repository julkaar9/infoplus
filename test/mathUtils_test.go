package infoplus_tests

import (
	"infoplus/server/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHumanReadable(t *testing.T) {
	humanReadableTests := []struct {
		num  float64
		want string
	}{
		{12, "12"},
		{1500, "1500"},
		{15000, "15K"},
		{15220, "15K"},
		{15551, "16K"},
		{15551, "16K"},
		{1000000, "1M"},
		{1500000000, "2B"},
	}

	for _, tt := range humanReadableTests {
		got := utils.HumanReadable(tt.num)
		assert.Equal(t, got, tt.want)
	}
}
