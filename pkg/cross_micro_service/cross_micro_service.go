package cross_micro_service

type CrossMicroServiceConfig struct {
	HarmonyAuthorizationToken string `env:"HARMONY_AUTHORIZATION_TOKEN"`
	HarmonyBaseUrl            string `env:"HARMONY_BASE_URL"`
}
