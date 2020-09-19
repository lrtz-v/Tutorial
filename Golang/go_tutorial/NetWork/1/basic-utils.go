package basicutils

import (
	"io/ioutil"
	"net"
)

// ParseIP parse ip addr
func ParseIP(ip string) net.IP {
	return net.ParseIP(ip)
}

// DefaultMask returns the default IP mask for the IP address ip (ipv4 only)
func DefaultMask(ip string) net.IPMask {
	addr := net.ParseIP(ip)
	return addr.DefaultMask()
}

// Mask returns the result of masking the IP address ip with mask.
func Mask(ip string) net.IP {
	addr := net.ParseIP(ip)
	return addr.Mask(addr.DefaultMask())
}

// LookupPort looks up the port for the given network and service.
func LookupPort(networkType, service string) (int, error) {
	return net.LookupPort(networkType, service)
}

// ResolveIPAddr returns an address of IP end point.
func ResolveIPAddr(network, name string) (*net.IPAddr, error) {
	return net.ResolveIPAddr(network, name)
}

// LookupHost looks up the given host using the local resolver.
func LookupHost(host string) (addrs []string, err error) {
	return net.LookupHost(host)
}

// GetHeadInfo return service content
func GetHeadInfo(service string) ([]byte, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}
	return result, nil
}
