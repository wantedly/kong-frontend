package kong

type RateLimitingPluginConfig struct {
	Second        uint   `form:"second" json:"second,omitempty" binding:"omitempty"`
	Minute        uint   `form:"minute" json:"minute,omitempty" binding:"omitempty"`
	Hour          uint   `form:"hour" json:"hour,omitempty" binding:"omitempty"`
	Day           uint   `form:"day" json:"day,omitempty" binding:"omitempty"`
	Month         uint   `form:"month" json:"month,omitempty" binding:"omitempty"`
	Year          uint   `form:"year" json:"year,omitempty" binding:"omitempty"`
	LimitBy       string `form:"limit_by" json:"limit_by,omitempty" binding:"omitempty"`
	Policy        string `form:"policy" json:"policy,omitempty" binding:"omitempty"`
	FaultTolerant bool   `form:"fault_tolerant" json:"fault_tolerant,omitempty" binding:"omitempty"`
	RedisHost     string `form:"redis_host" json:"redis_host,omitempty" binding:"omitempty"`
	RedisPort     uint   `form:"redis_port" json:"redis_port,omitempty" binding:"omitempty"`
	RedisTimeout  uint   `form:"redis_timeout" json:"redis_timeout,omitempty" binding:"omitempty"`
}
