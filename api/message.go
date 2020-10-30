package rocket

type MessageService service

type MessagePayload struct {
	Channel     string               `json:"channel" yaml:"channel"`
	Text        string               `json:"text,omitempty" yaml:"text,omitempty"`
	Alias       string               `json:"alias,omitempty" yaml:"alias,omitempty"`
	Emoji       string               `json:"emoji,omitempty" yaml:"emoji,omitempty"`
	Avatar      string               `json:"avatar,omitempty" yaml:"avatar,omitempty"`
	Attachments []MessageAttachement `json:"attachments,omitempty" yaml:"attachments,omitempty"`
}

type MessageAttachement struct {
	Color             string            `json:"color,omitempty" yaml:"color,omitempty"`
	Text              string            `json:"text,omitempty" yaml:"text,omitempty"`
	AuthorName        string            `json:"author_name,omitempty" yaml:"author_name,omitempty"`
	AuthorLink        string            `json:"author_link,omitempty" yaml:"author_link,omitempty"`
	AuthorIcon        string            `json:"author_icon,omitempty" yaml:"author_icon,omitempty"`
	Title             string            `json:"title,omitempty" yaml:"title,omitempty"`
	TitleLink         string            `json:"title_link,omitempty" yaml:"title_link,omitempty"`
	TitleLinkDownload bool              `json:"title_link_download,omitempty" yaml:"title_link_download,omitempty"`
	ImageUrl          string            `json:"image_url,omitempty" yaml:"image_url,omitempty"`
	AudioUrl          string            `json:"audio_url,omitempty" yaml:"audio_url,omitempty"`
	VideoUrl          string            `json:"video_url,omitempty" yaml:"video_url,omitempty"`
	Fields            []AttachmentField `json:"fields" yaml:"fields"`
}

type AttachmentField struct {
	Short bool   `json:"short,omitempty" yaml:"short,omitempty"`
	Title string `json:"title" yaml:"title"`
	Value string `json:"value" yaml:"value"`
}
