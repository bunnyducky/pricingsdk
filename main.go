package sdk

import (
	"fmt"
	"time"
)

type TokenBalance struct {
	OwnerAccount string     `json:"ownerAccount"`
	Amount       uint64     `json:"amount,omitempty"`
	SnapshotAt   *time.Time `json:"snapshotAt"`
}

func (c *Client) FetchTokenBalance(tokenMint string, timestamp uint64) ([]TokenBalance, error) {
	path := fmt.Sprintf("/token-balance/%s?timestamp=%d", tokenMint, timestamp)
	result := []TokenBalance{}
	err := c.get(path, &result)
	return result, err
}

type ParrotVault struct {
	Owner            string     `json:"owner"`
	Vault            string     `json:"vault"`
	VaultType        string     `json:"vaultType"`
	DebtAmount       uint64     `json:"debtAmount,omitempty"`
	CollateralAmount uint64     `json:"collateralAmount,omitempty"`
	SnapshotAt       *time.Time `json:"snapshotAt,omitempty"`
}

func (c *Client) FetchParrotVault(vaultType string, timestamp uint64) ([]ParrotVault, error) {
	path := fmt.Sprintf("/parrot-vaults/%s?timestamp=%d", vaultType, timestamp)
	result := []ParrotVault{}
	err := c.get(path, &result)
	return result, err
}