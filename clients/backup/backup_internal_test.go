package backup

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
	"github.com/influxdata/influx-cli/v2/internal/mock"
	"github.com/stretchr/testify/require"
)

func TestBackup_DownloadMetadata(t *testing.T) {
	t.Parallel()

	fakeKV := strings.Repeat("I'm the bolt DB\n", 1234)
	fakeSQL := strings.Repeat("I'm the SQL!\n", 1234)

	bucketMetadata := []api.BucketMetadataManifest{
		{
			OrganizationID:         "123",
			OrganizationName:       "org",
			BucketID:               "456",
			BucketName:             "bucket1",
			DefaultRetentionPolicy: "foo",
			RetentionPolicies: []api.RetentionPolicyManifest{
				{Name: "foo"},
				{Name: "bar"},
			},
		},
		{
			OrganizationID:         "123",
			OrganizationName:       "org",
			BucketID:               "789",
			BucketName:             "bucket2",
			DefaultRetentionPolicy: "baz",
			RetentionPolicies: []api.RetentionPolicyManifest{
				{Name: "qux"},
				{Name: "baz"},
			},
		},
	}

	testCases := []struct {
		name                string
		compression         FileCompression
		responseCompression FileCompression
	}{
		{
			name:                "no gzip",
			compression:         NoCompression,
			responseCompression: NoCompression,
		},
		{
			name:                "response gzip, no local gzip",
			compression:         NoCompression,
			responseCompression: GzipCompression,
		},
		{
			name:                "no response gzip, local gzip",
			compression:         GzipCompression,
			responseCompression: NoCompression,
		},
		{
			name:                "all gzip",
			compression:         GzipCompression,
			responseCompression: GzipCompression,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			backupApi := mock.NewMockBackupApi(ctrl)
			backupApi.EXPECT().GetBackupMetadata(gomock.Any()).
				Return(api.ApiGetBackupMetadataRequest{ApiService: backupApi})
			backupApi.EXPECT().GetBackupMetadataExecute(gomock.Any()).
				DoAndReturn(func(request api.ApiGetBackupMetadataRequest) (*http.Response, error) {
					out := bytes.Buffer{}
					var outW io.Writer = &out
					if tc.responseCompression == GzipCompression {
						gzw := gzip.NewWriter(outW)
						defer gzw.Close()
						outW = gzw
					}

					parts := []struct {
						name        string
						contentType string
						writeFn     func(io.Writer) error
					}{
						{
							name:        "kv",
							contentType: "application/octet-stream",
							writeFn: func(w io.Writer) error {
								_, err := w.Write([]byte(fakeKV))
								return err
							},
						},
						{
							name:        "sql",
							contentType: "application/octet-stream",
							writeFn: func(w io.Writer) error {
								_, err := w.Write([]byte(fakeSQL))
								return err
							},
						},
						{
							name:        "buckets",
							contentType: "application/json",
							writeFn: func(w io.Writer) error {
								enc := json.NewEncoder(w)
								return enc.Encode(bucketMetadata)
							},
						},
					}

					writer := multipart.NewWriter(outW)
					for _, part := range parts {
						pw, err := writer.CreatePart(map[string][]string{
							"Content-Type":        {part.contentType},
							"Content-Disposition": {fmt.Sprintf("attachment; name=%s", part.name)},
						})
						require.NoError(t, err)
						require.NoError(t, part.writeFn(pw))
					}
					require.NoError(t, writer.Close())

					res := http.Response{Header: http.Header{}, Body: ioutil.NopCloser(&out)}
					res.Header.Add("Content-Type", fmt.Sprintf("multipart/mixed; boundary=%s", writer.Boundary()))
					if tc.responseCompression == GzipCompression {
						res.Header.Add("Content-Encoding", "gzip")
					}
					return &res, nil
				})

			stdio := mock.NewMockStdIO(ctrl)
			writtenBytes := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()
			log.SetOutput(stdio)

			cli := Client{
				CLI:       clients.CLI{StdIO: stdio},
				BackupApi: backupApi,
				baseName:  "test",
			}

			out, err := ioutil.TempDir("", "")
			require.NoError(t, err)
			defer os.RemoveAll(out)

			params := Params{
				Path:        out,
				Compression: tc.compression,
			}

			require.NoError(t, cli.downloadMetadata(context.Background(), &params))
			require.Equal(t, bucketMetadata, cli.bucketMetadata)

			localKv, err := os.Open(filepath.Join(out, cli.manifest.KV.FileName))
			require.NoError(t, err)
			defer localKv.Close()

			var kvReader io.Reader = localKv
			if tc.compression == GzipCompression {
				gzr, err := gzip.NewReader(kvReader)
				require.NoError(t, err)
				defer gzr.Close()
				kvReader = gzr
			}
			kvBytes, err := ioutil.ReadAll(kvReader)
			require.NoError(t, err)
			require.Equal(t, fakeKV, string(kvBytes))

			localSql, err := os.Open(filepath.Join(out, cli.manifest.SQL.FileName))
			require.NoError(t, err)
			defer localSql.Close()

			var sqlReader io.Reader = localSql
			if tc.compression == GzipCompression {
				gzr, err := gzip.NewReader(sqlReader)
				require.NoError(t, err)
				defer gzr.Close()
				sqlReader = gzr
			}
			sqlBytes, err := ioutil.ReadAll(sqlReader)
			require.NoError(t, err)
			require.Equal(t, fakeSQL, string(sqlBytes))
		})
	}
}

