package filter

import (
	"math"

	"github.com/svartvalp/paoias2/internal/model"
)

const mathC = 10.0

type Filter struct {
	coefficients []int
	lineCount    int
	filterCount  int
	w            float64
	stepW        float64
}

func (f *Filter) Run() []model.ArgFunc {
	sch := make([]model.ArgFunc, 0)
	for i := 0.0; i < f.w; i += f.stepW {
		sch = append(sch, f.lineTime(i)...)
	}
	return sch
}

func (f *Filter) lineTime(i float64) []model.ArgFunc {
	arr := make([]model.ArgFunc, 0)
	for d := 0; d < f.lineCount; d++ {
		f := model.ArgFunc{
			T: d,
			Y: f.calcNewRes(d, i),
		}
		arr = append(arr, f)
	}
	return arr
}

func (f *Filter) calcNewRes(d int, i float64) float64 {
	var sum float64
	for t := d; t < f.filterCount+d; t++ {
		sum += float64(f.coefficients[t-d]) * mathC * math.Sin(i*float64(t))
	}
	return sum
}

func New(
	coefficients []int,
	lineCount int,
	filterCount int,
	w float64,
	stepW float64,
) *Filter {
	return &Filter{
		coefficients: coefficients,
		lineCount:    lineCount,
		filterCount:  filterCount,
		w:            w,
		stepW:        stepW,
	}
}
