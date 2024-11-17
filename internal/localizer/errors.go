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

func CreateLinkNameAlreadyExistError(linkName string) error {
	return LocalizedError{
		Config: &i18n.LocalizeConfig{
			MessageID: "error.linkname_already_exist",
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

func CreateNoSuchTagError(linkName, tag string) error {
	return LocalizedError{
		Config: &i18n.LocalizeConfig{
			MessageID: "error.no_such_tag",
			TemplateData: map[string]string{
				"LinkName": linkName,
				"Tag":      tag,
			},
		},
	}
}

func CreateTagAlreadyExistError(linkname, tagName string) error {
	return LocalizedError{
		Config: &i18n.LocalizeConfig{
			MessageID: "error.tag_already_exist",
			TemplateData: map[string]string{
				"Linkname": linkname,
				"Tag":      tagName,
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
