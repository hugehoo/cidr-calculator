package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"os"
	"strings"
)

var cidrCmd = &cobra.Command{
	Use:   "cidr",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples and usage of using your command. 

For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	//Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Input CIDR :")
		reader := bufio.NewReader(os.Stdin)
		calculateCidr(reader)
	},
}

func init() {
	rootCmd.AddCommand(cidrCmd)
}

func calculateCidr(reader *bufio.Reader) {
	cidr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error input", err)
		return
	}

	cidr = strings.TrimSpace(cidr)
	startIP, endIP, err := cidrToIPRange(cidr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Start IP: %s\n", startIP)
	fmt.Printf("End IP: %s\n", endIP)
}

func cidrToIPRange(cidr string) (string, string, error) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", "", err
	}

	// Get the network address and the broadcast address
	startIP := ip.Mask(ipNet.Mask)
	endIP := make(net.IP, len(startIP))
	copy(endIP, startIP)
	for i := range startIP {
		endIP[i] |= ^ipNet.Mask[i]
	}

	// Convert IP addresses to strings
	startIPStr := startIP.String()
	endIPStr := endIP.String()

	return startIPStr, endIPStr, nil
}
