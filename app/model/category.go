// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

// 栏目树形列表
type CategoryTreeItem struct {
	Id       uint                `json:"id"`              // 分类ID，自增主键
	ParentId uint                `json:"parent_id"`       // 父级分类ID，用于层级管理
	Name     string              `json:"name"`            // 分类名称
	Thumb    string              `json:"thumb"`           // 封面图
	Brief    string              `json:"brief"`           // 简述
	Content  string              `json:"content"`         // 详细介绍
	Indent   string              `json:"indent"`          // 缩进字符串，包含：&nbsp;, " │", " ├", " └"
	Items    []*CategoryTreeItem `json:"items,omitempty"` // 子级数据项
}

// 查询栏目树形列表
type CategoryGetTreeReq struct {
	// 栏目类型：topic, question, article
	// 当传递空时表示获取所有类型的栏目
	ContentType string
}
