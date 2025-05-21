package go_praxis

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
}

// notice:  input params is fixed
func getAuthHeaders(auth string) map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"charset":      "utf-8",
		SIGN_HEAD_NAME: auth,
	}
}
