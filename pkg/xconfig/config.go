package xconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// LoadYaml 从指定路径加载 YAML 文件，并将其解析为 Viper 配置对象。
// 参数:
//
//	path: YAML 文件的路径。字段名大小写不敏感，推荐大驼峰命名；不支持短线或下划线命名。
//	conf: 一个指向要填充的配置对象的指针，该对象应与 YAML 文件的结构相匹配。如果为 nil，则不解析到该对象。
//	envPrefix: 环境变量的前缀，用于自动从环境变量中读取配置值。
//
// 返回值:
//
//	返回一个 *viper.Viper 指针，表示加载的配置对象。
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
