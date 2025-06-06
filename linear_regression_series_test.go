package chart

import (
	"testing"

	"github.com/Strubbl/go-chart/v2/testutil"
)

func TestLinearRegressionSeries(t *testing.T) {
	// replaced new assertions helper

	mainSeries := ContinuousSeries{
		Name:    "A test series",
		XValues: LinearRange(1.0, 100.0),
		YValues: LinearRange(1.0, 100.0),
	}

	linRegSeries := &LinearRegressionSeries{
		InnerSeries: mainSeries,
	}

	lrx0, lry0 := linRegSeries.GetValues(0)
	testutil.AssertInDelta(t, 1.0, lrx0, 0.0000001)
	testutil.AssertInDelta(t, 1.0, lry0, 0.0000001)

	lrxn, lryn := linRegSeries.GetLastValues()
	testutil.AssertInDelta(t, 100.0, lrxn, 0.0000001)
	testutil.AssertInDelta(t, 100.0, lryn, 0.0000001)
}

func TestLinearRegressionSeriesDesc(t *testing.T) {
	// replaced new assertions helper

	mainSeries := ContinuousSeries{
		Name:    "A test series",
		XValues: LinearRange(100.0, 1.0),
		YValues: LinearRange(100.0, 1.0),
	}

	linRegSeries := &LinearRegressionSeries{
		InnerSeries: mainSeries,
	}

	lrx0, lry0 := linRegSeries.GetValues(0)
	testutil.AssertInDelta(t, 100.0, lrx0, 0.0000001)
	testutil.AssertInDelta(t, 100.0, lry0, 0.0000001)

	lrxn, lryn := linRegSeries.GetLastValues()
	testutil.AssertInDelta(t, 1.0, lrxn, 0.0000001)
	testutil.AssertInDelta(t, 1.0, lryn, 0.0000001)
}

func TestLinearRegressionSeriesWindowAndOffset(t *testing.T) {
	// replaced new assertions helper

	mainSeries := ContinuousSeries{
		Name:    "A test series",
		XValues: LinearRange(100.0, 1.0),
		YValues: LinearRange(100.0, 1.0),
	}

	linRegSeries := &LinearRegressionSeries{
		InnerSeries: mainSeries,
		Offset:      10,
		Limit:       10,
	}

	testutil.AssertEqual(t, 10, linRegSeries.Len())

	lrx0, lry0 := linRegSeries.GetValues(0)
	testutil.AssertInDelta(t, 90.0, lrx0, 0.0000001)
	testutil.AssertInDelta(t, 90.0, lry0, 0.0000001)

	lrxn, lryn := linRegSeries.GetLastValues()
	testutil.AssertInDelta(t, 80.0, lrxn, 0.0000001)
	testutil.AssertInDelta(t, 80.0, lryn, 0.0000001)
}
