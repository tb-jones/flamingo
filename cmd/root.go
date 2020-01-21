package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ToolName = "flamingo"
var Version = "0.0.2"
var BuildDate = "2020-01-20"

type flamingoParameters struct {
	Verbose        bool
	IgnoreFailures bool
	SNMPPorts      string
	SSHPorts       string
	SSHHostKey     string
	LDAPPorts      string
	LDAPSPorts     string
	HTTPPorts      string
	HTTPSPorts     string
	HTTPBasicRealm string
	TLSCertFile    string
	TLSCertData    string
	TLSKeyFile     string
	TLSKeyData     string
	TLSName        string
	TLSOrgName     string
	Protocols      string
}

var params = &flamingoParameters{}

var rootCmd = &cobra.Command{
	Use:   ToolName,
	Short: fmt.Sprintf("%s captures inbound credentials", ToolName),
	Long:  fmt.Sprintf(`flamingo v%s [%s]`, Version, BuildDate),
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		startCapture(cmd, args)
	},
}

// Execute is the main entry point for this tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// General options
	rootCmd.PersistentFlags().BoolVarP(&params.Verbose, "verbose", "v", false, "Display verbose output")
	rootCmd.PersistentFlags().BoolVarP(&params.IgnoreFailures, "ignore", "I", false, "Ignore individual listener failures")

	rootCmd.Flags().StringVarP(&params.Protocols, "protocols", "", "ssh,snmp,ldap,http", "Specify a comma-separated list of protocols")

	// SNMP parameters
	rootCmd.Flags().StringVarP(&params.SNMPPorts, "snmp-ports", "", "161", "The list of UDP ports to listen on for SNMP")

	// SSH parameters
	rootCmd.Flags().StringVarP(&params.SSHPorts, "ssh-ports", "", "22", "The list of TCP ports to listen on for SSH")
	rootCmd.Flags().StringVarP(&params.SSHHostKey, "ssh-host-key", "", "", "An optional path to a SSH host key on disk")

	// LDAP(S) parameters
	rootCmd.Flags().StringVarP(&params.LDAPPorts, "ldap-ports", "", "389", "The list of TCP ports to listen on for LDAP")
	rootCmd.Flags().StringVarP(&params.LDAPSPorts, "ldaps-ports", "", "636", "The list of TCP ports to listen on for LDAPS")

	// HTTP(S) parameters
	rootCmd.Flags().StringVarP(&params.HTTPPorts, "http-ports", "", "80", "The list of TCP ports to listen on for HTTP")
	rootCmd.Flags().StringVarP(&params.HTTPSPorts, "https-ports", "", "443", "The list of TCP ports to listen on for HTTPS")
	rootCmd.Flags().StringVarP(&params.HTTPBasicRealm, "http-realm", "", "Administration", "The authentication realm to present")

	rootCmd.Flags().StringVarP(&params.TLSCertFile, "tls-cert", "", "", "An optional x509 certificate for TLS listeners")
	rootCmd.Flags().StringVarP(&params.TLSKeyFile, "tls-key", "", "", "An optional x509 key for TLS listeners")
	rootCmd.Flags().StringVarP(&params.TLSName, "tls-name", "", "localhost", "A server name to use with TLS listeners")
	rootCmd.Flags().StringVarP(&params.TLSOrgName, "tls-org", "", "Flamingo Feed, Inc.", "An organization to use for self-signed certificates")

}
