package main

import (
	"fmt"
	"os"
	"strings"
)

func getUserTableF(sqlTmpl string, index int64) string {
	return fmt.Sprintf(sqlTmpl, index, index)
}

func genSubTables(mod int, filename string, sqlTmpls []string) {
	sqls := []string{}
	for i := 0; i < mod; i++ {
		for _, sqlTmpl := range sqlTmpls {
			if strings.Index(sqlTmpl, "#") >= 0 {
				// 报错
				fmt.Println(fmt.Errorf("has #"))
				return
			}
			sqls = append(sqls, getUserTableF(sqlTmpl, int64(i)))
		}
	}
	f, err := os.OpenFile(filename, os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(strings.Join(sqls, "\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// var sqlTmpls = []string{
	// 	"DROP TABLE IF EXISTS `send_award_log_%d`; CREATE TABLE `send_award_log_%d` ( `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT, `aid` varchar(32) NOT NULL DEFAULT '' COMMENT 'account ID', `region` varchar(32) NOT NULL DEFAULT '' COMMENT '游戏区服', `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID', `lang` varchar(16) NOT NULL DEFAULT '' COMMENT '语言code', `ticket` varchar(64) NOT NULL DEFAULT '' COMMENT '票据', `shop_reward_no` int(11) NOT NULL DEFAULT '0' COMMENT '奖励商店的奖励编号', `reward_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发奖平台：发放奖励id', `reward_num` int(11) NOT NULL DEFAULT '1' COMMENT '发奖平台：发放奖励id的数量', `error_message` text NOT NULL COMMENT '错误原因', `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态，0->失败，1->成功', `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间', `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次更新时间', PRIMARY KEY (`id`), KEY `idx_region_uid` (`region`, `uid`, `shop_reward_no`), KEY `idx_create_time` (`create_time`), KEY `idx_update_time` (`update_time`) ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT ='奖励兑换发放日志表';",
	// }
	genSubTables(20, "basic_study/test/db_sub_table/tables.sql", sqlTmpls)
}
