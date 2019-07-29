package repositories

import "inventory/iInfrastructures"

type InventoryRepository struct {
	iInfrastructures.IDbHandler
}
