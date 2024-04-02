package dbio

import (
	"time"

	"github.com/centretown/xray/model"
)

var SchemaGame = &Schema{
	Version: model.Version{
		Major:     0,
		Minor:     1,
		Patch:     0,
		Extension: 0,
		Created:   time.Now(),
		Updated:   time.Now(),
		Origin:    model.IN_HOUSE,
	},

	Create: []string{

		`CREATE TABLE version (
major SMALLINT,
minor SMALLINT,
patch SMALLINT,
extension SMALLINT);`,

		`CREATE TABLE items (
id VARCHAR(40) NOT NULL,
origin VARCHAR(40) NOT NULL,
title VARCHAR(128) NOT NULL,
category VARCHAR(40) NOT NULL,
created TIMESTAMP,
updated TIMESTAMP,
encoding VARCHAR(16),
content TEXT,
PRIMARY KEY (id));`,

		"CREATE INDEX items_title ON items (title,origin);",
		"CREATE INDEX items_category ON items (category,title,origin);",
		"CREATE INDEX items_origin_t ON items (origin,title,category);",
		"CREATE INDEX items_origin_c ON items (origin,category,title);",

		`CREATE TABLE tags (
id VARCHAR(40) NOT NULL,
link VARCHAR(40) NOT NULL,
title VARCHAR(128) NOT NULL,
weight INTEGER,
PRIMARY KEY (id));`,

		"CREATE INDEX tags_id ON tags (id,link,weight);",
		"CREATE INDEX tags_title ON tags (title,id,weight,link);",
		"CREATE INDEX tags_weight ON tags (id,weight,link);",
		"CREATE INDEX tags_link ON tags (link,id,weight);",

		`CREATE TABLE links (
id VARCHAR(40) NOT NULL,
link VARCHAR(40) NOT NULL,
PRIMARY KEY (id,link));`,

		"CREATE INDEX links_link ON links (link,id);",
	},

	InsertVersion: "INSERT INTO version (major, minor, patch, extension) " +
		"VALUES (:major, :minor, :patch, :extension);",
	InsertItem: "INSERT INTO items (id, origin, title, category, created, updated, encoding, content) " +
		"VALUES (:id, :origin, :title, :category, :created, :updated, :encoding, :content);",
}
