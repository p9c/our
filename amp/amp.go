package amp

import "html/template"

type AMP struct {
	AMPImg template.HTML `json:"ampimg"`
}

func AMPS() AMP {
	amp := AMP{AMPImg: template.HTML(`<a class="cglc" href="/coin/{{g}}"><amp-img width="32" height="32" alt="{{n}}" src="{{i}}"></amp-img><noscript><img src="{{i}}" width="32" height="32" alt="{{n}}"></noscript></a>`)}
	return amp
}
