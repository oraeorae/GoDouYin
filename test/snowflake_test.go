package test

import (
	"fmt"
	"go_douyin/utils/snow_flake"
	"testing"
)

//测试雪花算法
func TestSnow(t *testing.T) {
	snowflake, err := snow_flake.NewSnowflake(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	id := snowflake.Generate()
	fmt.Println(id)
}
