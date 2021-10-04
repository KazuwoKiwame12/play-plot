package main

import (
	"fmt"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	NUM_OF_DATA = 2
	ONE_WEEK    = 7
)

type sleepRecord struct {
	bedinTime float64
	wakeTime  float64
}

func main() {
	sleepRecords := []sleepRecord{
		{
			bedinTime: -2,
			wakeTime:  6,
		},
		{
			bedinTime: -1,
			wakeTime:  7,
		},
		{
			bedinTime: 0,
			wakeTime:  8,
		},
		{
			bedinTime: -2,
			wakeTime:  6,
		},
		{
			bedinTime: -4,
			wakeTime:  5,
		},
		{
			bedinTime: 2,
			wakeTime:  9,
		},
		{
			bedinTime: 1,
			wakeTime:  6,
		},
	}
	days := make([]plotter.Values, ONE_WEEK)
	for i, sleepRecord := range sleepRecords {
		days[i] = plotter.Values{sleepRecord.bedinTime, sleepRecord.wakeTime}
	}

	p := plot.New()

	p.Title.Text = "Sleep Record"
	p.Y.Label.Text = "time(based AM0:00)"

	now := time.Now()
	if err := plotutil.AddBoxPlots(p, vg.Points(20),
		fmt.Sprintf("%d/%d", int(now.AddDate(0, 0, -6).Month()), now.AddDate(0, 0, -6).Day()), days[6],
		fmt.Sprintf("%d/%d", int(now.AddDate(0, 0, -5).Month()), now.AddDate(0, 0, -5).Day()), days[5],
		fmt.Sprintf("%d/%d", int(now.AddDate(0, 0, -4).Month()), now.AddDate(0, 0, -4).Day()), days[4],
		fmt.Sprintf("%d/%d", int(now.AddDate(0, 0, -3).Month()), now.AddDate(0, 0, -3).Day()), days[3],
		fmt.Sprintf("%d/%d", int(now.AddDate(0, 0, -2).Month()), now.AddDate(0, 0, -2).Day()), days[2],
		fmt.Sprintf("%d/%d", int(now.AddDate(0, 0, -1).Month()), now.AddDate(0, 0, -1).Day()), days[1],
		fmt.Sprintf("%d/%d", int(now.Month()), now.Day()), days[0],
	); err != nil {
		panic(err)
	}

	if err := p.Save(12*vg.Inch, 7*vg.Inch, "sleeprecord-plot.png"); err != nil {
		panic(err)
	}
}
