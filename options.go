package sentinel

import "time"

type Option func(*Sentinel)

func WithReferrerPolicy(directives ...ReferrerPolicyDirective) Option {
	return func(s *Sentinel) {
		s.middlewares = append(s.middlewares, newReferrerPolicy(directives...))
	}
}

func WithXContentTypeOptions() Option {
	return func(s *Sentinel) {
		s.middlewares = append(s.middlewares, xContentTypeOptions(DirectiveNoSniff))
	}
}

func WithStrictTransportSecurity(maxAge time.Duration, includeSubdomains bool, preload bool) Option {
	return func(s *Sentinel) {
		sts := newStrictTransportSecurity(int(maxAge.Seconds()), includeSubdomains, preload)
		s.middlewares = append(s.middlewares, sts)
	}
}

func WithXFrameOptions(directive XFrameOptionsDirective) Option {
	return func(s *Sentinel) {
		s.middlewares = append(s.middlewares, xFrameOptions(directive))
	}
}

func WithCrossOriginEmbedderPolicy(directive CrossOriginEmbedderPolicyDirective) Option {
	return func(s *Sentinel) {
		s.middlewares = append(s.middlewares, crossOriginEmbedderPolicy(directive))
	}
}

func WithCrossOriginOpenerPolicy(directive CrossOriginOpenerPolicyDirective) Option {
	return func(s *Sentinel) {
		s.middlewares = append(s.middlewares, crossOriginOpenerPolicy(directive))
	}
}

func WithCrossOriginResourcePolicy(directive CrossOriginResourcePolicyDirective) Option {
	return func(s *Sentinel) {
		s.middlewares = append(s.middlewares, crossOriginResourcePolicy(directive))
	}
}

func WithXDNSPrefetchControl(directive XDNSPrefetchControlDirective) Option {
	return func(s *Sentinel) {
		s.middlewares = append(s.middlewares, xDNSPrefetchControl(directive))
	}
}

func WithXSSProtection(enabled, block bool, reportURI string) Option {
	return func(s *Sentinel) {
		p := newXXSSProtection(enabled, block, reportURI)
		s.middlewares = append(s.middlewares, p)
	}
}
