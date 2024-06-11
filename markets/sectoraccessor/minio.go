package sectoraccessor

import (
	"context"
	"fmt"
	"github.com/filecoin-project/dagstore/mount"
	"github.com/minio/minio-go"
	"net/url"
)

type MinioMount struct {
	client         *minio.Client
	bucketName     string
	objectName     string
	originalObject *minio.Object
}

func (m *MinioMount) Close() error {
	return m.originalObject.Close()
}

func (m *MinioMount) Fetch(ctx context.Context) (mount.Reader, error) {
	log.Infof("minio fetch values: %s, %s", m.bucketName, m.objectName)
	o, err := m.client.GetObject(m.bucketName, m.objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Errorf("minio fetch error: %v", err)
		return nil, err
	}
	m.originalObject = o
	return m.originalObject, nil
}

func (m *MinioMount) Info() mount.Info {
	return mount.Info{
		Kind:             mount.KindRemote,
		AccessRandom:     true,
		AccessSeek:       true,
		AccessSequential: true,
	}
}

func (m *MinioMount) Stat(ctx context.Context) (mount.Stat, error) {
	oi, err := m.client.StatObject(m.bucketName, m.objectName, minio.StatObjectOptions{})
	if err != nil {
		return mount.Stat{}, fmt.Errorf("minio stat error: %v", err)
	}
	return mount.Stat{
		Exists: true,
		Size:   oi.Size,
		Ready:  true,
	}, nil
}

func (m *MinioMount) Serialize() *url.URL {
	u, _ := m.client.PresignedGetObject(m.bucketName, m.objectName, 0, url.Values{})
	return u
}

func (m *MinioMount) Deserialize(url *url.URL) error {
	return nil
}

var _ mount.Mount = (*MinioMount)(nil)
