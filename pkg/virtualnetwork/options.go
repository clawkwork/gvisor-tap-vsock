package virtualnetwork

import (
	"github.com/containers/gvisor-tap-vsock/pkg/services/dns"
	"github.com/containers/gvisor-tap-vsock/pkg/services/forwarder"
)

// Option configures optional, runtime-only behaviour on a VirtualNetwork at
// construction time. Options are separate from types.Configuration because
// they carry live callbacks rather than serializable settings.
type Option func(*options)

type options struct {
	tcpFilter   forwarder.AddressFilter
	udpFilter   forwarder.AddressFilter
	icmpFilter  forwarder.AddressFilter
	dnsObserver dns.Observer
}

// WithTCPFilter gates each new outbound TCP connection through filter. A
// non-nil error from filter drops the connection.
func WithTCPFilter(filter forwarder.AddressFilter) Option {
	return func(o *options) { o.tcpFilter = filter }
}

// WithUDPFilter gates each new outbound UDP flow through filter. A non-nil
// error from filter drops the flow.
func WithUDPFilter(filter forwarder.AddressFilter) Option {
	return func(o *options) { o.udpFilter = filter }
}

// WithICMPFilter gates each outbound ICMP echo through filter (addr is the
// destination IP). A non-nil error from filter drops the packet.
func WithICMPFilter(filter forwarder.AddressFilter) Option {
	return func(o *options) { o.icmpFilter = filter }
}

// WithDNSObserver registers observer, called for each A answer the embedded
// DNS server returns to the guest.
func WithDNSObserver(observer dns.Observer) Option {
	return func(o *options) { o.dnsObserver = observer }
}
