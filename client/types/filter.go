// Package types @author: Violet-Eva @date  : 2025/6/10 @notes :
package types

type (
	DataType       int
	SpliceType     int
	FilterTypeEnum int
	DateRange      int
	DateRangeType  int
)

const (
	// Text 文本
	Text DataType = iota + 2
	// MobilePhoneNum 移动电话号码
	MobilePhoneNum
	// LandlineNum 固定电话
	LandlineNum
	// Email 邮件
	Email
	// Number 数值
	Number
	// Certificates 证件
	Certificates
	// Amount 金额
	Amount
	// TileRadio 单选 - 平铺模式
	TileRadio
	// MultipleChoices 多选
	MultipleChoices
	// DropDownRadio 单选 - 下拉模式
	DropDownRadio
	// Annex 附件
	Annex = iota + 4
	// Data 2006-01-02
	Data
	// DataTime 2006-01-02 15:04
	DataTime
)

const (
	And SpliceType = iota + 1
	Or
)

const (
	Default FilterTypeEnum = iota
	// Like 包含
	Like
	// Eq 等于
	Eq
	// Star 开头为
	Star
	// End 结尾为
	End
	// NContain 不包含
	NContain
	// Ne 不等于
	Ne
	// IsNull 为空
	IsNull
	// HasValue 不为空
	HasValue
	// Between 在范围内
	Between = iota + 2
	// NBetween 不在范围内
	NBetween
	// Gt >
	Gt
	// Gte >=
	Gte
	// Lt <
	Lt
	// Lte <=
	Lte
	// DateEnum 日期是
	DateEnum
	// NDateEnum 日期不是
	NDateEnum
	// MySelf 我拥有的
	MySelf = iota + 4
	// UnRead 未读
	UnRead
	// Sub 下属
	Sub
	// RCEq 关联控件是
	RCEq
	// RCNe 关联控件不是
	RCNe
	// ArrEq 数组等于
	ArrEq
	// ArrNe 数组不等于
	ArrNe
	// DataBetween 在范围内
	DataBetween = iota + 7
	// DateNBetween 不在范围内
	DateNBetween
	// DateGt >
	DateGt
	// DateGte >=
	DateGte
	// DateLt <
	DateLt
	// DateLte <=
	DateLte
	// NormalUser 常规用户
	NormalUser = iota + 11
	// PortalUser 外部用户
	PortalUser
)

const (
	DateDefault DateRange = iota
	Today
	Yesterday
	Tomorrow
	ThisWeek
	LastWeek
	NextWeek
	ThisMonth
	LastMonth
	NextMonth
	LastEnum
	NextEnum
	ThisQuarter
	LastQuarter
	NextQuarter
	ThisYear
	LastYear
	NextYear
	Customize
	Last7Day = iota + 2
	Last14Day
	Last30Day
	Next7Day = iota + 9
	Next14Day
	Next33Day
)

const (
	Day DateRangeType = iota + 1
	Week
	Month
	Quarter
	Year
)

type Filter struct {
	ControlId     string         `json:"controlId"`
	DataType      DataType       `json:"dataType"`
	SpliceType    SpliceType     `json:"spliceType"`
	FilterType    FilterTypeEnum `json:"filterType"`
	Value         string         `json:"value,omitempty"`
	Values        []string       `json:"values,omitempty"`
	DateRange     DateRange      `json:"dateRange,omitempty"`
	DateRangeType DateRangeType  `json:"dateRangeType,omitempty"`
}

func NewFilter(controlId string) *Filter {
	return &Filter{
		ControlId: controlId,
	}
}

func (f *Filter) SetDataType(dataType DataType) *Filter {
	f.DataType = dataType
	return f
}

func (f *Filter) SetSpliceType(spliceType SpliceType) *Filter {
	f.SpliceType = spliceType
	return f
}

func (f *Filter) SetFilterType(filterType FilterTypeEnum) *Filter {
	f.FilterType = filterType
	return f
}

func (f *Filter) SetValue(value string) *Filter {
	f.Value = value
	return f
}

func (f *Filter) SetValues(values ...string) *Filter {
	f.Values = values
	return f
}

func (f *Filter) AddValue(value string) *Filter {
	f.Values = append(f.Values, value)
	return f
}

func (f *Filter) SetDateRange(date DateRange) *Filter {
	f.DateRange = date
	return f
}

func (f *Filter) SetDateRangeType(date DateRangeType) *Filter {
	f.DateRangeType = date
	return f
}
