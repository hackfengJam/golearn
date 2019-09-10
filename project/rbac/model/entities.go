package model

import "time"

// ip_address 表
type IpAddress struct {
	Id         uint64    `json:"id" db:"id"`
	Name       string    `json:"name,omitempty" db:"name"`                 // 名称
	Ip         string    `json:"ip,omitempty" db:"ip"`                     // ip
	Mask       string    `json:"mask,omitempty" db:"mask"`                 // 子网掩码 8, 32 即 255.0.0.0, 255.255.255.255（认为没有网络号，只代表一个主机号）
	IsAdmin    bool      `json:"is_admin,omitempty" db:"is_admin"`         // 是否是 admin 级 ip。1 表示是 | 0 表示不是
	Status     bool      `json:"status,omitempty" db:"status"`             // 状态 1：有效 | 0：无效
	UpdateTime time.Time `json:"updated_time,omitempty" db:"updated_time"` // 最后一次更新时间
	CreateTime time.Time `json:"created_time,omitempty" db:"created_time"` // 插入时间
}

// 角色表
type Role struct {
	Id         uint64    `json:"id" db:"id"`
	Name       string    `json:"name,omitempty" db:"name"`                 // 角色名称
	Status     bool      `json:"status,omitempty" db:"status"`             // 状态 1：有效 0：无效
	UpdateTime time.Time `json:"updated_time,omitempty" db:"updated_time"` // 最后一次更新时间
	CreateTime time.Time `json:"created_time,omitempty" db:"created_time"` // 插入时间
}

// object 角色表
type ObjectRole struct {
	Id         uint64    `json:"id" db:"id"`
	ObjectType int       `json:"object_type,omitempty" db:"object_type"`   // 对象类型 0: unknown 1：Ip Whitelisting 2：Access Key
	ObjectId   uint64    `json:"object_id,omitempty" db:"object_id"`       // 对象 ID
	RoleId     uint64    `json:"role_id,omitempty" db:"role_id"`           // 角色ID
	CreateTime time.Time `json:"created_time,omitempty" db:"created_time"` // 插入时间
}

// 权限详情表
type Permission struct {
	Id         uint64    `json:"id" db:"id"`
	Title      string    `json:"title,omitempty" db:"title"`               // 权限名称
	EntityKey  string    `json:"entity_key,omitempty" db:"entity_key"`     // 可访问的 实体 key
	Status     bool      `json:"status,omitempty" db:"status"`             // 状态 1：有效 0：无效
	UpdateTime time.Time `json:"updated_time,omitempty" db:"updated_time"` // 最后一次更新时间
	CreateTime time.Time `json:"created_time,omitempty" db:"created_time"` // 插入时间
}

// 角色权限表
type RolePermission struct {
	Id           uint64    `json:"id" db:"id"`
	RoleId       uint64    `json:"role_id,omitempty" db:"role_id"`             // 对象 ID
	PermissionId uint64    `json:"permission_id,omitempty" db:"permission_id"` // 权限 ID
	CreateTime   time.Time `json:"created_time,omitempty" db:"created_time"`   // 插入时间
}

// AK 配置
type AkConfig struct {
	Id         uint64    `json:"id" db:"id"`
	Namespace  string    `json:"namespace,omitempty" db:"namespace"`       // 命名空间
	Name       string    `json:"name,omitempty" db:"name"`                 // name 唯一
	Title      string    `json:"title,omitempty" db:"title"`               // 标题
	Remark     string    `json:"remark,omitempty" db:"remark"`             // 备注
	Value      string    `json:"value,omitempty" db:"value"`               // 插入时间
	UpdateTime time.Time `json:"updated_time,omitempty" db:"updated_time"` // 最后一次更新时间
	CreateTime time.Time `json:"value,omitempty" db:"value"`               // 插入时间
}
