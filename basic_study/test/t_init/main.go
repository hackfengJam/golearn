package main

import (
	"fmt"
	"golearn/basic_study/test/t_init/t_init"
)

func main() {
	t_init.AA()

	strs := []string{
		// "inset_%d_title,插画 %d 标题",
		// "inset_%d_desc,插画 %d 描述",
		// "inset_%d_img_url,插画 %d 图片地址",
		// "inset_%d_img_url_popup,插画 %d 图片地址（弹窗插画图片）",
		"inset_%d_img_url_share,插画 %d 图片地址（插画分享图片）",
	}

	for i := 214; i <= 225; i++ {
		for _, str := range strs {
			fmt.Printf(str+"\n", i, i)
		}
	}
}
