package model

// Page is for data paging.
type Page struct {
	// Page is current page.
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// pageSize is the data number shown per page.
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// Total is the total number of referents.
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	// Pages is the number of pages.
	Pages int32 `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
}

// OrderBy determines the data list order.
type OrderBy int32

const (
	// Desc stands for descending order.
	Desc OrderBy = 0
	// Asc stands for ascending order.
	Asc OrderBy = 1
)

// SortBy determines the data list order reference.
type SortBy string
