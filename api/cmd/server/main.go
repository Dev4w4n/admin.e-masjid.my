package main

import (
	"context"
	"log"
	"net"

	"github.com/Dev4w4n/admin.e-masjid.my/api/model"
	"github.com/Dev4w4n/admin.e-masjid.my/api/pb"
	"github.com/Dev4w4n/admin.e-masjid.my/api/repository"
	"github.com/Dev4w4n/e-masjid.my/api/core/config"
	"github.com/Dev4w4n/e-masjid.my/api/core/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	log.Println("Starting server ...")

	env, err := env.GetEnvironment()
	if err != nil {
		log.Fatalf("Error getting environment: %v", err)
	}

	db, err := config.DatabaseConnection(env)
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}

	tenantRepository := repository.NewTenantRepository(db)

	addr := ":" + env.ServerPort
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error can't listen: %s", err)
	}

	srv := grpc.NewServer()
	tenantSrv := &TenantsServer{
		tenantRepository: tenantRepository,
	}

	pb.RegisterTenantsServer(srv, tenantSrv)
	reflection.Register(srv)

	log.Printf("Server started on %s", addr)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type TenantsServer struct {
	pb.UnimplementedTenantsServer
	tenantRepository repository.TenantRepository
}

func (s *TenantsServer) FindAll(ctx context.Context, req *emptypb.Empty) (*pb.TenantList, error) {
	// Use tenantRepository here to fetch tenants
	tenants, err := s.tenantRepository.FindAll()
	if err != nil {
		log.Printf("Error fetching tenants: %v", err)
		return nil, err
	}

	// Convert tenants to pb.TenantList
	tenantList := &pb.TenantList{
		TenantList: make([]*pb.Tenant, len(tenants)),
	}
	for i, tenant := range tenants {
		tenantList.TenantList[i] = &pb.Tenant{
			// Fill in tenant fields
			Id:               tenant.Id,
			DbHost:           tenant.DbHost,
			DbUser:           tenant.DbUser,
			DbPassword:       tenant.DbPassword,
			DbName:           tenant.DbName,
			AllowedOrigin:    tenant.AllowedOrigin,
			ManagerRole:      tenant.ManagerRole,
			UserRole:         tenant.UserRole,
			KeycloakClientId: tenant.KeycloakClientId,
			KeycloakServer:   tenant.KeycloakServer,
			KeycloakJwksUrl:  tenant.KeycloakJwksUrl,
			CreatedAt:        tenant.CreatedAt,
		}
	}

	return tenantList, nil
}

func (s *TenantsServer) Upsert(ctx context.Context, req *pb.Tenant) (*pb.Tenant, error) {
	// Convert pb.Tenant to model.Tenant
	mTenant := model.Tenant{
		DbHost:           req.DbHost,
		DbUser:           req.DbUser,
		DbPassword:       req.DbPassword,
		DbName:           req.DbName,
		AllowedOrigin:    req.AllowedOrigin,
		ManagerRole:      req.ManagerRole,
		UserRole:         req.UserRole,
		KeycloakClientId: req.KeycloakClientId,
		KeycloakServer:   req.KeycloakServer,
		KeycloakJwksUrl:  req.KeycloakJwksUrl,
	}

	// Use tenantRepository here to upsert tenant
	_tenant, err := s.tenantRepository.Upsert(&mTenant)

	if err != nil {
		log.Printf("Error upserting tenant: %v", err)
		return nil, err
	}

	// Convert model.Tenant to pb.Tenant
	pbTenant := &pb.Tenant{
		// Fill in tenant fields
		Id:               _tenant.Id,
		DbHost:           _tenant.DbHost,
		DbUser:           _tenant.DbUser,
		DbPassword:       _tenant.DbPassword,
		DbName:           _tenant.DbName,
		AllowedOrigin:    _tenant.AllowedOrigin,
		ManagerRole:      _tenant.ManagerRole,
		UserRole:         _tenant.UserRole,
		KeycloakClientId: _tenant.KeycloakClientId,
		KeycloakServer:   _tenant.KeycloakServer,
		KeycloakJwksUrl:  _tenant.KeycloakJwksUrl,
		CreatedAt:        _tenant.CreatedAt,
	}

	return pbTenant, nil
}

func (s *TenantsServer) Delete(ctx context.Context, req *pb.TenantIdRequest) (*pb.TenantServiceResponse, error) {
	// Use tenantRepository here to delete tenant
	response := &pb.TenantServiceResponse{
		Status: false,
	}

	err := s.tenantRepository.Delete(req.Id)

	if err != nil {
		log.Printf("Error deleting tenant: %v", err)
		return response, err
	}

	// Use pb.TenantServiceResponse to return status
	response.Status = true
	return response, nil
}
