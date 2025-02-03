package config

type Env string

const (
	Local      Env = "local"
	Demo       Env = "demo"
	Stage      Env = "stg"
	Preprod    Env = "preprod"
	QA         Env = "qa"
	Production Env = "prod"
	Test       Env = "test"
)

type Base struct {
	Env                Env     `config:"dd_env"`
	ServiceName        string  `config:"dd_service"`
	ServiceVersion     string  `config:"dd_version"`
	TraceAddress       string  `config:"trace_address"`
	MetricsAddress     string  `config:"metrics_address"`
	LogLevel           string  `config:"log_level"`
	SystemPort         int     `config:"system_port"`
	TraceMaxBatchCount int     `config:"trace_max_batch_count"`
	TraceSampleRate    float64 `config:"trace_sample_rate"`
	DisableLogSampling bool    `config:"disable_log_sampling"`
	DisableStackTraces bool    `config:"disable_stack_traces"`
	BindAddress        string  `config:"bind_address"`
}
