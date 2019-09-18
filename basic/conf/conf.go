package conf

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	err                     error
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	profiles                defaultProfiles
	consulConfig            defaultConsulConfig
	mysqlConfig             defaultMysqlConfig
	m                       sync.RWMutex
	inited                  bool
	sp                      = string(filepath.Separator)
)

// Init 初始化配置
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[Init] 配置已经初始化过")
		return
	}

	// 加载yml配置
	// 先加载基础配置
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))

	pt := filepath.Join(appPath, "config")
	os.Chdir(appPath)

	// 找到application.yml文件
	if err = config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}

	// 找到需要引入的新配置文件
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	//log.Logf("[Init] 加载配置文件：path: %s, %+v\n", pt+sp+"application.yml", profiles)

	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")

		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"

			//log.Logf("[Init] 加载配置文件：path: %s\n", filePath)

			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		// 加载include的文件
		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}

	// 赋值
	config.Get(defaultRootPath, "consul").Scan(&consulConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)

	//log.Logf("consulConfig %v \n", consulConfig)
	//log.Logf("mysqlConfig %v \n", mysqlConfig)

	// 标记已经初始化
	inited = true
}

// GetMysqlConfig 获取mysql配置
func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// GetConsulConfig 获取Consul配置
func GetConsulConfig() (ret ConsulConfig) {
	return consulConfig
}

// Profiles 属性配置文件
type Profiles interface {
	GetInclude() string
}

// defaultProfiles 属性配置文件
type defaultProfiles struct {
	Include string `json:"include"`
}

// Include 包含的配置文件
// 名称前缀为"application-"，格式为yml，如："application-xxx.yml"
// 多个文件名以逗号隔开，并省略掉前缀"application-"，如：dn, jpush, mysql
func (p defaultProfiles) GetInclude() string {
	return p.Include
}

type ConsulConfig interface {
	GetHost() string
	GetEnabled() bool
	GetPort() int
}

// defaultConsulConfig 默认consul 配置
type defaultConsulConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// URL mysql 连接
func (m defaultConsulConfig) GetHost() string {
	return m.Host
}

// Enabled 激活
func (m defaultConsulConfig) GetEnabled() bool {
	return m.Enabled
}

// 闲置连接数
func (m defaultConsulConfig) GetPort() int {
	return m.Port
}

// MysqlConfig mysql 配置 接口
type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

// defaultMysqlConfig mysql 配置
type defaultMysqlConfig struct {
	URL               string `json:"url"`
	Enabled           bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

// URL mysql 连接
func (m defaultMysqlConfig) GetURL() string {
	return m.URL
}

// Enabled 激活
func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enabled
}

// 闲置连接数
func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

// 打开连接数
func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}
