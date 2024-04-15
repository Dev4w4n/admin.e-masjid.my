/*
Copyright © 2024 Rohaizan Roosley rohaizanr@gmail.com
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/Dev4w4n/admin.e-masjid.my/api/pb"
	"github.com/Dev4w4n/admin.e-masjid.my/cli/grpcutils"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tenants in E-Masjid.My Saas.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")

		list, err := retrieveTenants()

		if err != nil {
			log.Printf("Cannot retrieve the list of tenants: %s", err)
		}

		displayTenant(list)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func displayTenant(list *pb.TenantList) {
	// Displays the tenant list in a table
	w := tabwriter.NewWriter(os.Stdout, 10, 10, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "TENANT ID\t NAMESPACE\t DB NAME\t DB USER")
	for _, tenant := range list.TenantList {
		fmt.Fprintln(w, tenant.Id, "\t", tenant.NameSpace, "\t", tenant.DbName, "\t", tenant.DbUser)
	}

	w.Flush()
}

func retrieveTenants() (*pb.TenantList, error) {
	// Get gRPC connection
	conn, err := grpcutils.NewGrpcConnection()
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
		return nil, err
	}
	defer grpcutils.CloseGrpcConnection(conn)

	client := pb.NewTenantsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tenantList, err := client.FindAll(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	conn.Close()

	return tenantList, nil
}
