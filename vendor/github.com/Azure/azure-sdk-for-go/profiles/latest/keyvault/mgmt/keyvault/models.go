// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package keyvault

import original "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2016-10-01/keyvault"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type CertificatePermissions = original.CertificatePermissions

const (
	Create         CertificatePermissions = original.Create
	Delete         CertificatePermissions = original.Delete
	Deleteissuers  CertificatePermissions = original.Deleteissuers
	Get            CertificatePermissions = original.Get
	Getissuers     CertificatePermissions = original.Getissuers
	Import         CertificatePermissions = original.Import
	List           CertificatePermissions = original.List
	Listissuers    CertificatePermissions = original.Listissuers
	Managecontacts CertificatePermissions = original.Managecontacts
	Manageissuers  CertificatePermissions = original.Manageissuers
	Purge          CertificatePermissions = original.Purge
	Recover        CertificatePermissions = original.Recover
	Setissuers     CertificatePermissions = original.Setissuers
	Update         CertificatePermissions = original.Update
)

type CreateMode = original.CreateMode

const (
	CreateModeDefault CreateMode = original.CreateModeDefault
	CreateModeRecover CreateMode = original.CreateModeRecover
)

type KeyPermissions = original.KeyPermissions

const (
	KeyPermissionsBackup    KeyPermissions = original.KeyPermissionsBackup
	KeyPermissionsCreate    KeyPermissions = original.KeyPermissionsCreate
	KeyPermissionsDecrypt   KeyPermissions = original.KeyPermissionsDecrypt
	KeyPermissionsDelete    KeyPermissions = original.KeyPermissionsDelete
	KeyPermissionsEncrypt   KeyPermissions = original.KeyPermissionsEncrypt
	KeyPermissionsGet       KeyPermissions = original.KeyPermissionsGet
	KeyPermissionsImport    KeyPermissions = original.KeyPermissionsImport
	KeyPermissionsList      KeyPermissions = original.KeyPermissionsList
	KeyPermissionsPurge     KeyPermissions = original.KeyPermissionsPurge
	KeyPermissionsRecover   KeyPermissions = original.KeyPermissionsRecover
	KeyPermissionsRestore   KeyPermissions = original.KeyPermissionsRestore
	KeyPermissionsSign      KeyPermissions = original.KeyPermissionsSign
	KeyPermissionsUnwrapKey KeyPermissions = original.KeyPermissionsUnwrapKey
	KeyPermissionsUpdate    KeyPermissions = original.KeyPermissionsUpdate
	KeyPermissionsVerify    KeyPermissions = original.KeyPermissionsVerify
	KeyPermissionsWrapKey   KeyPermissions = original.KeyPermissionsWrapKey
)

type SecretPermissions = original.SecretPermissions

const (
	SecretPermissionsBackup  SecretPermissions = original.SecretPermissionsBackup
	SecretPermissionsDelete  SecretPermissions = original.SecretPermissionsDelete
	SecretPermissionsGet     SecretPermissions = original.SecretPermissionsGet
	SecretPermissionsList    SecretPermissions = original.SecretPermissionsList
	SecretPermissionsPurge   SecretPermissions = original.SecretPermissionsPurge
	SecretPermissionsRecover SecretPermissions = original.SecretPermissionsRecover
	SecretPermissionsRestore SecretPermissions = original.SecretPermissionsRestore
	SecretPermissionsSet     SecretPermissions = original.SecretPermissionsSet
)

type SkuName = original.SkuName

const (
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type StoragePermissions = original.StoragePermissions

const (
	StoragePermissionsDelete        StoragePermissions = original.StoragePermissionsDelete
	StoragePermissionsDeletesas     StoragePermissions = original.StoragePermissionsDeletesas
	StoragePermissionsGet           StoragePermissions = original.StoragePermissionsGet
	StoragePermissionsGetsas        StoragePermissions = original.StoragePermissionsGetsas
	StoragePermissionsList          StoragePermissions = original.StoragePermissionsList
	StoragePermissionsListsas       StoragePermissions = original.StoragePermissionsListsas
	StoragePermissionsRegeneratekey StoragePermissions = original.StoragePermissionsRegeneratekey
	StoragePermissionsSet           StoragePermissions = original.StoragePermissionsSet
	StoragePermissionsSetsas        StoragePermissions = original.StoragePermissionsSetsas
	StoragePermissionsUpdate        StoragePermissions = original.StoragePermissionsUpdate
)

type AccessPolicyEntry = original.AccessPolicyEntry
type DeletedVault = original.DeletedVault
type DeletedVaultListResult = original.DeletedVaultListResult
type DeletedVaultProperties = original.DeletedVaultProperties
type Permissions = original.Permissions
type Resource = original.Resource
type ResourceListResult = original.ResourceListResult
type Sku = original.Sku
type Vault = original.Vault
type VaultCreateOrUpdateParameters = original.VaultCreateOrUpdateParameters
type VaultListResult = original.VaultListResult
type VaultProperties = original.VaultProperties
type VaultsClient = original.VaultsClient

func NewVaultsClient(subscriptionID string) VaultsClient {
	return original.NewVaultsClient(subscriptionID)
}
func NewVaultsClientWithBaseURI(baseURI string, subscriptionID string) VaultsClient {
	return original.NewVaultsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
