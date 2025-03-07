package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
)

var version = "1.0.0"

// Prints a detailed help page
func printHelp() {
	helpText := `
CIDR Calculation Tool - Helps with understanding and calculating CIDR ranges.

Usage:
  1. Calculate CIDR details:
     $ ./cidr-tool calculate 192.168.1.0/24
     
  2. Split a CIDR block into smaller subnets:
     $ ./cidr-tool split 192.168.1.0/24 26
     
  3. Merge multiple CIDR blocks:
     $ ./cidr-tool merge 192.168.1.0/24 192.168.2.0/24
  
CIDR Basics:
  - CIDR notation defines IP ranges (e.g., 192.168.1.0/24).
  - The prefix length (/24) determines how many addresses are available.
  - The subnet mask controls how the IP space is divided.
  - Network Address: First IP in the range.
  - Broadcast Address: Last IP in the range.
  - Usable IPs: IPs available for devices (excluding network/broadcast IPs).
  
Examples:
  - /24 = 256 total IPs (254 usable)
  - /26 = 64 total IPs (62 usable)
  - /16 = 65,536 total IPs (65,534 usable)
`
	fmt.Println(helpText)
}

// Calculates CIDR details
func calculateCIDR(cidr string) {
	_, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Invalid CIDR:", err)
		return
	}

	networkIP := ipv4Net.IP
	maskSize, bits := ipv4Net.Mask.Size()
	totalIPs := int(math.Pow(2, float64(bits-maskSize)))
	usableIPs := totalIPs - 2

	broadcastIP := make(net.IP, len(networkIP))
	copy(broadcastIP, networkIP)
	for i := range broadcastIP {
		broadcastIP[i] |= ^ipv4Net.Mask[i]
	}

	firstUsableIP := make(net.IP, len(networkIP))
	copy(firstUsableIP, networkIP)
	firstUsableIP[3]++

	lastUsableIP := make(net.IP, len(broadcastIP))
	copy(lastUsableIP, broadcastIP)
	lastUsableIP[3]--

	fmt.Println("CIDR:", cidr)
	fmt.Println("Network Address:", networkIP)
	fmt.Println("Subnet Mask:", ipv4Net.Mask)
	fmt.Println("Total Addresses:", totalIPs)
	fmt.Println("Usable Addresses:", usableIPs)
	fmt.Println("First Usable IP:", firstUsableIP)
	fmt.Println("Last Usable IP:", lastUsableIP)
	fmt.Println("Broadcast Address:", broadcastIP)
}

// Splits a CIDR block into smaller subnets
func splitCIDR(cidr string, newPrefix int) {
	_, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Invalid CIDR:", err)
		return
	}

	currentPrefix, totalBits := ipv4Net.Mask.Size()
	if newPrefix <= currentPrefix || newPrefix > totalBits {
		fmt.Println("Invalid new subnet prefix!")
		return
	}

	subnetCount := int(math.Pow(2, float64(newPrefix-currentPrefix)))
	fmt.Printf("Splitting %s into %d /%d subnets:\n", cidr, subnetCount, newPrefix)

	networkIP := ipv4Net.IP
	for i := 0; i < subnetCount; i++ {
		subnet := fmt.Sprintf("%s/%d", networkIP, newPrefix)
		fmt.Println(subnet)

		// Move to the next subnet
		networkIP = nextSubnet(networkIP, newPrefix)
	}
}

// Calculates the next subnet address
func nextSubnet(ip net.IP, newPrefix int) net.IP {
	ipInt := ipToInt(ip)
	increment := 1 << (32 - newPrefix)
	nextIPInt := ipInt + uint32(increment)
	return intToIP(nextIPInt)
}

// Converts an IP to an integer
func ipToInt(ip net.IP) uint32 {
	return uint32(ip[12])<<24 | uint32(ip[13])<<16 | uint32(ip[14])<<8 | uint32(ip[15])
}

// Converts an integer to an IP
func intToIP(ipInt uint32) net.IP {
	return net.IPv4(byte(ipInt>>24), byte(ipInt>>16), byte(ipInt>>8), byte(ipInt))
}

// Merges multiple CIDR blocks
func mergeCIDRs(cidrs []string) {
	if len(cidrs) < 2 {
		fmt.Println("Provide at least two CIDR blocks to merge.")
		return
	}

	// Convert CIDRs to IP ranges
	ipRanges := []struct {
		startIP net.IP
		endIP   net.IP
	}{}

	for _, cidr := range cidrs {
		_, ipv4Net, err := net.ParseCIDR(cidr)
		if err != nil {
			fmt.Println("Invalid CIDR:", err)
			return
		}
		startIP := ipv4Net.IP
		endIP := make(net.IP, len(startIP))
		copy(endIP, startIP)
		for i := range endIP {
			endIP[i] |= ^ipv4Net.Mask[i]
		}
		ipRanges = append(ipRanges, struct {
			startIP net.IP
			endIP   net.IP
		}{startIP, endIP})
	}

	// Find the smallest range covering all CIDRs
	startIP := ipRanges[0].startIP
	endIP := ipRanges[0].endIP
	for _, r := range ipRanges[1:] {
		if bytesCompare(r.startIP, startIP) < 0 {
			startIP = r.startIP
		}
		if bytesCompare(r.endIP, endIP) > 0 {
			endIP = r.endIP
		}
	}

	fmt.Printf("Merged CIDR covers range: %s - %s\n", startIP, endIP)
}

// Compares two byte slices (IP addresses)
func bytesCompare(a, b net.IP) int {
	return strings.Compare(a.String(), b.String())
}

// Validates if an IP belongs to a CIDR block
func validateIPInCIDR(ipStr, cidr string) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		fmt.Println("Invalid IP address")
		return
	}

	_, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Invalid CIDR:", err)
		return
	}

	if ipv4Net.Contains(ip) {
		fmt.Printf("IP %s is within the CIDR block %s\n", ipStr, cidr)
	} else {
		fmt.Printf("IP %s is NOT within the CIDR block %s\n", ipStr, cidr)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./cidr-tool [calculate|split|merge|validate|help|version] <CIDR> [new prefix]")
		return
	}

	command := os.Args[1]
	switch command {
	case "help":
		printHelp()
	case "version":
		fmt.Println("CIDR Calculation Tool Version:", version)
	case "calculate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ./cidr-tool calculate <CIDR>")
			return
		}
		calculateCIDR(os.Args[2])
	case "split":
		if len(os.Args) < 4 {
			fmt.Println("Usage: ./cidr-tool split <CIDR> <new prefix>")
			return
		}
		newPrefix, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Invalid prefix:", err)
			return
		}
		splitCIDR(os.Args[2], newPrefix)
	case "merge":
		if len(os.Args) < 4 {
			fmt.Println("Usage: ./cidr-tool merge <CIDR1> <CIDR2> [CIDR3] ...")
			return
		}
		mergeCIDRs(os.Args[2:])
	case "validate":
		if len(os.Args) < 4 {
			fmt.Println("Usage: ./cidr-tool validate <IP> <CIDR>")
			return
		}
		validateIPInCIDR(os.Args[2], os.Args[3])
	default:
		fmt.Println("Invalid command. Use './cidr-tool help' for usage details.")
	}
}
