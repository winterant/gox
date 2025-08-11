package xconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func LoadYaml(path string, conf any, envPrefix string) *viper.Viper {
	v := viper.New()

	// 读取原始文件内容
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// 展开环境变量
	expanded := os.ExpandEnv(string(content))

	// 加载进 Viper（通过读取字符串）
	v.SetConfigType("yaml")
	err = v.ReadConfig(strings.NewReader(expanded))
	if err != nil {
		panic(fmt.Errorf("load config error: %w", err))
	}

	// 允许自动读取环境变量覆盖配置文件。例如 {envPrefix}_LOG_MAX_DAYS 将覆盖配置中的 log.max-days
	v.AutomaticEnv()
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")) // 将配置Key中的点号(.)和横杠(-)替换为下划线(_)。

	if conf != nil {
		if err := v.Unmarshal(conf); err != nil {
			panic(fmt.Errorf("unmarshal config error: %w", err))
		}
	}
	return v
}
