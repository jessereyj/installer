// Package ipnet wraps net.IPNet to get CIDR serialization.
package ipnet

import (
	"encoding/json"
	"net"
	"reflect"
)

var nullString = "null"
var nullBytes = []byte(nullString)
var emptyIPNet = net.IPNet{}

// IPNet wraps net.IPNet to get CIDR serialization.
type IPNet struct {
	net.IPNet
}

// String returns a CIDR serialization of the subnet, or an empty
// string if the subnet is nil.
func (ipnet *IPNet) String() string {
	if ipnet == nil {
		return ""
	}
	return ipnet.IPNet.String()
}

// DeepCopyInto copies the receiver into out.  out must be non-nil.
func (ipnet *IPNet) DeepCopyInto(out *IPNet) {
	if ipnet == nil {
		*out = *new(IPNet)
	} else {
		*out = *ipnet
	}
	return
}

// DeepCopy copies the receiver, creating a new IPNet.
func (ipnet *IPNet) DeepCopy() *IPNet {
	if ipnet == nil {
		return nil
	}
	out := new(IPNet)
	ipnet.DeepCopyInto(out)
	return out
}

// MarshalJSON interface for an IPNet
func (ipnet IPNet) MarshalJSON() (data []byte, err error) {
	if reflect.DeepEqual(ipnet.IPNet, emptyIPNet) {
		return nullBytes, nil
	}

	return json.Marshal(ipnet.String())
}

// UnmarshalJSON interface for an IPNet
func (ipnet *IPNet) UnmarshalJSON(b []byte) (err error) {
	if string(b) == nullString {
		ipnet.IP = net.IP{}
		ipnet.Mask = net.IPMask{}
		return nil
	}

	var cidr string
	err = json.Unmarshal(b, &cidr)
	if err != nil {
		return err
	}

	ip, net, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	ipnet.IP = ip
	ipnet.Mask = net.Mask
	return nil
}
