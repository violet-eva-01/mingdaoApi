// Package types @author: Violet-Eva @date  : 2025/6/10 @notes :
package types

type WorkSheetRequestBody struct {
	AppKey          string   `json:"appKey"`
	Sign            string   `json:"sign"`
	WorksheetId     string   `json:"worksheetId,omitempty"`
	ViewId          string   `json:"viewId,omitempty"`
	RowId           string   `json:"rowId,omitempty"`
	PageSize        int      `json:"pageSize,omitempty"`
	PageIndex       int      `json:"pageIndex,omitempty"`
	ListType        int      `json:"listType,omitempty"`
	Controls        []string `json:"controls,omitempty"`
	Filters         []Filter `json:"filters,omitempty"`
	SortId          string   `json:"sortId,omitempty"`
	IsAsc           bool     `json:"isAsc,omitempty"`
	NotGetTotal     bool     `json:"notGetTotal,omitempty"`
	UseControlId    bool     `json:"useControlId,omitempty"`
	IsSystemControl bool     `json:"getSystemControl,omitempty"`
}

func NewWSReqBody() *WorkSheetRequestBody {
	return &WorkSheetRequestBody{}
}

func (req *WorkSheetRequestBody) SetAppKey(appKey string) *WorkSheetRequestBody {
	req.AppKey = appKey
	return req
}

func (req *WorkSheetRequestBody) SetSign(sign string) *WorkSheetRequestBody {
	req.Sign = sign
	return req
}

func (req *WorkSheetRequestBody) SetWorksheetId(worksheetId string) *WorkSheetRequestBody {
	req.WorksheetId = worksheetId
	return req
}

func (req *WorkSheetRequestBody) SetViewId(viewId string) *WorkSheetRequestBody {
	req.ViewId = viewId
	return req
}

func (req *WorkSheetRequestBody) SetRowId(rowId string) *WorkSheetRequestBody {
	req.RowId = rowId
	return req
}

func (req *WorkSheetRequestBody) SetPageSize(pageSize int) *WorkSheetRequestBody {
	req.PageSize = pageSize
	return req
}

func (req *WorkSheetRequestBody) SetPageIndex(pageIndex int) *WorkSheetRequestBody {
	req.PageIndex = pageIndex
	return req
}

func (req *WorkSheetRequestBody) SetListType(listType int) *WorkSheetRequestBody {
	req.ListType = listType
	return req
}

func (req *WorkSheetRequestBody) SetControls(controls ...string) *WorkSheetRequestBody {
	req.Controls = controls
	return req
}

func (req *WorkSheetRequestBody) SetFilters(filters ...Filter) *WorkSheetRequestBody {
	req.Filters = filters
	return req
}

func (req *WorkSheetRequestBody) SetSortId(sortId string) *WorkSheetRequestBody {
	req.SortId = sortId
	return req
}

func (req *WorkSheetRequestBody) SetAsc() *WorkSheetRequestBody {
	req.IsAsc = true
	return req
}

func (req *WorkSheetRequestBody) SetNotGetTotal() *WorkSheetRequestBody {
	req.NotGetTotal = true
	return req
}

func (req *WorkSheetRequestBody) SetUseControlId() *WorkSheetRequestBody {
	req.UseControlId = true
	return req
}

func (req *WorkSheetRequestBody) GetSystemControl() *WorkSheetRequestBody {
	req.IsSystemControl = true
	return req
}
