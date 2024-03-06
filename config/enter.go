package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Upload   Upload   `yaml:"upload"`
	QQ       QQ       `yaml:"qq"`
	Email    Email    `yaml:"email"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Jwt      Jwt      `yaml:"jwt"`
	Redis    Redis    `yaml:"redis"`
	ES       ES       `yaml:"es"`
}
