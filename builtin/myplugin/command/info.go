package command

import (
	"strings"

	"github.com/DavidGamba/go-getoptions/option"
	"github.com/hashicorp/vagrant-plugin-sdk/component"
	plugincore "github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/docs"
	"github.com/hashicorp/vagrant-plugin-sdk/terminal"
)

// Info is a Command implementation for myplugin.
// It is a subcommand of myplugin
type Info struct {
	*Command
}

func (c *Info) ConfigSet(v interface{}) error {
	return nil
}

func (c *Info) CommandFunc() interface{} {
	return nil
}

func (c *Info) Config() (interface{}, error) {
	return &c.config, nil
}

func (c *Info) Documentation() (*docs.Documentation, error) {
	doc, err := docs.New(docs.FromConfig(&CommandConfig{}))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// ExecuteFunc implements component.Command
func (c *Info) ExecuteFunc(cliArgs []string) interface{} {
	return c.Execute
}

// CommandInfoFunc implements component.Command
func (c *Info) CommandInfoFunc() interface{} {
	return c.CommandInfo
}

func (c *Info) CommandInfo() (*component.CommandInfo, error) {
	return &component.CommandInfo{
		Name:     "info",
		Help:     c.Help(),
		Synopsis: c.Synopsis(),
		Flags:    c.Flags(),
	}, nil
}

func (c *Info) Synopsis() string {
	return "Output some project information!"
}

func (c *Info) Help() string {
	return "Output some project information!"
}

func (c *Info) Flags() []*option.Option {
	return []*option.Option{}
}

func (c *Info) Execute(trm terminal.UI, p plugincore.Project) int64 {
	mn, _ := p.MachineNames()
	trm.Output("\nMachines in this project")
	trm.Output(strings.Join(mn[:], "\n"))

	cwd, _ := p.CWD()
	datadir, _ := p.DataDir()
	vagrantfileName, _ := p.VagrantfileName()
	home, _ := p.Home()
	localDataPath, _ := p.LocalData()
	defaultPrivateKeyPath, _ := p.DefaultPrivateKey()

	trm.Output("\nEnvironment information")
	trm.Output("Working directory: " + cwd)
	trm.Output("Data directory: " + datadir.DataDir().String())
	trm.Output("Vagrantfile name: " + vagrantfileName)
	trm.Output("Home directory: " + home)
	trm.Output("Local data directory: " + localDataPath)
	trm.Output("Default private key path: " + defaultPrivateKeyPath)

	return 0
}

var (
	_ component.Command = (*Info)(nil)
)