package models

type Role string

const (
	Tenant      Role = "tenant"
	TenantAdmin Role = "tenantAdmin"
	Root        Role = "root"
)
