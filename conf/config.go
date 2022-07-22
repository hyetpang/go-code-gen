package conf

type Config struct {
	Global  Global    `mapstructure:"global"`
	Methods []*Method `mapstructure:"methods" validate:"min=1"`
}

// 全局配置
type Global struct {
	ModelName              string `mapstructure:"model_name"`
	DocTag                 string `mapstructure:"doc_tag"`
	LogicPath              string `mapstructure:"logic_path"`
	RepoName               string `mapstructure:"repo_name"`
	DependencyName         string `mapstructure:"dependency_name"`
	AddIpToReqParam        *bool  `mapstructure:"add_ip_to_req_param"`
	AddUserIdToReqParam    *bool  `mapstructure:"add_user_id_to_req_param"`
	AddCompanyIdToReqParam *bool  `mapstructure:"add_company_id_to_req_param"`
}

// 自定义配置
type Method struct {
	ModelName              string `mapstructure:"model_name" validate:"required"`
	DocTag                 string `mapstructure:"doc_tag" validate:"required"`
	LogicPath              string `mapstructure:"logic_path" validate:"required"`
	RepoName               string `mapstructure:"repo_name" validate:"required"`
	DependencyName         string `mapstructure:"dependency_name" validate:"required"`
	AddIpToReqParam        *bool  `mapstructure:"add_ip_to_req_param" validate:"required"`
	AddUserIdToReqParam    *bool  `mapstructure:"add_user_id_to_req_param" validate:"required"`
	AddCompanyIdToReqParam *bool  `mapstructure:"add_company_id_to_req_param" validate:"required"`
	MethodName             string `mapstructure:"method_name" validate:"required"`
	DocDesc                string `mapstructure:"doc_desc" validate:"required"`
	DocUrl                 string `mapstructure:"doc_url" validate:"required"`
	DocUrlMethod           string `mapstructure:"doc_url_method" validate:"required"`
}
