package secrets

func (e *Env) GetSecret(key string) string {
	return e.secrets[key]
}
