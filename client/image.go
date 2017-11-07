package client

import (
	"net/url"

	"github.com/alibaba/pouch/apis/types"
)

// ImagePull requests daemon to pull an image from registry.
func (cli *Client) ImagePull(name, tag string) error {
	q := url.Values{}
	q.Set("fromImage", name)
	q.Set("tag", tag)

	resp, err := cli.post("/images/create", q, nil)
	ensureCloseReader(resp)

	return err
}

// ImageList requests daemon to list all images
func (cli *Client) ImageList() ([]types.Image, error) {
	resp, err := cli.get("/images/json", nil)
	if err != nil {
		return nil, err
	}

	imageList := []types.Image{}

	err = decodeBody(&imageList, resp.Body)
	ensureCloseReader(resp)

	return imageList, err

}
