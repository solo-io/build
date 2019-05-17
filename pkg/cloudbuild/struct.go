package cloudbuild

type Cloudbuild struct {
	Steps   []*Step   `json:"steps,omitempty"`
	Secrets []*Secret `json:"secrets,omitempty"`
	Timeout string    `json:"timeout,omitempty"`
}

type Step struct {
	Name       string    `json:"name"`
	Args       []string  `json:"args,omitempty"`
	Entrypoint string    `json:"entrypoint,omitempty"`
	Env        []string  `json:"env,omitempty"`
	Id         string    `json:"id,omitempty"`
	SecretEnv  []string  `json:"secretEnv,omitempty"`
	Volumes    []*Volume `json:"volumes,omitempty"`
}

type Volume struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

type Secret struct {
	KmsKeyName string            `json:"kmsKeyName,omitempty"`
	SecretEnv  map[string]string `json:"secretEnv,omitempty"`
}
