package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env           string `env:"ENV" env-default:"local"`
	HTTP          HTTPConfig
	DBConfig      DBConfig
	ServiceConst  ServiceConfig
	ApiConst      ApiConfig
	URLGetCarInfo string `env:"URL_GETCARINFO"`
}

type ServiceConfig struct {
	DefaultPropertyOffset int `env:"DEFAULTPROPERTYOFFSET"`
	DefaultPropertyLimit  int `env:"DEFAULTPROPERTYLIMIT"`
}

type ApiConfig struct {
	ID         string `env:"API_ID"`         //"id"
	LIMIT      string `env:"API_LIMIT"`      //"limit"
	OFFSET     string `env:"API_OFFSET"`     //"offset"
	REGNUM     string `env:"API_REGNUM"`     //"regNum"
	MARK       string `env:"API_MARK"`       //"mark"
	MODEL      string `env:"API_MODEL"`      //"model"
	YEAR       string `env:"API_YEAR"`       //"year"
	NAME       string `env:"API_NAME"`       //name
	SURNAME    string `env:"API_SURNAME"`    //surname
	PATRONYMIC string `env:"API_PATRONYMIC"` //patronymic
}

type DBConfig struct {
	User     string `env:"DB_USER" env-default:"user"`
	Password string `env:"DB_PASSWORD" env-default:"user"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	Dbname   string `env:"DB_DBNAME" env-default:"test"`
	Sslmode  string `env:"DB_SSLMODE" env-default:"disable"`
}

type HTTPConfig struct {
	HostPort string `env:"HTTP_PORT" env-default:"8080"`
}

// загрузка конфига из .env
func MustLoad() *Config {

	if err := godotenv.Load(); err != nil {
		panic("No .env file found" + err.Error())
	}
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
