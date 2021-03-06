package api

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"net/http"
	"testing"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestLinkAlbum(t *testing.T) {
	t.Run("create share link", func(t *testing.T) {
		app, router, _ := NewApiTest()

		var link entity.Link

		CreateAlbumLink(router)

		resp := PerformRequestWithBody(app, "POST", "/api/v1/albums/at9lxuqxpogaaba7/links", `{"Password": "foobar", "Expires": 0, "CanEdit": true}`)

		if resp.Code != http.StatusOK {
			t.Fatal(resp.Body.String())
		}

		if err := json.Unmarshal(resp.Body.Bytes(), &link); err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, link.LinkUID)
		assert.NotEmpty(t, link.ShareUID)
		assert.NotEmpty(t, link.LinkToken)
		assert.Equal(t, true, link.CanEdit)
		assert.Equal(t, 0, link.LinkExpires)
		assert.False(t, link.CanComment)
		assert.True(t, link.CanEdit)
	})
	t.Run("album does not exist", func(t *testing.T) {
		app, router, _ := NewApiTest()
		CreateAlbumLink(router)
		resp := PerformRequestWithBody(app, "POST", "/api/v1/albums/xxx/links", `{"Password": "foobar", "Expires": 0, "CanEdit": true}`)

		if resp.Code != http.StatusNotFound {
			t.Fatal(resp.Body.String())
		}

		val := gjson.Get(resp.Body.String(), "error")
		assert.Equal(t, "Album not found", val.String())
	})
	t.Run("invalid request", func(t *testing.T) {
		app, router, _ := NewApiTest()

		CreateAlbumLink(router)

		resp := PerformRequestWithBody(app, "POST", "/api/v1/albums/at9lxuqxpogaaba7/links", `{"Password": "foobar", "Expires": "abc", "CanEdit": true}`)

		if resp.Code != http.StatusBadRequest {
			t.Fatal(resp.Body.String())
		}
	})
}

func TestLinkPhoto(t *testing.T) {
	t.Run("create share link", func(t *testing.T) {
		app, router, _ := NewApiTest()

		var link entity.Link

		CreatePhotoLink(router)

		resp := PerformRequestWithBody(app, "POST", "/api/v1/photos/pt9jtdre2lvl0yh7/links", `{"Password": "foobar", "Expires": 0, "CanEdit": true}`)
		assert.Equal(t, http.StatusOK, resp.Code)

		if err := json.Unmarshal(resp.Body.Bytes(), &link); err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, link.LinkUID)
		assert.NotEmpty(t, link.ShareUID)
		assert.NotEmpty(t, link.LinkToken)
		assert.Equal(t, 0, link.LinkExpires)
		assert.False(t, link.CanComment)
		assert.True(t, link.CanEdit)
	})
	t.Run("photo not found", func(t *testing.T) {
		app, router, _ := NewApiTest()

		CreatePhotoLink(router)

		resp := PerformRequestWithBody(app, "POST", "/api/v1/photos/xxx/link", `{"Password": "foobar", "Expires": 0, "CanEdit": true}`)

		if resp.Code != http.StatusNotFound {
			t.Fatal(resp.Body.String())
		}
	})
	t.Run("invalid request", func(t *testing.T) {
		app, router, _ := NewApiTest()
		CreatePhotoLink(router)
		r := PerformRequestWithBody(app, "POST", "/api/v1/photos/pt9jtdre2lvl0yh7/links", `{"xxx": 123, "Expires": "abc", "CanEdit": "xxx"}`)
		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestLinkLabel(t *testing.T) {
	t.Run("create share link", func(t *testing.T) {
		app, router, _ := NewApiTest()

		var link entity.Link

		CreateLabelLink(router)

		resp := PerformRequestWithBody(app, "POST", "/api/v1/labels/lt9k3pw1wowuy3c2/links", `{"Password": "foobar", "Expires": 0, "CanEdit": true}`)
		assert.Equal(t, http.StatusOK, resp.Code)

		if err := json.Unmarshal(resp.Body.Bytes(), &link); err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, link.LinkUID)
		assert.NotEmpty(t, link.ShareUID)
		assert.NotEmpty(t, link.LinkToken)
		assert.Equal(t, 0, link.LinkExpires)
		assert.False(t, link.CanComment)
		assert.True(t, link.CanEdit)
	})
	t.Run("label not found", func(t *testing.T) {
		app, router, _ := NewApiTest()
		CreateLabelLink(router)
		resp := PerformRequestWithBody(app, "POST", "/api/v1/labels/xxx/links", `{"Password": "foobar", "Expires": 0, "CanEdit": true}`)

		if resp.Code != http.StatusNotFound {
			t.Fatal(resp.Body.String())
		}
	})
	t.Run("invalid request", func(t *testing.T) {
		app, router, _ := NewApiTest()
		CreateLabelLink(router)
		r := PerformRequestWithBody(app, "POST", "/api/v1/labels/lt9k3pw1wowuy3c2/links", `{"xxx": 123, "Expires": "abc", "CanEdit": "xxx"}`)
		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}
