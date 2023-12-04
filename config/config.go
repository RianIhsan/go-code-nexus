package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	AppPort     int
	Token       string
	Database    database
	EmailConfig email
}

type database struct {
	DbHost string
	DbPort int
	DbUser string
	DbPass string
	DbName string
}

type email struct {
	Host     string
	Sender   string
	Password string
	MailPort int
}

func BootConfig() *Config {
	return loadConfig()

}

func loadConfig() *Config {
	var res = new(Config)
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	if value, found := os.LookupEnv("SERVER_PORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid server port", err.Error())
			return nil
		}
		res.AppPort = port
	}

	if value, found := os.LookupEnv("TOKEN"); found {
		res.Token = value
	}

	if value, found := os.LookupEnv("DBHOST"); found {
		res.Database.DbHost = value
	}

	if value, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid db port", err.Error())
			return nil
		}
		res.Database.DbPort = port
	}

	if value, found := os.LookupEnv("DBUSER"); found {
		res.Database.DbUser = value
	}

	if value, found := os.LookupEnv("DBPASS"); found {
		res.Database.DbPass = value
	}

	if value, found := os.LookupEnv("DBNAME"); found {
		res.Database.DbName = value
	}

	if value, found := os.LookupEnv("SMTP_USER"); found {
		res.EmailConfig.Sender = value
	}

	if value, found := os.LookupEnv("SMTP_PASS"); found {
		res.EmailConfig.Password = value
	}

	if value, found := os.LookupEnv("SMTP_PORT"); found {
		mailPort, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid smtp port", err.Error())
			return nil
		}
		res.EmailConfig.MailPort = mailPort
	}

	if value, found := os.LookupEnv("SMTP_HOST"); found {
		res.EmailConfig.Host = value
	}

	return res
}
