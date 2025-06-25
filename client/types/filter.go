// Package types @author: Violet-Eva @date  : 2025/6/10 @notes :
package types

type Filter struct {
	ControlId  string   `json:"controlId"`
	DataType   int      `json:"dataType"`
	SpliceType int      `json:"spliceType"`
	FilterType int      `json:"filterType"`
	Value      string   `json:"value,omitempty"`
	Values     []string `json:"values,omitempty"`
}

func NewFilter(controlId string, dataType, spliceType, filterType int, values ...string) *Filter {
	if len(values) == 0 {
		return &Filter{
			ControlId:  controlId,
			DataType:   dataType,
			SpliceType: spliceType,
			FilterType: filterType,
		}
	} else if len(values) == 1 {
		return &Filter{
			ControlId:  controlId,
			DataType:   dataType,
			SpliceType: spliceType,
			FilterType: filterType,
			Value:      values[0],
		}
	} else {
		return &Filter{
			ControlId:  controlId,
			DataType:   dataType,
			SpliceType: spliceType,
			FilterType: filterType,
			Values:     values,
		}
	}
}
