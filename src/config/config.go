package config

type (
	Config struct {
		Server     Server     `yaml:"server"`
		Database   Database   `yaml:"database"`
		Encryption Encryption `yaml:"encryption"`
		I18n       I18n       `yaml:"i18n"`
	}
	I18n struct {
		BundlePath string `yaml:"bundle_path"`
	}
	Server struct {
		Name               string `yaml:"name"`
		Port               int    `yaml:"port"`
		GrpcClientPort     int    `yaml:"grpc_client_port"`
		GrpcManagerPort    int    `yaml:"grpc_manager_port"`
		ClientAccessToken  string `yaml:"client_access_token"`
		ManagerAccessToken string `yaml:"manager_access_token"`
	}
	Database struct {
		MongoDb MongoDb `yaml:"mongodb"`
		Redis   Redis   `yaml:"redis"`
	}
	MongoDb struct {
		Url      string `yaml:"url"`
		Database string `yaml:"database"`
	}
	Redis struct {
		Server   string `yaml:"server"`
		Password string `yaml:"password"`
		Database int    `yaml:"database"`
	}
	Encryption struct {
		Key string `yaml:"key"`
	}
)
