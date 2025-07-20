package middleware_helper

import (
	"encoding/json"
	"regexp"
	"strings"
)

var sensitiveKeys = []string{"password", "email", "card_number", "pan"}

// MaskSensitiveData tries to sanitize sensitive fields in JSON string
func MaskSensitiveData(rawBody []byte) string {

	if rawBody != nil {
		var data map[string]interface{}
		if err := json.Unmarshal(rawBody, &data); err != nil {
			return string(rawBody) // fallback: return original
		}

		for _, key := range sensitiveKeys {
			for k := range data {
				if strings.EqualFold(k, key) {
					data[k] = maskValue(k, data[k])
				}
			}
		}

		masked, err := json.Marshal(data)
		if err != nil {
			return string(rawBody)
		}

		return string(masked)
	} else {
		return ""
	}
}

func maskValue(key string, value interface{}) string {
	str, ok := value.(string)
	if !ok {
		return "***"
	}

	switch strings.ToLower(key) {
	case "password":
		return strings.Repeat("*", len(str))
	case "email":
		return maskEmail(str)
	case "card_number", "pan":
		return maskCardNumber(str)
	default:
		return "***"
	}
}

func maskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "***"
	}
	return parts[0][:1] + "*****@" + parts[1]
}

func maskCardNumber(card string) string {
	if !regexp.MustCompile(`^\d+$`).MatchString(card) {
		return "***"
	}
	if len(card) <= 4 {
		return strings.Repeat("*", len(card))
	}
	return strings.Repeat("*", len(card)-4) + card[len(card)-4:]
}
