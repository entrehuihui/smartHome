package databases

// 用户列表
type Userinfo struct {
	ID         int    `gorm:"primary_key;unique_index;AUTO_INCREMENT:10000"` // 平台唯一ID
	Nicename   string `gorm:"not null;size:100"`                             // 昵称
	Emails     string `gorm:"not null;size:100;index"`                       // 登录账户
	Password   string `gorm:"not null;size:32"`                              // 密码
	Createtime int64  `gorm:"not null"`                                      // 创建时间
	Updatetime int64  `gorm:"not null"`                                      // 更新时间
	Del        int    `gorm:"not null"`                                      // 删除 1删除
	State      int    `gorm:"not null"`                                      // 1未激活 2已激活 3停用
}

// 设备列表
type Device struct {
	ID         int    `gorm:"primary_key;unique_index;AUTO_INCREMENT:10000"` // 平台唯一ID
	Name       string `gorm:"not null;size:100"`                             // 设备名称
	Sn         string `gorm:"not null;size:100;index"`                       // 设备SN
	Createtime int64  `gorm:"not null"`                                      // 创建时间
	Updatetime int64  `gorm:"not null"`                                      // 更新时间
	Userid     int    `gorm:"not null;index;default:-1"`                     // 所属用户
	Del        int    `gorm:"not null"`                                      // 删除 1删除
	Groupid    int    `gorm:"not null;index;default:-1"`                     // 设备组ID
	Details    string `gorm:"size:1000;default:''"`                          // 备注
	Typeid     int    `gorm:"not null;index;default:-1"`                     // 设备类型ID
	Devicekey  string `gorm:"not null;size:32"`                              // 设备连接密钥
	Devicemd5  string `gorm:"not null;size:32"`                              // 设备连接密钥MD5值
}

// 设备组
type Groupinfo struct {
	ID         int    `gorm:"primary_key;unique_index;AUTO_INCREMENT:10000"` // 平台唯一ID
	Name       string `gorm:"not null;size:100"`                             // 设备名称
	Createtime int64  `gorm:"not null"`                                      // 创建时间
	Userid     int    `gorm:"not null;index"`                                // 所属用户
	Del        int    `gorm:"not null"`                                      // 删除 1删除
	Details    string `gorm:"size:1000;default:''"`                          // 备注
}

// 设备类型
type Typesinfo struct {
	ID         int    `gorm:"primary_key;unique_index;AUTO_INCREMENT:10000"` // 平台唯一ID
	Name       string `gorm:"not null;size:100"`                             // 设备名称
	Createtime int64  `gorm:"not null"`                                      // 创建时间
	Userid     int    `gorm:"not null;index"`                                // 所属用户
	Del        int    `gorm:"not null"`                                      // 删除 1删除
	Details    string `gorm:"size:1000;default:''"`                          // 备注
}

// 设备联动
type Linkageinfo struct {
	ID         int    `gorm:"primary_key;unique_index;AUTO_INCREMENT:10000"` // 平台唯一ID
	Name       string `gorm:"not null;size:100"`                             // 设备名称
	Createtime int64  `gorm:"not null"`                                      // 创建时间
	Userid     int    `gorm:"not null;index"`                                // 所属用户
	Del        int    `gorm:"not null"`                                      // 删除 1删除
	Details    string `gorm:"size:1000;default:''"`                          // 备注
}
