package sys

type Menus []MenuItem

type MenuItem struct {
	Path     string      `json:"path"`
	Name     string      `json:"name"`
	Meta     MenuMeta    `json:"meta"`
	Children []MenuChild `json:"children"`
}

type MenuChild struct {
	Path string   `json:"path"`
	Name string   `json:"name"`
	Meta MenuMeta `json:"meta"`
}

type MenuMeta struct {
	Locale       string   `json:"locale"`
	RequiresAuth bool     `json:"requires_auth"`
	Icon         string   `json:"icon,omitempty"`
	Order        int      `json:"order,omitempty"`
	Roles        []string `json:"roles,omitempty"`
	HideInMenu   bool     `json:"hide_in_menu,omitempty"`
	ImgPath      string   `json:"img,omitempty"`
	SvgIcon      string   `json:"svg,omitempty"`
	TopLevel     bool     `json:"top_level,omitempty"`
	Placeholder  bool     `json:"placeholder,omitempty"`
}
