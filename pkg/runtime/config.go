package runtime

type Config struct {
	LogStreamCreation     bool `yaml:"log_stream_creation"`
	LogPushRequest        bool `yaml:"log_push_request"`
	LogPushRequestStreams bool `yaml:"log_push_request_streams"`

	// LimitedLogPushErrors is to be implemented and will allow logging push failures at a controlled pace.
	LimitedLogPushErrors bool `yaml:"limited_log_push_errors"`
}

var EmptyConfig = &Config{}

// TenantConfigProvider serves a tenant or default config.
type TenantConfigProvider interface {
	TenantConfig(userID string) *Config
}

// TenantConfigs periodically fetch a set of per-user configs, and provides convenience
// functions for fetching the correct value.
type TenantConfigs struct {
	TenantConfigProvider
}

// DefaultTenantConfigs creates and returns a new TenantConfigs with the defaults populated.
func DefaultTenantConfigs() *TenantConfigs {
	return &TenantConfigs{
		TenantConfigProvider: nil,
	}
}

// NewTenantConfig makes a new TenantConfigs
func NewTenantConfigs(configProvider TenantConfigProvider) (*TenantConfigs, error) {
	return &TenantConfigs{
		TenantConfigProvider: configProvider,
	}, nil
}

func (o *TenantConfigs) getOverridesForUser(userID string) *Config {
	if o.TenantConfigProvider != nil {
		l := o.TenantConfigProvider.TenantConfig(userID)
		if l != nil {
			return l
		}
	}
	return EmptyConfig
}

func (o *TenantConfigs) LogStreamCreation(userID string) bool {
	return o.getOverridesForUser(userID).LogStreamCreation
}

func (o *TenantConfigs) LogPushRequest(userID string) bool {
	return o.getOverridesForUser(userID).LogPushRequest
}

func (o *TenantConfigs) LogPushRequestStreams(userID string) bool {
	return o.getOverridesForUser(userID).LogPushRequestStreams
}

func (o *TenantConfigs) LimitedLogPushErrors(userID string) bool {
	return o.getOverridesForUser(userID).LimitedLogPushErrors
}
