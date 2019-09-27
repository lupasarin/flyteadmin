package mocks

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/lyft/flytestdlib/storage"
)

type NopCloser struct {
	io.Reader
}

func (NopCloser) Close() error { return nil }

type TestDataStore struct {
	HeadCb          func(ctx context.Context, reference storage.DataReference) (storage.Metadata, error)
	ReadProtobufCb  func(ctx context.Context, reference storage.DataReference, msg proto.Message) error
	WriteProtobufCb func(
		ctx context.Context, reference storage.DataReference, opts storage.Options, msg proto.Message) error
}

func (t *TestDataStore) Head(ctx context.Context, reference storage.DataReference) (storage.Metadata, error) {
	return t.HeadCb(ctx, reference)
}

func (t *TestDataStore) ReadProtobuf(ctx context.Context, reference storage.DataReference, msg proto.Message) error {
	return t.ReadProtobufCb(ctx, reference, msg)
}

func (t *TestDataStore) WriteProtobuf(
	ctx context.Context, reference storage.DataReference, opts storage.Options, msg proto.Message) error {
	return t.WriteProtobufCb(ctx, reference, opts, msg)
}

func (t *TestDataStore) GetBaseContainerFQN(ctx context.Context) storage.DataReference {
	return "s3://bucket"
}

// Retrieves a byte array from the Blob store or an error
func (t *TestDataStore) ReadRaw(ctx context.Context, reference storage.DataReference) (io.ReadCloser, error) {
	return NopCloser{}, nil
}

// Stores a raw byte array.
func (t *TestDataStore) WriteRaw(
	ctx context.Context, reference storage.DataReference, size int64, opts storage.Options, raw io.Reader) error {
	return nil
}

// Copies from source to destination.
func (t *TestDataStore) CopyRaw(ctx context.Context, source, destination storage.DataReference, opts storage.Options) error {
	return nil
}

func (t *TestDataStore) ConstructReference(
	ctx context.Context, reference storage.DataReference, nestedKeys ...string) (storage.DataReference, error) {
	nestedPath := strings.Join(nestedKeys, "/")
	return storage.DataReference(fmt.Sprintf("%s/%v", "s3://bucket", nestedPath)), nil
}

func GetMockStorageClient() *storage.DataStore {
	mockStorageClient := TestDataStore{}
	return &storage.DataStore{
		ComposedProtobufStore: &mockStorageClient,
		ReferenceConstructor:  &mockStorageClient,
	}
}