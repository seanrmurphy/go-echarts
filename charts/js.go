package charts

import (
	"encoding/json"
	"log"
	"syscall/js"

	"github.com/fatih/structs"
)

func generateTitleOpts(o TitleOpts) interface{} {
	m := structs.Map(o)
	//r := js.ValueOf(m)
	//r, _ := json.Marshal(o)
	return m
}

func generateTooltipOpts(o TooltipOpts) interface{} {
	m := structs.Map(o)
	//r := js.ValueOf(m)
	return m
}

func generateLegend(o LegendOpts) interface{} {
	m := structs.Map(o)
	//r := jsmValueOf(m)
	return m
}

func generateGeo(c int) interface{} {
	return c
}

func generateRadar(c int) interface{} {
	return c
}

func generateParallel(c int) interface{} {
	return c
}

func generateParallelAxis(c int) interface{} {
	return c
}

func generateSingleAxis(c int) interface{} {
	return c
}

func generateToolboxOpts(o ToolboxOpts) interface{} {
	m := structs.Map(o)
	log.Printf("toolboxopts = %v\n", m)
	r := js.ValueOf(m)
	return r
}

func generateDataZoomOpts(o DataZoomOptsList) interface{} {
	r := make([]interface{}, len(o))
	for i, op := range o {
		m := structs.Map(op)
		jsVal := js.ValueOf(m)
		r[i] = jsVal
	}
	return r
}

func generateVisualMapOptsList(o VisualMapOptsList) interface{} {
	r, _ := json.Marshal(o)
	return r
}

func generateXAxis(o []XAxisOpts) interface{} {
	r := make([]interface{}, len(o))
	for i, op := range o {
		m := structs.Map(op)
		jsVal := js.ValueOf(m)
		r[i] = jsVal
	}
	return r
}

func generateYAxis(o []YAxisOpts) interface{} {
	r := make([]interface{}, len(o))
	for i, op := range o {
		m := structs.Map(op)
		jsVal := js.ValueOf(m)
		r[i] = jsVal
	}
	return r
}

func generateXAxis3D(c int) interface{} {
	return c
}

func generateYAxis3D(c int) interface{} {
	return c
}

func generateZAxis3D(c int) interface{} {
	return c
}

func generateSeries(s Series) interface{} {
	r := make([]interface{}, len(s))
	for i, op := range s {
		m := structs.Map(op)
		log.Printf("map = %v\n", m)
		d := make([]interface{}, len(op.Data.([]int)))
		for _, q := range op.Data.([]int) {
			d[i] = q
		}
		m["data"] = d
		jsVal := js.ValueOf(m)
		r[i] = jsVal
	}
	return r
}

func generateColors(o ColorOpts) interface{} {
	r := make([]interface{}, len(o))
	for i, s := range o {
		r[i] = s
	}
	return r
}

func generateBackgroundColors(c string) interface{} {
	return c
}

// This function takes a given chart and generates the set of options in a
// JS compatible Go struct which can be passed to the setOption method within
// the echart library
//func GenerateOptions(c charter) (r map[string]interface{}) {

////let myChart___x__{{ .ChartID }}__x__ = echarts.init(document.getElementById('{{ .ChartID }}'), "{{ .Theme }}");
////let option___x__{{ .ChartID }}__x__ = {
//r = make(map[string]interface{})

////title: {{ .TitleOpts  }},
//r["title"] = generateTitleOpts(c.BaseOpts.TitleOpts)

////tooltip: {{ .TooltipOpts }},
//r["tooltip"] = generateTooltipOpts(c.TooltipOpts)

////legend: {{ .LegendOpts }},
//r["legend"] = generateLegend(c.LegendOpts)

////{{- if .HasGeo }}
////geo: {{ .GeoComponentOpts }},
////{{- end }}
//if c.HasGeo {
//r["geo"] = generateGeo(c.GeoComponentOpts)
//}

////{{- if .HasRadar }}
////radar: {{ .RadarComponentOpts }},
////{{- end }}
//if c.HasRadar {
//r["radar"] = generateRadar(c.RadarComponentOpts)
//}

////{{- if .HasParallel }}
////parallel: {{ .ParallelComponentOpts }},
////parallelAxis: {{ .ParallelAxisOpts }},
////{{- end }}
//if c.HasParallel {
//r["parallel"] = generateParallel(c.ParallelComponentOpts)
//r["parallelAxis"] = generateParallelAxis(c.ParallelAxisOpts)
//}

////{{- if .HasSingleAxis }}
////singleAxis: {{ .SingleAxisOpts }},
////{{- end }}
//if c.HasSingleAxis {
//r["singleAxis"] = generateSingleAxis(c.SingleAxisOpts)
//}

////{{- if .ToolboxOpts.Show }}
////toolbox: {{ .ToolboxOpts }},
////{{- end }}
//if c.ToolboxOpts.Show {
//r["toolbox"] = generateToolboxOpts(c.ToolboxOpts)
//}

////{{- if gt .DataZoomOptsList.Len 0 }}
////dataZoom:{{ .DataZoomOptsList }},
////{{- end }}
//if len(c.DataZoomOptsList) > 0 {
//r["dataZoom"] = generateDataZoomOpts(c.DataZoomOptsList)
//}

////{{- if gt .VisualMapOptsList.Len 0 }}
////visualMap:{{ .VisualMapOptsList }},
////{{- end }}
//if len(c.VisualMapOptsList) > 0 {
//r["visualMap"] = generateVisualMapOptsList(c.VisualMapOptsList)
//}

////{{- if .HasXYAxis }}
////xAxis: {{ .XAxisOptsList }},
////yAxis: {{ .YAxisOptsList }},
////{{- end }}
//if c.HasXYAxis {
//r["xAxis"] = generateXAxis(c.XAxisOptsList)
//r["yAxis"] = generateYAxis(c.YAxisOptsList)
//}

////{{- if .Has3DAxis }}
////xAxis3D: {{ .XAxis3D }},
////yAxis3D: {{ .YAxis3D }},
////zAxis3D: {{ .ZAxis3D }},
////grid3D: {{ .Grid3D }},
////{{- end }}
//if c.Has3DAxis {
//r["xAxis3D"] = generateXAxis3D(c.XAxis3D)
//r["yAxis3D"] = generateYAxis3D(c.YAxis3D)
//r["zAxis3D"] = generateZAxis3D(c.ZAxis3D)
//}

////series: [
////{{ range .Series }}
////{{- . }},
////{{ end -}}
////],
//r["series"] = generateSeries(c.Series)

////{{- if eq .Theme "white" }}
////color: {{ .Colors }},
////{{- end }}
//if c.Theme == "white" {
//r["color"] = generateColors(c.Colors)
//}

////{{- if ne .BackgroundColor "" }}
////backgroundColor: {{ .BackgroundColor }}
////{{- end }}
//if c.BackgroundColor != "" {
//r["backgroundColor"] = generateBackgroundColors(c.BackgroundColor)
//}

////};
//return
//}