func TestBackup_DownloadShardData(t *testing.T) {
	t.Parallel()

	fakeTsm := strings.Repeat("Time series data!\n", 1024)

	testCases := []struct {
		name                string
		compression         FileCompression
		responseCompression FileCompression
	}{
		{
			name:                "no gzip",
			compression:         NoCompression,
			responseCompression: NoCompression,
		},
		{
			name:                "response gzip, no local gzip",
			compression:         NoCompression,
			responseCompression: GzipCompression,
		},
		{
			name:                "no response gzip, local gzip",
			compression:         GzipCompression,
			responseCompression: NoCompression,
		},
		{
			name:                "all gzip",
			compression:         GzipCompression,
			responseCompression: GzipCompression,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			backupApi := mock.NewMockBackupApi(ctrl)
			req := api.ApiGetBackupShardIdRequest{ApiService: backupApi}.ShardID(1)
			backupApi.EXPECT().GetBackupShardId(gomock.Any(), gomock.Eq(req.GetShardID())).Return(req)
			backupApi.EXPECT().GetBackupShardIdExecute(gomock.Any()).
				DoAndReturn(func(api.ApiGetBackupShardIdRequest) (*http.Response, error) {
					out := bytes.Buffer{}
					var outW io.Writer = &out
					if tc.responseCompression == GzipCompression {
						gzw := gzip.NewWriter(outW)
						defer gzw.Close()
						outW = gzw
					}
					_, err := outW.Write([]byte(fakeTsm))
					require.NoError(t, err)
					res := http.Response{Header: http.Header{}, Body: ioutil.NopCloser(&out)}
					res.Header.Add("Content-Type", "application/octet-stream")
					if tc.responseCompression == GzipCompression {
						res.Header.Add("Content-Encoding", "gzip")
					}
					return &res, nil
				})

			stdio := mock.NewMockStdIO(ctrl)
			writtenBytes := bytes.Buffer{}
			stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()
			log.SetOutput(stdio)

			cli := Client{
				CLI:       clients.CLI{StdIO: stdio},
				BackupApi: backupApi,
				baseName:  "test",
			}

			out, err := ioutil.TempDir("", "")
			require.NoError(t, err)
			defer os.RemoveAll(out)

			params := Params{
				Path:        out,
				Compression: tc.compression,
			}

			metadata, err := cli.downloadShardData(context.Background(), &params, req.GetShardID())
			require.NoError(t, err)
			require.NotNil(t, metadata)
			localShard, err := os.Open(filepath.Join(out, metadata.FileName))
			require.NoError(t, err)
			defer localShard.Close()

			var shardReader io.Reader = localShard
			if tc.compression == GzipCompression {
				gzr, err := gzip.NewReader(shardReader)
				require.NoError(t, err)
				defer gzr.Close()
				shardReader = gzr
			}
			shardBytes, err := ioutil.ReadAll(shardReader)
			require.NoError(t, err)
			require.Equal(t, fakeTsm, string(shardBytes))
		})
	}

	t.Run("shard deleted", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		backupApi := mock.NewMockBackupApi(ctrl)
		req := api.ApiGetBackupShardIdRequest{ApiService: backupApi}.ShardID(1)
		backupApi.EXPECT().GetBackupShardId(gomock.Any(), gomock.Eq(req.GetShardID())).Return(req)
		backupApi.EXPECT().GetBackupShardIdExecute(gomock.Any()).Return(nil, &notFoundErr{})

		stdio := mock.NewMockStdIO(ctrl)
		writtenBytes := bytes.Buffer{}
		stdio.EXPECT().Write(gomock.Any()).DoAndReturn(writtenBytes.Write).AnyTimes()
		log.SetOutput(stdio)

		cli := Client{
			CLI:       clients.CLI{StdIO: stdio},
			BackupApi: backupApi,
			baseName:  "test",
		}

		out, err := ioutil.TempDir("", "")
		require.NoError(t, err)
		defer os.RemoveAll(out)

		params := Params{
			Path: out,
		}

		metadata, err := cli.downloadShardData(context.Background(), &params, req.GetShardID())
		require.NoError(t, err)
		require.Nil(t, metadata)
		require.Contains(t, writtenBytes.String(), fmt.Sprintf("WARN: Shard %d removed during backup", req.GetShardID()))
	})
}

type notFoundErr struct{}

func (e *notFoundErr) Error() string {
	return "not found"
}

func (e *notFoundErr) ErrorCode() api.ErrorCode {
	return api.ERRORCODE_NOT_FOUND
}