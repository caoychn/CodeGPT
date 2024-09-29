package cmd

import (
	"errors"

	"github.com/caoychn/CodeGPT/core"
	"github.com/caoychn/CodeGPT/gemini"
	"github.com/caoychn/CodeGPT/openai"

	"github.com/spf13/viper"
)

func NewOpenAI() (*openai.Client, error) {
	return openai.New(
		openai.WithToken(viper.GetString("openai.api_key")),
		openai.WithModel(viper.GetString("openai.model")),
		openai.WithOrgID(viper.GetString("openai.org_id")),
		openai.WithProxyURL(viper.GetString("openai.proxy")),
		openai.WithSocksURL(viper.GetString("openai.socks")),
		openai.WithBaseURL(viper.GetString("openai.base_url")),
		openai.WithTimeout(viper.GetDuration("openai.timeout")),
		openai.WithMaxTokens(viper.GetInt("openai.max_tokens")),
		openai.WithTemperature(float32(viper.GetFloat64("openai.temperature"))),
		openai.WithProvider(core.Platform(viper.GetString("openai.provider"))),
		openai.WithSkipVerify(viper.GetBool("openai.skip_verify")),
		openai.WithHeaders(viper.GetStringSlice("openai.headers")),
		openai.WithAPIVersion(viper.GetString("openai.api_version")),
		openai.WithTopP(float32(viper.GetFloat64("openai.top_p"))),
		openai.WithFrequencyPenalty(float32(viper.GetFloat64("openai.frequency_penalty"))),
		openai.WithPresencePenalty(float32(viper.GetFloat64("openai.presence_penalty"))),
	)
}

// NewGemini returns a new Gemini client
func NewGemini() (*gemini.Client, error) {
	return gemini.New(
		gemini.WithToken(viper.GetString("openai.api_key")),
		gemini.WithModel(viper.GetString("openai.model")),
		gemini.WithMaxTokens(viper.GetInt32("openai.max_tokens")),
		gemini.WithTemperature(float32(viper.GetFloat64("openai.temperature"))),
		gemini.WithTopP(float32(viper.GetFloat64("openai.top_p"))),
	)
}

// GetClient returns the generative client based on the platform
func GetClient(p core.Platform) (core.Generative, error) {
	switch p {
	case core.Gemini:
		return NewGemini()
	case core.OpenAI, core.Azure:
		return NewOpenAI()
	}
	return nil, errors.New("invalid provider")
}
