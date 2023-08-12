package adapter

import (
	"github.com/thuongtruong1009/zoomer/internal/server/api"
	"log"
)

func WithVersion(version string) Options {
	return func(opts *api.IApi) error {
		log.Printf("Starting API version: %s\n", version)
		// opts.Version = version
		return nil
	}
}

// func WithPort(port string) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API port: %s\n", port)
// 		return nil
// 	}
// }

// func WithHost(host string) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API host: %s\n", host)
// 		return nil
// 	}
// }

// func WithLogger(logger *log.Logger) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API logger: %s\n", logger)
// 		return nil
// 	}
// }

// func WithInterceptor(interceptor *api.Interceptor) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API interceptor: %s\n", interceptor)
// 		return nil
// 	}
// }

// func WithNotify(notify chan error) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API notify: %s\n", notify)
// 		return nil
// 	}
// }

// func WithEcho(echo *api.Echo) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API echo: %s\n", echo)
// 		return nil
// 	}
// }

// func WithConfiguration(configuration *api.Configuration) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API configuration: %s\n", configuration)
// 		return nil
// 	}
// }

// func WithParameterConfig(parameterConfig *api.ParameterConfig) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API parameterConfig: %s\n", parameterConfig)
// 		return nil
// 	}
// }

// func WithPgDB(pgDB *api.PgDB) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API pgDB: %s\n", pgDB)
// 		return nil
// 	}
// }

// func WithRedisDB(redisDB *api.RedisDB) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API redisDB: %s\n", redisDB)
// 		return nil
// 	}
// }

// func WithMinioClient(minioClient *api.MinioClient) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API minioClient: %s\n", minioClient)
// 		return nil
// 	}
// }

// func WithResourceAdapter(resourceAdapter *api.ResourceAdapter) Options {
// 	return func(opts *api.IApi) error {
// 		log.Printf("Starting API resourceAdapter: %s\n", resourceAdapter)
// 		return nil
// 	}
// }
