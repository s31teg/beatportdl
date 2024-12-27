package beatport

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Label struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Slug    string    `json:"slug"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func (l *Label) DirectoryName(template string, whitespace string, aLimit int, aShortForm string) string {
	templateValues := map[string]string{
		"id":           strconv.Itoa(int(l.ID)),
		"name":         SanitizeForPath(l.Name),
		"slug":         l.Slug,
		"created_date": l.Created.Format("2006-01-02"),
		"updated_date": l.Updated.Format("2006-01-02"),
	}
	directoryName := ParseTemplate(template, templateValues)
	return SanitizePath(directoryName, whitespace)
}

func (l *Label) StoreUrl() string {
	return fmt.Sprintf("https://www.beatport.com/label/%s/%d", l.Slug, l.ID)
}

func (b *Beatport) GetLabel(id int64) (*Label, error) {
	res, err := b.fetch(
		"GET",
		fmt.Sprintf("/catalog/labels/%d/", id),
		nil,
		"",
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	response := &Label{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}
	return response, nil
}

func (b *Beatport) GetLabelReleases(id int64, page int) (*Paginated[Release], error) {
	res, err := b.fetch(
		"GET",
		fmt.Sprintf("/catalog/labels/%d/releases/?page=%d", id, page),
		nil,
		"",
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response Paginated[Release]
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
