package azure

import (
	"context"
	"net/http"
	"strings"
	"testing"

	kedav1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
)

func TestGetBlobLength(t *testing.T) {
	httpClient := http.DefaultClient
	length, err := GetAzureBlobListLength(context.TODO(), httpClient, kedav1alpha1.AuthPodIdentity{}, "", "blobContainerName", "", "", "", "")
	if length != -1 {
		t.Error("Expected length to be -1, but got", length)
	}

	if err == nil {
		t.Error("Expected error for empty connection string, but got nil")
	}

	if !strings.Contains(err.Error(), "parse storage connection string") {
		t.Error("Expected error to contain parsing error message, but got", err.Error())
	}

	length, err = GetAzureBlobListLength(context.TODO(), httpClient, kedav1alpha1.AuthPodIdentity{}, "DefaultEndpointsProtocol=https;AccountName=name;AccountKey=key==;EndpointSuffix=core.windows.net", "blobContainerName", "", "", "", "")

	if length != -1 {
		t.Error("Expected length to be -1, but got", length)
	}

	if err == nil {
		t.Error("Expected error for empty connection string, but got nil")
	}

	if !strings.Contains(err.Error(), "illegal base64") {
		t.Error("Expected error to contain base64 error message, but got", err.Error())
	}
}
