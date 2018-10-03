package amp

import "html/template"

type AMPi struct {
	AMPImg template.HTML `json:"ampimg"`
}

func AMPI() AMPi {
	amp := AMPi{AMPImg: template.HTML(`<a class="cglc" href="/coin/{{g}}"><amp-img width="32" height="32" alt="{{n}}" src="{{i}}"></amp-img><noscript><img src="{{i}}" width="32" height="32" alt="{{n}}"></noscript></a>`)}
	return amp
}

type AMPb struct {
	BlockHeight   template.HTML `json:"blockheight"`
	BlockHash     template.HTML `json:"blockhash"`
	Tx            template.HTML `json:"tx"`
	Addr          template.HTML `json:"addr"`
	NextBlockHash template.HTML `json:"nbh"`
	PrevBlockHash template.HTML `json:"pbh"`
}

func AMPB() AMPb {
	amp := AMPb{
		BlockHeight:   template.HTML(`<a href="/block/{{height}}">{{height}}</a>`),
		BlockHash:     template.HTML(`<a href="/blockhash/{{hash}}">{{hash}}</a>`),
		Tx:            template.HTML(`<a href="/tx/{{.}}">{{.}}</a>`),
		Addr:          template.HTML(`<a href="/addr/{{.}}">{{.}}</a>`),
		NextBlockHash: template.HTML(`<a href="/blockhash/{{nextblockhash}}">{{nextblockhash}}</a>`),
		PrevBlockHash: template.HTML(`<a href="/blockhash/{{previousblockhash}}">{{previousblockhash}}</a>`),
	}
	return amp
}
