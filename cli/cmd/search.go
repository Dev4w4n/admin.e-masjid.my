/*
Copyright © 2024 Rohaizan Roosley rohaizanr@gmail.com
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Dev4w4n/admin.e-masjid.my/api/pb"
	"github.com/Dev4w4n/admin.e-masjid.my/cli/grpcutils"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for tenant by namespace",
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")

		tenant, err := searchTenant(namespace)

		if err != nil {
			log.Printf("Cannot find tenant by namespace: %s", err)
			return
		}

		displaySingleTenant(tenant)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("namespace", "n", "", "Define the new tenant namespace")
	searchCmd.MarkFlagRequired("namespace")
}

func searchTenant(namespace string) (*pb.Tenant, error) {
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

	request := &pb.TenantNamespaceRequest{NameSpace: namespace}

	tenant, err := client.FindByNamespace(ctx, request)

	if err != nil {
		return nil, err
	}

	conn.Close()

	return tenant, nil
}

func displaySingleTenant(tenant *pb.Tenant) {
	// Displays the single tenant in details
	fmt.Println("")
	fmt.Printf("TENANT ID: %d\n", tenant.Id)
	fmt.Printf("NAMESPACE: %s\n", tenant.NameSpace)
	fmt.Printf("DB NAME: %s\n", tenant.DbName)
	fmt.Printf("DB USER: %s\n", tenant.DbUser)
	fmt.Printf("DB PASSWORD: %s\n", tenant.DbUser)
	fmt.Printf("ALLOWED ORIGIN: %s\n", tenant.AllowedOrigin)
	fmt.Printf("MANAGER ROLE: %s\n", tenant.ManagerRole)
	fmt.Printf("USER ROLE: %s\n", tenant.UserRole)
	fmt.Printf("KEYCLOAK CLIENT ID: %s\n", tenant.KeycloakClientId)
	fmt.Printf("KEYCLOAK SERVER: %s\n", tenant.KeycloakServer)
	fmt.Printf("KEYCLOAK JWKS URL: %s\n", tenant.KeycloakJwksUrl)
	fmt.Println("")
}
