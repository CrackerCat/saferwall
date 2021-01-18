// Copyright 2021 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package client

import (
	"context"

	"github.com/saferwall/saferwall/pkg/grpc/multiav"
	pb "github.com/saferwall/saferwall/pkg/grpc/multiav/eset/proto"
)

// GetVerion returns version
func GetVerion(client pb.EsetScannerClient) (*pb.VersionResponse, error) {
	versionRequest := &pb.VersionRequest{}
	return client.GetVersion(context.Background(), versionRequest)
}

// ScanFile scans file
func ScanFile(client pb.EsetScannerClient, path string) (multiav.ScanResult, error) {
	scanFile := &pb.ScanFileRequest{Filepath: path}
	ctx, cancel := context.WithTimeout(context.Background(), multiav.ScanTimeout)
	defer cancel()
	res, err := client.ScanFile(ctx, scanFile)
	if err != nil {
		return multiav.ScanResult{}, err
	}

	return multiav.ScanResult{
		Output:   res.Output,
		Infected: res.Infected,
		Update:   res.Update,
	}, nil
}
