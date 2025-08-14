package hexlet

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		return "en." + domain
	} else {
		return domain + "." + locale
	}
	
}
