package infoplus_tests

import (
	"infoplus/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackoverflowReputationhandler(t *testing.T) {
	const url = "/api/stackoverflow"
	r, err := GetGinRouter()
	if err != nil {
		t.Fail()
	}

	r.GET(url, server.StackoverflowReputationhandler)

	w, req := GetHttpRecorder(url, map[string]string{"userid": "1111111"})
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	want := `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="134" height="20"
		id="svg1">
		<linearGradient id="smooth" x2="0" y2="100%">
			<stop offset="0" stop-color="#bbb" stop-opacity=".1"></stop>
			<stop offset="1" stop-opacity=".1"></stop>
		</linearGradient>
		<clipPath id="round">
			<rect width="134" height="20" rx="4" fill="#fff"></rect>
		</clipPath>
		<g clip-path="url(#round)">
			<rect width="92" height="20" fill="#555555"></rect>
			<rect x="92" width="42" height="20" fill="#97ca00"></rect>
			<rect width="134" height="20" fill="url(#smooth)"></rect>
		</g>
		<g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
			<text x="47" y="15" fill="#010101" fill-opacity=".3" lengthAdjust="spacing">stackoverflow</text>
			<text x="47" y="14" lengthAdjust="spacing">stackoverflow</text>
			<text x="112" y="15" fill="#010101" fill-opacity=".3" lengthAdjust="spacing">error</text>
			<text x="112" y="14" lengthAdjust="spacing">error</text>
		</g>
	</svg>`
	assert.Equal(t, StandardizeSpaces(want), StandardizeSpaces(w.Body.String()))
}
