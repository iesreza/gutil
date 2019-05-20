package configuration

import (
	"github.com/spf13/viper"
	"gutil/log"
	"gutil/path"
	"os"
	"runtime"
	"strings"
	"time"
)

type cfg struct {
	App       string
	Type      string
	Path      []string
	vp        *viper.Viper
	container interface{}
}

func GetInstance(container interface{}) *cfg {
	cfg := cfg{}
	cfg.App = os.Args[0]
	cfg.Path = []string{"/etc/" + cfg.App + "/", "$HOME/."}
	cfg.Type = "json"
	cfg.container = container
	cfg.vp = viper.New()
	return &cfg
}

func (cfg cfg) Load() error {
	viper.Set("Verbose", true)
	cfg.vp.SetConfigName(cfg.App + "." + cfg.Type)

	if runtime.GOOS == "windows" {

		cfg.vp.SetConfigFile("./" + cfg.App + "." + cfg.Type)
		err := cfg.vp.ReadInConfig()
		if err != nil {
			log.WarningF("Config cannot be found. %s", err)
			return err
		}
	} else {
		for _, path := range cfg.Path {
			cfg.vp.AddConfigPath(path)
		}
		cfg.vp.SetConfigType(cfg.Type)
		err := cfg.vp.ReadInConfig()
		if err != nil {
			localConf := path.File("./" + cfg.App + "." + cfg.Type)
			if localConf.Exist() {
				log.WarningF("Config cannot be found. try to create %s", cfg.Path[0])
				p := path.Dir(cfg.Path[0])
				p.Create()
				localConf.Copy(strings.TrimRight(cfg.Path[0], "/") + "/" + cfg.App + ".conf")
				err = cfg.Load()
				if err != nil {
					return err
				}
			} else {
				log.WarningF("Config cannot be found. new config cannot be created due default config is not exists.")
				return nil
			}
		}
	}

	return cfg.vp.Unmarshal(cfg.container)

}

func (cfg cfg) SetDefault(key string, value interface{}) *cfg {
	cfg.vp.SetDefault(key, value)
	return &cfg
}
func (cfg cfg) Get(key string) interface{} {
	return cfg.vp.Get(key)
}
func (cfg cfg) GetBool(key string) bool {
	return cfg.vp.GetBool(key)
}
func (cfg cfg) GetFloat64(key string) float64 {
	return cfg.vp.GetFloat64(key)
}
func (cfg cfg) GetInt(key string) int {
	return cfg.vp.GetInt(key)
}
func (cfg cfg) GetString(key string) string {
	return cfg.vp.GetString(key)
}
func (cfg cfg) GetStringMap(key string) map[string]interface{} {
	return cfg.vp.GetStringMap(key)
}
func (cfg cfg) GetStringMapString(key string) map[string]string {
	return cfg.vp.GetStringMapString(key)
}
func (cfg cfg) GetStringSlice(key string) []string {
	return cfg.vp.GetStringSlice(key)
}
func (cfg cfg) GetTime(key string) time.Time {
	return cfg.vp.GetTime(key)
}
func (cfg cfg) GetDuration(key string) time.Duration {
	return cfg.vp.GetDuration(key)
}
func (cfg cfg) IsSet(key string) bool {
	return cfg.vp.IsSet(key)
}

func (cfg cfg) Set(key string, value interface{}) {
	cfg.vp.Set(key, value)
	cfg.vp.Unmarshal(cfg.container)
}

func (cfg cfg) Update() error {
	return cfg.vp.WriteConfig()
}
