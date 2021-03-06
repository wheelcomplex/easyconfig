package examplelib

import "gopkg.in/hlandau/easyconfig.v1/cflag"

var g = cflag.NewGroup(nil, "Server Options")
var bindFlag = cflag.String(g, "bind", ":53", "Address to bind to (e.g. 0.0.0.0:53)")
var publicKeyFlag = cflag.String(g, "publicKey", "", "Path to the DNSKEY KSK public key file")
var privateKeyFlag = cflag.String(g, "privateKey", "", "Path to the KSK's corresponding private key file")
var zonePublicKeyFlag = cflag.String(g, "zonePublicKey", "", "Path to the DNSKEY ZSK public key file; if one is not specified, a temporary one is generated on startup and used only for the duration of that process")
var zonePrivateKeyFlag = cflag.String(g, "zonePrivateKey", "", "Path to the ZSK's corresponding private key file")
var namecoinRPCUsernameFlag = cflag.String(g, "namecoinRPCUsername", "", "Namecoin RPC username")
var namecoinRPCPasswordFlag = cflag.String(g, "namecoinRPCPassword", "", "Namecoin RPC password")
var namecoinRPCAddressFlag = cflag.String(g, "namecoinRPCAddress", "localhost:8336", "Namecoin RPC server address")
var cacheMaxEntriesFlag = cflag.Int(g, "cacheMaxEntries", 100, "Maximum name cache entries")
var selfNameFlag = cflag.String(g, "selfName", "", "The FQDN of this nameserver; if empty, a psuedo-hostname is generated")
var selfIPFlag = cflag.String(g, "selfIP", "127.127.127.127", "The canonical IP address for this service")
var httpListenAddrFlag = cflag.String(g, "httpListenAddr", "", "Address for the webserver to listen at (default: disabled)")
var canonicalSuffixFlag = cflag.String(g, "canonicalSuffix", "bit", "Suffix to advertise via HTTP")
var canonicalNameserversFlag = cflag.String(g, "canonicalNameservers", "", "Comma-separated list of nameservers to use for NS records; if blank, selfName (or autogenerated psuedo-hostname) is used")
var hostmasterFlag = cflag.String(g, "hostmaster", "", "Hostmaster e. mail address")
var vanityIPs = cflag.String(g, "vanityIPs", "", "Comma separated list of IP addresses to place in A/AAAA records at the zone apex (default: don't add any records)")
var doStuff = cflag.Bool(g, "doStuff", false, "Do stuff")
