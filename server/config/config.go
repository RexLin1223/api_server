package config

// CustomizedConfig to handler
type CustomizedConfig struct {
	Domain string
	Port   string
	DBInfo DatabaseInfo
	Router RouteRule
}

// DatabaseInfo represents router of some route
type DatabaseInfo struct {
	Domain     string
	Port       string
	User       string
	Password   string
	TargetName string
}

// RouteRule is customized route
type RouteRule struct {
	Home      string
	Authorize string
	Request   string
}

// Load from config.yaml
func Load() *CustomizedConfig {

	// Fix me
	conf := CustomizedConfig{
		Domain: "0.0.0.0",
		Port:   "8857",
		Router: RouteRule{
			Home:      "/home",
			Authorize: "/auth",
		},
		DBInfo: DatabaseInfo{
			Domain:     "172.16.10.18",
			Port:       "3306",
			User:       "remote_root",
			Password:   "Aa123456",
			TargetName: "dagogo",
		},
	}
	return &conf
}
