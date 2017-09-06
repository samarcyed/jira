package jiracmd

import (
	"github.com/coryb/figtree"
	"github.com/coryb/oreo"
	"gopkg.in/Netflix-Skunkworks/go-jira.v1"
	"gopkg.in/Netflix-Skunkworks/go-jira.v1/jiracli"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func CmdIssueLinkTypesRegistry() *jiracli.CommandRegistryEntry {
	opts := jiracli.CommonOptions{
		Template: figtree.NewStringOption("issuelinktypes"),
	}

	return &jiracli.CommandRegistryEntry{
		"Show the issue link types",
		func(fig *figtree.FigTree, cmd *kingpin.CmdClause) error {
			jiracli.LoadConfigs(cmd, fig, &opts)
			return CmdIssueLinkTypesUsage(cmd, &opts)
		},
		func(o *oreo.Client, globals *jiracli.GlobalOptions) error {
			return CmdIssueLinkTypes(o, globals, &opts)
		},
	}
}

func CmdIssueLinkTypesUsage(cmd *kingpin.CmdClause, opts *jiracli.CommonOptions) error {
	jiracli.TemplateUsage(cmd, opts)
	return nil
}

// CmdIssueLinkTypes will get issue link type data and send to "issuelinktypes" template
func CmdIssueLinkTypes(o *oreo.Client, globals *jiracli.GlobalOptions, opts *jiracli.CommonOptions) error {
	data, err := jira.GetIssueLinkTypes(o, globals.Endpoint.Value)
	if err != nil {
		return err
	}
	return jiracli.RunTemplate(opts.Template.Value, data, nil)
}
