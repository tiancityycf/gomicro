package conf

// RedisConfig redis 配置
type RedisConfig interface {
	GetEnabled() bool
	GetConn() string
	GetPassword() string
	GetDBNum() int
	GetSentinelConfig() RedisSentinelConfig
}

// RedisSentinelConfig 哨兵配置
type RedisSentinelConfig interface {
	GetEnabled() bool
	GetMaster() string
	GetNodes() []string
}

// defaultMysqlConfig mysql 配置
type defaultRedisConfig struct {
	Conn           string              `json:"conn"`
	Enabled        bool                `json:"enabled"`
	Password       string              `json:"password"`
	DBNum          int                 `json:"db_num"`
	SentinelConfig RedisSentinelConfig `json:"sentinel_config"`
}

func (m defaultRedisConfig) GetConn() string {
	return m.Conn
}
func (m defaultRedisConfig) GetEnabled() bool {
	return m.Enabled
}
func (m defaultRedisConfig) GetPassword() string {
	return m.Password
}

func (m defaultRedisConfig) GetDBNum() int {
	return m.DBNum
}
func (m defaultRedisConfig) GetSentinelConfig() RedisSentinelConfig {
	return m.SentinelConfig
}
