package charts

import "github.com/fatih/structs"

// Line represents a line chart.
type Line struct {
	RectChart
}

func (Line) chartType() string { return ChartType.Line }

// LineOpts is the options set for a line chart.
type LineOpts struct {
	// 数据堆叠，同个类目轴上系列配置相同的 stack 值可以堆叠放置
	Stack string
	// 曲线是否平滑
	Smooth bool
	// 是否使用阶梯图
	Step bool
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
	// 是否连接空数据。
	ConnectNulls bool
}

func (LineOpts) MarkSeries() {}

func (opt *LineOpts) setChartOpt(s *singleSeries) {
	s.Stack = opt.Stack
	s.Smooth = opt.Smooth
	s.Step = opt.Step
	s.XAxisIndex = opt.XAxisIndex
	s.YAxisIndex = opt.YAxisIndex
	s.ConnectNulls = opt.ConnectNulls
}

// NewLine creates a new line chart.
func NewLine(routers ...RouterOpts) *Line {
	chart := new(Line)
	chart.initBaseOpts(routers...)
	chart.initXYOpts()
	chart.HasXYAxis = true
	return chart
}

// AddXAxis adds the X axis.
func (c *Line) AddXAxis(xAxis interface{}) *Line {
	c.xAxisData = xAxis
	return c
}

// AddYAxis adds the Y axis.
func (c *Line) AddYAxis(name string, yAxis interface{}, options ...SeriesOptser) *Line {
	series := singleSeries{Name: name, Type: ChartType.Line, Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

// This function takes a given chart and generates the set of options in a
// JS compatible Go struct which can be passed to the setOption method within
// the echart library
func (c *Line) GenerateOptions() (r map[string]interface{}) {
	c.validateOpts()

	r = make(map[string]interface{})

	r["title"] = structs.Map(c.TitleOpts)
	r["tooltip"] = structs.Map(c.TooltipOpts)
	r["legend"] = structs.Map(c.LegendOpts)

	// showing the toolbox component is a bit tricky right now and does not
	// seem to be fully supported by go-echarts
	//if c.ToolboxOpts.Show {
	//r["toolbox"] = generateToolboxOpts(c.ToolboxOpts)
	//}

	if len(c.DataZoomOptsList) > 0 {
		r["dataZoom"] = generateDataZoomOpts(c.DataZoomOptsList)
	}

	if len(c.VisualMapOptsList) > 0 {
		r["visualMap"] = generateVisualMapOptsList(c.VisualMapOptsList)
	}

	if c.HasXYAxis {
		r["xAxis"] = generateXAxis(c.XAxisOptsList)
		r["yAxis"] = generateYAxis(c.YAxisOptsList)
	}

	r["series"] = generateSeries(c.Series)

	if c.Theme == "white" {
		r["color"] = generateColors(c.Colors)
	}

	if c.BackgroundColor != "" {
		r["backgroundColor"] = generateBackgroundColors(c.BackgroundColor)
	}
	return
}
