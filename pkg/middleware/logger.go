package middleware

import "github.com/gofiber/fiber/v2/middleware/logger"

func Logger() logger.Config {
	return logger.Config{
		Format:     "${time} ${ip} ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "2006/01/02 15:04:05",
		TimeZone:   "America/Sao_Paulo",
	}
}
