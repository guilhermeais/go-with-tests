package clockface_test

import (
	"bytes"
	"encoding/xml"
	"go-with-tests/clockface"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			xmlBuffer := bytes.Buffer{}
			clockface.SVGWriter(&xmlBuffer, c.time)

			svg := SVG{}
			xml.Unmarshal(xmlBuffer.Bytes(), &svg)

			want := c.line

			if !containsLine(svg, want) {
				t.Errorf("Expect to find the second hand with x2 of %+v and y2 of %+v, in the SVG output %v", want.X2, want.Y2, xmlBuffer.String())
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			xmlBuffer := bytes.Buffer{}
			clockface.SVGWriter(&xmlBuffer, c.time)

			svg := SVG{}
			xml.Unmarshal(xmlBuffer.Bytes(), &svg)

			want := c.line

			if !containsLine(svg, want) {
				t.Errorf("Expect to find the minute hand line %+v, in the SVG lines %+v", want, svg.Line)
			}
		})
	}

}

func containsLine(svg SVG, want Line) bool {
	for _, line := range svg.Line {
		if line == want {
			return true
		}
	}
	return false
}

func simpleTime(hour, minute, seconds int) time.Time {
	return time.Date(312, time.October, 28, hour, minute, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
