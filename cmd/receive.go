package cmd

import (
	"fmt"

	"github.com/claudiodangelis/qrcp/config"
	"github.com/claudiodangelis/qrcp/qr"
	"github.com/claudiodangelis/qrcp/server"
	"github.com/spf13/cobra"
)

func receiveCmdFunc(comman *cobra.Command, args []string) error {
	// Load configuration
	cfg := config.Load()
	// Create the server
	// TODO: Remove duplicate code
	// TODO: Maybe move this somewhere else?
	if portFlag > 0 {
		cfg.Port = portFlag
	}
	if interfaceFlag != "" {
		cfg.Interface = interfaceFlag
	}
	srv, err := server.New(cfg.Interface, cfg.Port, false)
	if err != nil {
		return err
	}
	srv.ReceiveTo(outputFlag)
	fmt.Println(srv.ReceiveURL)
	qr.RenderString(srv.ReceiveURL)
	if err := srv.Wait(); err != nil {
		return err
	}
	return nil
}

var receiveCmd = &cobra.Command{
	Use:     "receive",
	Aliases: []string{"r"},
	// TODO add usage
	RunE: receiveCmdFunc,
}
