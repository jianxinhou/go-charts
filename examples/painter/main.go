package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	charts "github.com/vicanso/go-charts"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func writeFile(buf []byte) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, "painter.png")
	err = ioutil.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	p, err := charts.NewPainter(charts.PainterOptions{
		Width:  400,
		Height: 1200,
		Type:   charts.ChartOutputPNG,
	})
	if err != nil {
		panic(err)
	}
	// 背景色
	p.SetBackground(p.Width(), p.Height(), drawing.ColorWhite)

	top := 0

	// 画线
	p.SetDrawingStyle(charts.Style{
		StrokeColor: drawing.ColorBlack,
		FillColor:   drawing.ColorBlack,
		StrokeWidth: 1,
	})
	p.LineStroke([]charts.Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 100,
			Y: 10,
		},
		{
			X: 200,
			Y: 0,
		},
		{
			X: 300,
			Y: 10,
		},
	})

	// 圆滑曲线
	// top += 50
	// p.Child(charts.PainterPaddingOption(charts.Box{
	// 	Top: top,
	// })).SetDrawingStyle(charts.Style{
	// 	StrokeColor: drawing.ColorBlack,
	// 	FillColor:   drawing.ColorBlack,
	// 	StrokeWidth: 1,
	// }).SmoothLineStroke([]charts.Point{
	// 	{
	// 		X: 0,
	// 		Y: 0,
	// 	},
	// 	{
	// 		X: 100,
	// 		Y: 10,
	// 	},
	// 	{
	// 		X: 200,
	// 		Y: 0,
	// 	},
	// 	{
	// 		X: 300,
	// 		Y: 10,
	// 	},
	// })

	// 标线
	top += 50
	p.Child(charts.PainterPaddingOption(charts.Box{
		Top: top,
	})).SetDrawingStyle(charts.Style{
		StrokeColor: drawing.ColorBlack,
		FillColor:   drawing.ColorBlack,
		StrokeWidth: 1,
		StrokeDashArray: []float64{
			4,
			2,
		},
	}).MarkLine(0, 0, p.Width())

	top += 60
	// Polygon
	p.Child(charts.PainterBoxOption(charts.Box{
		Top: top,
	})).SetDrawingStyle(charts.Style{
		StrokeColor: drawing.ColorBlack,
		FillColor:   drawing.ColorBlack,
		StrokeWidth: 1,
	}).Polygon(charts.Point{
		X: 100,
		Y: 0,
	}, 50, 6)

	// FillArea
	top += 60
	p.Child(charts.PainterPaddingOption(charts.Box{
		Top: top,
	})).SetDrawingStyle(charts.Style{
		FillColor: drawing.ColorBlack,
	}).FillArea([]charts.Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 100,
			Y: 0,
		},
		{
			X: 150,
			Y: 40,
		},
		{
			X: 80,
			Y: 30,
		},
		{
			X: 0,
			Y: 0,
		},
	})

	// 坐标轴的点
	top += 50
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   1,
			Right:  p.Width() - 1,
			Bottom: 20,
		}),
	).SetDrawingStyle(charts.Style{
		StrokeColor: drawing.ColorBlack,
		FillColor:   drawing.ColorBlack,
		StrokeWidth: 1,
	}).Ticks(charts.TicksOption{
		Count:  7,
		Length: 5,
	})

	// 坐标轴的点，每2格显示一个
	top += 20
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   1,
			Right:  p.Width() - 1,
			Bottom: 20,
		}),
	).SetDrawingStyle(charts.Style{
		StrokeColor: drawing.ColorBlack,
		FillColor:   drawing.ColorBlack,
		StrokeWidth: 1,
	}).Ticks(charts.TicksOption{
		Unit:   2,
		Count:  7,
		Length: 5,
	})

	// 坐标轴的点，纵向
	top += 20
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   1,
			Right:  p.Width() - 1,
			Bottom: top + 100,
		}),
	).SetDrawingStyle(charts.Style{
		StrokeColor: drawing.ColorBlack,
		FillColor:   drawing.ColorBlack,
		StrokeWidth: 1,
	}).Ticks(charts.TicksOption{
		Orient: charts.OrientVertical,
		Count:  7,
		Length: 5,
	})

	// 横向展示文本
	top += 120
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   1,
			Right:  p.Width() - 1,
			Bottom: 20,
		}),
	).OverrideTextStyle(charts.Style{
		FontColor: drawing.ColorBlack,
		FontSize:  10,
	}).MultiText(charts.MultiTextOption{
		TextList: []string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		},
	})

	// 横向显示文本，靠左
	top += 20
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   1,
			Right:  p.Width() - 1,
			Bottom: 20,
		}),
	).OverrideTextStyle(charts.Style{
		FontColor: drawing.ColorBlack,
		FontSize:  10,
	}).MultiText(charts.MultiTextOption{
		Position: charts.PositionLeft,
		TextList: []string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		},
	})

	// 纵向显示文本
	top += 20
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   1,
			Right:  50,
			Bottom: top + 150,
		}),
	).OverrideTextStyle(charts.Style{
		FontColor: drawing.ColorBlack,
		FontSize:  10,
	}).MultiText(charts.MultiTextOption{
		Orient: charts.OrientVertical,
		Align:  charts.AlignRight,
		TextList: []string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		},
	})
	// 纵向 文本居中
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   50,
			Right:  100,
			Bottom: top + 150,
		}),
	).OverrideTextStyle(charts.Style{
		FontColor: drawing.ColorBlack,
		FontSize:  10,
	}).MultiText(charts.MultiTextOption{
		Orient: charts.OrientVertical,
		Align:  charts.AlignCenter,
		TextList: []string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		},
	})
	// 纵向 文本置顶
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   100,
			Right:  150,
			Bottom: top + 150,
		}),
	).OverrideTextStyle(charts.Style{
		FontColor: drawing.ColorBlack,
		FontSize:  10,
	}).MultiText(charts.MultiTextOption{
		Orient:   charts.OrientVertical,
		Position: charts.PositionTop,
		Align:    charts.AlignCenter,
		TextList: []string{
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun",
		},
	})

	// grid
	top += 150
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   1,
			Right:  p.Width() - 1,
			Bottom: top + 100,
		}),
	).OverrideTextStyle(charts.Style{
		FontColor: drawing.ColorBlack,
		FontSize:  10,
	}).Grid(charts.GridOption{
		Column:            8,
		IgnoreColumnLines: []int{0, 8},
		Row:               8,
		IgnoreRowLines:    []int{0, 8},
	})

	// dots
	top += 100
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   1,
			Right:  p.Width() - 1,
			Bottom: top + 20,
		}),
	).OverrideDrawingStyle(charts.Style{
		FillColor:   drawing.ColorWhite,
		StrokeColor: drawing.ColorBlack,
		StrokeWidth: 1,
	}).Dots([]charts.Point{
		{
			X: 0,
			Y: 0,
		},
		{
			X: 50,
			Y: 0,
		},
		{
			X: 100,
			Y: 10,
		},
	})

	// rect
	top += 30
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   1,
			Right:  200,
			Bottom: top + 50,
		}),
	).OverrideDrawingStyle(charts.Style{
		StrokeColor: drawing.ColorBlack,
		FillColor:   drawing.ColorBlack,
	}).Rect(charts.Box{
		Left:   10,
		Top:    0,
		Right:  110,
		Bottom: 20,
	})
	// legend line dot
	p.Child(
		charts.PainterBoxOption(charts.Box{
			Top:    top,
			Left:   200,
			Right:  p.Width() - 1,
			Bottom: top + 50,
		}),
	).OverrideDrawingStyle(charts.Style{
		StrokeColor: drawing.ColorBlack,
		FillColor:   drawing.ColorBlack,
	}).LegendLineDot(charts.Box{
		Left:   10,
		Top:    0,
		Right:  50,
		Bottom: 20,
	})

	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}
	err = writeFile(buf)
	if err != nil {
		panic(err)
	}
}
