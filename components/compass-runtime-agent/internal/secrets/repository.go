package secrets

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/types"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ManagerConstructor creates Secret Manager for specified namespace
type ManagerConstructor func(namespace string) Manager

// Manager contains operations for managing k8s secrets
//
//go:generate mockery --name=Manager
type Manager interface {
	Get(ctx context.Context, name string, options metav1.GetOptions) (*v1.Secret, error)
	Create(ctx context.Context, secret *v1.Secret, options metav1.CreateOptions) (*v1.Secret, error)
	Update(ctx context.Context, secret *v1.Secret, options metav1.UpdateOptions) (*v1.Secret, error)
	Delete(ctx context.Context, name string, options metav1.DeleteOptions) error
}

// Repository contains operations for managing client credentials
//
//go:generate mockery --name=Repository
type Repository interface {
	Exists(name types.NamespacedName) (bool, error)
	Get(name types.NamespacedName) (map[string][]byte, error)
	UpsertWithReplace(name types.NamespacedName, data map[string][]byte) error
	UpsertWithMerge(name types.NamespacedName, data map[string][]byte) error
	Delete(secretName types.NamespacedName) error
}

type repository struct {
	secretsManagerConstructor ManagerConstructor
}

// NewRepository creates a new secrets repository
func NewRepository(secretsManagerConstructor ManagerConstructor) Repository {
	return &repository{
		secretsManagerConstructor: secretsManagerConstructor,
	}
}

func (r *repository) Exists(name types.NamespacedName) (bool, error) {
	secretManager := r.secretsManagerConstructor(name.Namespace)

	_, err := secretManager.Get(context.Background(), name.Name, metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// UpsertWithReplace creates a new Kubernetes secret, if secret with specified name already exists overrides it
func (r *repository) UpsertWithReplace(name types.NamespacedName, data map[string][]byte) error {
	secretManager := r.secretsManagerConstructor(name.Namespace)

	secret := makeSecret(name, data)

	_, err := secretManager.Create(context.Background(), secret, metav1.CreateOptions{})
	if err != nil {
		if k8serrors.IsAlreadyExists(err) {
			return r.replace(secretManager, secret)
		}

		return errors.Wrapf(err, "Replacing %s secret failed", name)
	}

	return err
}

func (r *repository) replace(secretManager Manager, secret *v1.Secret) error {
	err := secretManager.Delete(context.Background(), secret.Name, metav1.DeleteOptions{})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Deleting %s secret failed", secret.Name))
	}

	_, err = secretManager.Create(context.Background(), secret, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

// Get returns secret data for specified name
func (r *repository) Get(name types.NamespacedName) (map[string][]byte, error) {
	secretManager := r.secretsManagerConstructor(name.Namespace)

	secret, err := secretManager.Get(context.Background(), name.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return secret.Data, nil
}

// UpsertWithMerge updates secrets data with the provided values. If provided value already exists it will be updated.
// If secret does not exist it will be created
func (r *repository) UpsertWithMerge(name types.NamespacedName, data map[string][]byte) error {
	existingData, err := r.Get(name)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return errors.Wrap(err, "Failed to upsert secret data")
		}

		existingData = map[string][]byte{}
	}

	mergedData := mergeSecretData(existingData, data)
	return r.upsert(name, mergedData)
}

func (r *repository) Delete(name types.NamespacedName) error {
	secretManager := r.secretsManagerConstructor(name.Namespace)
	return secretManager.Delete(context.Background(), name.Name, metav1.DeleteOptions{})
}

func (r *repository) upsert(name types.NamespacedName, data map[string][]byte) error {
	secretManager := r.secretsManagerConstructor(name.Namespace)

	secret := makeSecret(name, data)

	_, err := secretManager.Update(context.Background(), secret, metav1.UpdateOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			_, err = secretManager.Create(context.Background(), secret, metav1.CreateOptions{})
			return err
		}
		return errors.Wrapf(err, "Updating %s secret failed while upserting", name)
	}
	return nil
}

func makeSecret(name types.NamespacedName, data map[string][]byte) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name.Name,
			Namespace: name.Namespace,
		},
		Data: data,
	}
}

func mergeSecretData(oldData, newData map[string][]byte) map[string][]byte {
	mergedMap := mergeMap(map[string][]byte{}, oldData)
	mergedMap = mergeMap(mergedMap, newData)

	return mergedMap
}

func mergeMap(base, merge map[string][]byte) map[string][]byte {
	for k, v := range merge {
		base[k] = v
	}

	return base
}
