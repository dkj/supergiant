package ui

import (
	"net/http"

	"github.com/supergiant/supergiant/pkg/client"
	"github.com/supergiant/supergiant/pkg/models"
)

func NewEntrypoint(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	return renderTemplate(w, "entrypoints/new.html", map[string]interface{}{
		"title":      "Entrypoints",
		"formAction": "/ui/entrypoints",
		"model": map[string]interface{}{
			"kube_id": nil,
			"name":    "",
		},
	})
}

func CreateEntrypoint(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	m := new(models.Entrypoint)
	if err := unmarshalFormInto(r, m); err != nil {
		return err
	}
	if err := sg.Entrypoints.Create(m); err != nil {
		return renderTemplate(w, "entrypoints/new.html", map[string]interface{}{
			"title":      "Entrypoints",
			"formAction": "/ui/entrypoints",
			"model":      m,
			"error":      err.Error(),
		})
	}
	http.Redirect(w, r, "/ui/entrypoints", http.StatusFound)
	return nil
}

func ListEntrypoints(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	fields := []map[string]interface{}{
		{
			"title": "Kube ID",
			"type":  "field_value",
			"field": "kube_id",
		},
		{
			"title": "Name",
			"type":  "field_value",
			"field": "name",
		},
		{
			"title": "Address",
			"type":  "field_value",
			"field": "address",
		},
	}
	return renderTemplate(w, "entrypoints/index.html", map[string]interface{}{
		"title":       "Entrypoints",
		"uiBasePath":  "/ui/entrypoints",
		"apiListPath": "/api/v0/entrypoints",
		"fields":      fields,
		"showNewLink": true,
		"batchActionPaths": map[string]string{
			"Delete": "/delete",
		},
	})
}

func GetEntrypoint(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r)
	if err != nil {
		return err
	}
	item := new(models.Entrypoint)
	if err := sg.Entrypoints.Get(id, item); err != nil {
		return err
	}
	return renderTemplate(w, "entrypoints/show.html", map[string]interface{}{
		"title": "Entrypoints",
		"model": item,
	})
}

func DeleteEntrypoint(sg *client.Client, w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(r)
	if err != nil {
		return err
	}
	item := new(models.Entrypoint)
	item.ID = id
	if err := sg.Entrypoints.Delete(item); err != nil {
		return err
	}
	// http.Redirect(w, r, "/ui/entrypoints", http.StatusFound)
	return nil
}
