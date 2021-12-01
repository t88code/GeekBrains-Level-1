package config

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/namsral/flag"
)

func LoadConfig(filename string) (string, error) {

	var port, dbUrl, jaegerUrl, sentryUrl, kafkaBroker, someAppId, someAppKey string

	flag.String(flag.DefaultConfigFlagname, filename, "Путь к файлу конфигурации")
	flag.StringVar(&port, "port", "", "Порт запускаемого сервера")
	flag.StringVar(&dbUrl, "db_url", "", "URL для подключения к базе данных")
	flag.StringVar(&jaegerUrl, "jaeger_url", "", "URL для подключения к jaeger-агенту")
	flag.StringVar(&sentryUrl, "sentry_url", "", "URL для подключения к sentry")
	flag.StringVar(&kafkaBroker, "kafka_broker", "", "URL для подключениz к брокеру kafka")
	flag.StringVar(&someAppId, "some_app_id", "", "Идентификатор приложения")
	flag.StringVar(&someAppKey, "some_app_key", "", "Ключ доступа")
	flag.Parse()

	if validURL := govalidator.IsPort(port); !validURL {
		return "", fmt.Errorf("в параметре port указан не валидный порт(1-65535) - %s", port)
	}

	if validURL := govalidator.IsURL(dbUrl); !validURL {
		return "", fmt.Errorf("в параметре db_url указан не валидный URL - %s", dbUrl)
	}

	if validURL := govalidator.IsURL(jaegerUrl); !validURL {
		return "", fmt.Errorf("в параметре jaeger_url указан не валидный URL - %s", jaegerUrl)
	}

	if validURL := govalidator.IsURL(sentryUrl); !validURL {
		return "", fmt.Errorf("в параметре sentry_url указан не валидный URL - %s", sentryUrl)
	}

	if validURL := govalidator.IsURL(kafkaBroker); !validURL {
		return "", fmt.Errorf("в параметре kafka_broker указан не валидный URL - %s", kafkaBroker)
	}

	result := fmt.Sprintf("port=%s - %s\n", port, "Порт запускаемого сервера") +
		fmt.Sprintf("db_url=%s - %s\n", dbUrl, "URL для подключения к базе данных") +
		fmt.Sprintf("jaeger_url=%s - %s\n", jaegerUrl, "URL для подключения к jaeger-агенту") +
		fmt.Sprintf("sentry_url=%s - %s\n", sentryUrl, "URL для подключения к sentry") +
		fmt.Sprintf("kafka_broker=%s - %s\n", kafkaBroker, "URL для подключениz к брокеру kafka") +
		fmt.Sprintf("some_app_id=%s - %s\n", someAppId, "Идентификатор приложения") +
		fmt.Sprintf("some_app_key=%s - %s\n", someAppKey, "Ключ доступа")

	return result, nil
}
