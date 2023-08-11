package constants

const (
	CertPath string = "/.docker/nginx/cert.pem"
	KeyPath  string = "/.docker/nginx/key.pem"
)

const (
	EnvConfPath   string = ".env"
	ParamConfPath string = "conf"
)

const (
	BucketName string = "zoomer"
)

const (
	UploadPath       string = "public/upload/"
	UploadPathReturn string = "/upload/"
	UploadPathInit   string = "public/upload"
)

const (
	RequestLogPath string = "logs/access.log"
	ErrorLogPath   string = "logs/errors.log"
	SystemLogPath  string = "logs/system.log"
)
