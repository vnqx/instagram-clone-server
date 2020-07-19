package database

func setIfNotEmpty(m map[string]string, key, val string) {
	if val != "" {
		m[key] = val
	}
}

func dbValues(config *Config) map[string]string {
	p := map[string]string{}
	setIfNotEmpty(p, "dbname", config.Name)
	setIfNotEmpty(p, "host", config.Host)
	setIfNotEmpty(p, "user", config.User)
	setIfNotEmpty(p, "password", config.Password)
	setIfNotEmpty(p, "port", config.Port)
	setIfNotEmpty(p, "sslmode", config.SSLMode)
	return p
}
