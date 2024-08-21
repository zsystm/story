//nolint:lll // Long lines are easier to read for flag descriptions.
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/storyprotocol/iliad/e2e/app"
	"github.com/storyprotocol/iliad/e2e/app/agent"
	"github.com/storyprotocol/iliad/e2e/app/key"
	"github.com/storyprotocol/iliad/e2e/types"
)

func bindDefFlags(flags *pflag.FlagSet, cfg *app.DefinitionConfig) {
	bindPromFlags(flags, &cfg.AgentSecrets)
	flags.StringVarP(&cfg.ManifestFile, "manifest-file", "f", cfg.ManifestFile, "path to manifest file")
	flags.StringVar(&cfg.InfraProvider, "infra", cfg.InfraProvider, "infrastructure provider: docker, vmcompose")
	flags.StringVar(&cfg.InfraDataFile, "infra-file", cfg.InfraDataFile, "infrastructure data file (not required for docker provider)")
	flags.StringVar(&cfg.DeployKeyFile, "deploy-key", cfg.DeployKeyFile, "path to deploy private key file")
	flags.StringVar(&cfg.FireAPIKey, "fireblocks-api-key", cfg.FireAPIKey, "FireBlocks api key")
	flags.StringVar(&cfg.FireKeyPath, "fireblocks-key-path", cfg.FireKeyPath, "FireBlocks RSA private key path")
	flags.StringVar(&cfg.IliadImgTag, "iliad-image-tag", cfg.IliadImgTag, "Iliad docker images tag (iliad). Defaults to working dir git commit.")
	flags.StringToStringVar(&cfg.RPCOverrides, "rpc-overrides", cfg.RPCOverrides, "Public chain rpc overrides: '<chain1>=<url1>,<url2>'")
	flags.StringVar(&cfg.TracingEndpoint, "tracing-endpoint", cfg.TracingEndpoint, "Tracing endpoint")
	flags.StringVar(&cfg.TracingHeaders, "tracing-headers", cfg.TracingHeaders, "Tracing headers")
}

func bindE2EFlags(flags *pflag.FlagSet, cfg *app.E2ETestConfig) {
	flags.BoolVar(&cfg.Preserve, "preserve", cfg.Preserve, "preserve infrastructure after test")
}

func bindPromFlags(flags *pflag.FlagSet, cfg *agent.Secrets) {
	flags.StringVar(&cfg.URL, "prom-url", cfg.URL, "prometheus url (only required if prometheus==true)")
	flags.StringVar(&cfg.User, "prom-user", cfg.User, "prometheus user")
	flags.StringVar(&cfg.Pass, "prom-password", cfg.Pass, "prometheus password")
}

func bindUpgradeFlags(flags *pflag.FlagSet, cfg *types.UpgradeConfig) {
	flags.StringVar(&cfg.ServiceRegexp, "services", cfg.ServiceRegexp, "Regexp applied to services per VM. Any match results in the VM being upgraded (all services on that VM are upgraded, not only matching services)")
}

func bindCreate3DeployFlags(flags *pflag.FlagSet, cfg *app.Create3DeployConfig) {
	flags.Uint64Var(&cfg.ChainID, "chain-id", cfg.ChainID, "chain id of the chain to deploy to")
}

func bindKeyCreateFlags(cmd *cobra.Command, cfg *key.UploadConfig) {
	cmd.Flags().StringVar(&cfg.Name, "name", cfg.Name, "key name: either node name or eoa account type")
	cmd.Flags().StringVar((*string)(&cfg.Type), "type", string(cfg.Type), "key type: validator, p2p_execution, p2p_consensus, eoa")

	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("type")
}
