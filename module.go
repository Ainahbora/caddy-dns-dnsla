package template

import (
	"fmt"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsla "github.com/r6c/dnsla"
)

type Provider struct {
	*libdnsla.Provider
}

func init() {
	caddy.RegisterModule(Provider{})
}

func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.dnsla",
		New: func() caddy.Module { return &Provider{new(libdnsla.Provider)} },
	}
}

func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIID = caddy.NewReplacer().ReplaceAll(p.Provider.APIID, "")
	p.Provider.APISecret = caddy.NewReplacer().ReplaceAll(p.Provider.APISecret, "")
	return fmt.Errorf("TODO: not implemented")
}

func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_id":
				if p.Provider.APIID != "" {
					return d.Err("API ID already set")
				}
				if d.NextArg() {
					p.Provider.APIID = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "api_secret":
				if p.Provider.APISecret != "" {
					return d.Err("API Secret already set")
				}
				if d.NextArg() {
					p.Provider.APISecret = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIID == "" {
		return d.Err("missing API ID")
	}
	if p.Provider.APISecret == "" {
		return d.Err("missing API Secret")
	}
	return nil
}

var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
