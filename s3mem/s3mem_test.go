/*
###############################################################################
# Licensed Materials - Property of IBM Copyright IBM Corporation 2017, 2019. All Rights Reserved.
# U.S. Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP
# Schedule Contract with IBM Corp.
#
# Contributors:
#  IBM Corporation - initial API and implementation
###############################################################################
*/

package s3mem

import (
	"context"
	"testing"

	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	S3MemBuckets.Mux.Lock()
	defer S3MemBuckets.Mux.Unlock()
	l := len(S3MemBuckets.Buckets)
	client := New(aws.Config{})
	//Create the request
	req := client.ListBucketsRequest(&s3.ListBucketsInput{})
	//Send the request
	listBucketsOutput, err := req.Send(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, l, len(listBucketsOutput.Buckets))
}

func TestNotImplemented(t *testing.T) {
	//Request a client
	client := New(aws.Config{})
	input := &s3.AbortMultipartUploadInput{}
	req := client.AbortMultipartUploadRequest(input)
	assert.Error(t, req.Error)
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}
