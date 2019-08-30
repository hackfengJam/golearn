package model

import "time"

// ip_address 表
type IpAddress struct {
	Id         uint64    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`                 // 名称
	Ip         string    `json:"ip" db:"ip"`                     // ip
	Mask       string    `json:"mask" db:"mask"`                 // 子网掩码 8, 32 即 255.0.0.0, 255.255.255.255（认为没有网络号，只代表一个主机号）
	IsAdmin    bool      `json:"is_admin" db:"is_admin"`         // 是否是 admin 级 ip。1 表示是 | 0 表示不是
	Status     bool      `json:"status" db:"status"`             // 状态 1：有效 | 0：无效
	UpdateTime time.Time `json:"updated_time" db:"updated_time"` // 最后一次更新时间
	CreateTime time.Time `json:"created_time" db:"created_time"` // 插入时间
}

// 角色表
type Role struct {
	Id         uint64    `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`                 // 角色名称
	Status     bool      `json:"status" db:"status"`             // 状态 1：有效 0：无效
	UpdateTime time.Time `json:"updated_time" db:"updated_time"` // 最后一次更新时间
	CreateTime time.Time `json:"created_time" db:"created_time"` // 插入时间
}

// object 角色表
type ObjectRole struct {
	Id         uint64    `json:"id" db:"id"`
	ObjectType int       `json:"object_type" db:"object_type"`   // 对象类型 0: unknown 1：Ip Whitelisting 2：Access Key
	ObjectId   uint64    `json:"object_id" db:"object_id"`       // 对象 ID
	RoleId     uint64    `json:"role_id" db:"role_id"`           // 角色ID
	CreateTime time.Time `json:"created_time" db:"created_time"` // 插入时间
}

// 权限详情表
type Permission struct {
	Id         uint64    `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`               // 权限名称
	Key        string    `json:"key" db:"key"`                   // 可访问的 key
	Status     bool      `json:"status" db:"status"`             // 状态 1：有效 0：无效
	UpdateTime time.Time `json:"updated_time" db:"updated_time"` // 最后一次更新时间
	CreateTime time.Time `json:"created_time" db:"created_time"` // 插入时间
}

// 角色权限表
type RolePermission struct {
	Id           uint64    `json:"id" db:"id"`
	RoleId       uint64    `json:"role_id" db:"role_id"`             // 对象 ID
	PermissionId uint64    `json:"permission_id" db:"permission_id"` // 权限 ID
	CreateTime   time.Time `json:"created_time" db:"created_time"`   // 插入时间
}
