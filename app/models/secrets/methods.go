package secrets

import "os"

func (e *Env) LoadSecrets() {
	e.secrets["GCP_PROJECT_ID"] = os.Getenv("GCP_PROJECT_ID")
}

func (e *Env) AddSecret(key, value string) {
	e.secrets[key] = value
}
