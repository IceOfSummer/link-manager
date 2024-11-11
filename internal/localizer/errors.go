package localizer

import "github.com/nicksnyder/go-i18n/v2/i18n"

func CreateNoSuchLinkError(linkName string) error {
	return LocalizedError{
		Config: &i18n.LocalizeConfig{
			MessageID: "error.no_such_link",
			TemplateData: map[string]string{
				"LinkName": linkName,
			},
		},
	}
}

func CreateError(messageId string) error {
	return LocalizedError{
		Config: &i18n.LocalizeConfig{
			MessageID: messageId,
		},
	}
}

func CreateNoSuchLinkValueError(linkName, tag string) error {
	return LocalizedError{
		Config: &i18n.LocalizeConfig{
			MessageID: "error.no_such_link_value",
			TemplateData: map[string]string{
				"LinkName": linkName,
				"Tag":      tag,
			},
		},
	}
}

// CreateNoSuchBindError
// Deprecated, use [CreateError] instead.
func CreateNoSuchBindError() error {
	return LocalizedError{
		Config: &i18n.LocalizeConfig{
			MessageID: ErrorNoSuchBind,
		},
	}
}
