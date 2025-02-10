package pythonwalk

import "docwiz/internal/badge"

var (
	shieldDjango = &badge.ShieldBadge{
		ID:        "Django",
		Label:     "django",
		Color:     "#092E20",
		Style:     badge.ShieldStyleDefault,
		Logo:      "django",
		LogoColor: "white",
		Href:      "https://www.djangoproject.com/",
	}

	shieldFlask = &badge.ShieldBadge{
		ID:        "Flask",
		Label:     "flask",
		Color:     "#000000",
		Style:     badge.ShieldStyleDefault,
		Logo:      "flask",
		LogoColor: "white",
		Href:      "https://flask.palletsprojects.com/",
	}

	shieldFastAPI = &badge.ShieldBadge{
		ID:        "FastAPI",
		Label:     "FastAPI",
		Color:     "#005571",
		Style:     badge.ShieldStyleDefault,
		Logo:      "fastapi",
		LogoColor: "white",
		Href:      "https://fastapi.tiangolo.com/",
	}
)
