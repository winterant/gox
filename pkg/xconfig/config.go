package xconfig

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func LoadYaml(path string, conf any, envPrefix string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)

	// 允许自动读取环境变量覆盖配置文件。例如 {envPrefix}_LOG_MAX_DAYS 将覆盖配置中的 log.max-days
	v.AutomaticEnv()
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")) // 将配置Key中的点号(.)和横杠(-)替换为下划线(_)。

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config file error: %w", err))
	}
	if conf != nil {
		if err := v.Unmarshal(conf); err != nil {
			panic(fmt.Errorf("unmarshal config error: %w", err))
		}
	}
	return v
}
