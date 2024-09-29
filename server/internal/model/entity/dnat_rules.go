// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DnatRules is the golang structure for table dnat_rules.
type DnatRules struct {
	Id        int64  `json:"id"        orm:"id"         ` // 主键
	Protocol  string `json:"protocol"  orm:"protocol"   ` //
	Dip       string `json:"dip"       orm:"dip"        ` //
	Iif       string `json:"iif"       orm:"iif"        ` //
	Port      string `json:"port"      orm:"port"       ` //
	Dnat      string `json:"dnat"      orm:"dnat"       ` //
	Comment   string `json:"comment"   orm:"comment"    ` // 备注，无意义给用户看的
	Position  int    `json:"position"  orm:"position"   ` // 规则位置，重启服务时按此从小到大排序
	CreatedAt int64  `json:"createdAt" orm:"created_at" ` //
	DeletedAt int64  `json:"deletedAt" orm:"deleted_at" ` //
}
