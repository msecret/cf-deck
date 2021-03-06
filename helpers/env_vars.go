package helpers

var (
	// ClientIDEnvVar is the environment variable key that represents the registered
	// Client ID for this web app.
	ClientIDEnvVar = "CONSOLE_CLIENT_ID"
	// ClientSecretEnvVar is the environment variable key that represents the
	// Client Secret associated with the registered Client ID for this web app.
	ClientSecretEnvVar = "CONSOLE_CLIENT_SECRET"
	// HostnameEnvVar is the environment variable key that represents the hostname
	// of this web app.
	HostnameEnvVar = "CONSOLE_HOSTNAME"
	// LoginURLEnvVar is the environment variable key that represents the
	// base login URL endpoint that this app should use to authenticate users.
	LoginURLEnvVar = "CONSOLE_LOGIN_URL"
	// UAAURLEnvVar is the environment variable key that represents the
	//  base uaa URL endpoint that this app should use to get tokens.
	UAAURLEnvVar = "CONSOLE_UAA_URL"
	// APIURLEnvVar is the environment variable key that represents the
	// base api URL endpoint that this app should use to access cloud foundry data.
	APIURLEnvVar = "CONSOLE_API_URL"
	// LogURLEnvVar is the environment variable key that represents the
	// endpoint to the loggregator.
	LogURLEnvVar = "CONSOLE_LOG_URL"
	// PProfEnabledEnvVar is the environment variable key that represents if the pprof routes
	// should be enabled. If no value is specified, it is assumed to be false.
	PProfEnabledEnvVar = "PPROF_ENABLED"
	// BuildInfoEnvVar is the environment variable key that represents the particular build number
	BuildInfoEnvVar = "BUILD_INFO"
)

// EnvVars holds all the environment variable values that a non-test server should have.
type EnvVars struct {
	ClientID     string
	ClientSecret string
	Hostname     string
	LoginURL     string
	UAAURL       string
	APIURL       string
	LogURL       string
	PProfEnabled string
	BuildInfo    string
}
