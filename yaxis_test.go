package chart

import (
	"testing"

	"github.com/Strubbl/go-chart/v2/testutil"
)

func TestYAxisGetTicks(t *testing.T) {
	// replaced new assertions helper

	r, err := PNG(1024, 1024)
	testutil.AssertNil(t, err)

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

	ya := YAxis{}
	yr := &ContinuousRange{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	ticks := ya.GetTicks(r, yr, styleDefaults, vf)
	testutil.AssertLen(t, ticks, 32)
}

func TestYAxisGetTicksWithUserDefaults(t *testing.T) {
	// replaced new assertions helper

	r, err := PNG(1024, 1024)
	testutil.AssertNil(t, err)

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

	ya := YAxis{
		Ticks: []Tick{{Value: 1.0, Label: "1.0"}},
	}
	yr := &ContinuousRange{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	ticks := ya.GetTicks(r, yr, styleDefaults, vf)
	testutil.AssertLen(t, ticks, 1)
}

func TestYAxisMeasure(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)
	style := Style{
		Font:     f,
		FontSize: 10.0,
	}
	r, err := PNG(100, 100)
	testutil.AssertNil(t, err)
	ticks := []Tick{{Value: 1.0, Label: "1.0"}, {Value: 2.0, Label: "2.0"}, {Value: 3.0, Label: "3.0"}}
	ya := YAxis{}
	yab := ya.Measure(r, NewBox(0, 0, 100, 100), &ContinuousRange{Min: 1.0, Max: 3.0, Domain: 100}, style, ticks)
	testutil.AssertEqual(t, 32, yab.Width())
	testutil.AssertEqual(t, 110, yab.Height())
}

func TestYAxisSecondaryMeasure(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)
	style := Style{
		Font:     f,
		FontSize: 10.0,
	}
	r, err := PNG(100, 100)
	testutil.AssertNil(t, err)
	ticks := []Tick{{Value: 1.0, Label: "1.0"}, {Value: 2.0, Label: "2.0"}, {Value: 3.0, Label: "3.0"}}
	ya := YAxis{AxisType: YAxisSecondary}
	yab := ya.Measure(r, NewBox(0, 0, 100, 100), &ContinuousRange{Min: 1.0, Max: 3.0, Domain: 100}, style, ticks)
	testutil.AssertEqual(t, 32, yab.Width())
	testutil.AssertEqual(t, 110, yab.Height())
}
