package tool

import (
	"testing"
)

func TestListBuckets(t *testing.T)  {
	listBuckets()
}

func TestListBucketItems(t *testing.T)  {
	listBucketItems("webapp.jing.zhang")
}

func TestDeleteFile(t *testing.T) {
	DeleteFile("webapp.jing.zhang", "4.png")
}