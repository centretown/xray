package dbg

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
	},

	Create: []string{

		`CREATE TABLE version (
itemMajor BIGINT, itemMinor BIGINT,			
major SMALLINT,
minor SMALLINT,
patch SMALLINT,
extension SMALLINT,
PRIMARY KEY (itemMajor,itemMajor,major,minor));`,

		`CREATE TABLE items (
title VARCHAR(80) NOT NULL,
category VARCHAR(16) NOT NULL,
content TEXT,
major BIGINT, minor BIGINT,			
origin_major BIGINT, origin_minor BIGINT,
created TIMESTAMP, updated TIMESTAMP,
encoding SMALLINT,
PRIMARY KEY (major,minor));`,

		"CREATE INDEX items_title ON items (title,category,origin_major,origin_minor);",
		"CREATE INDEX items_category ON items (category,title,origin_major,origin_minor);",
		"CREATE INDEX items_origin_t ON items (origin_major,origin_minor,title,category);",
		"CREATE INDEX items_origin_c ON items (origin_major,origin_minor,category,title);",

		`CREATE TABLE tags (
name VARCHAR(32) NOT NULL,
weight DOUBLE PRECISION,
major BIGINT, minor BIGINT,
PRIMARY KEY (name,major,minor));`,

		"CREATE INDEX tags_id ON tags (major,minor,weight);",
		"CREATE INDEX tags_name ON tags (name,weight,major,minor);",

		`CREATE TABLE links (
major BIGINT, minor BIGINT,
linked_major BIGINT, linked_minor BIGINT,
repeated SMALLINT, weight DOUBLE PRECISION,
PRIMARY KEY (major,minor,linked_major,linked_minor));`,

		// "CREATE UNIQUE INDEX links_link ON links (linked_major,linked_minor,major,minor);",
		"CREATE INDEX links_weight ON links (major,minor,weight);",
		"CREATE INDEX links_weight_link ON links (linked_major,linked_minor,weight);",
	},

	InsertVersion: `INSERT INTO version (major, minor, patch, extension) 
VALUES (:major, :minor, :patch, :extension);`,

	InsertItem: `INSERT INTO items (title, category, content, encoding, major, minor, origin_major, origin_minor, created, updated)
VALUES (:title, :category, :content, :encoding, :major, :minor, :origin_major, :origin_minor, :created, :updated);`,

	InsertLink: `INSERT INTO links (major, minor, linked_major, linked_minor, repeated, weight)
VALUES (:major, :minor, :linked_major, :linked_minor, :repeated, :weight);`,

	InsertTag: `INSERT INTO tags (name, weight, major, minor)
VALUES (:name, :weight, :major, :minor);`,

	GetVersion: `SELECT * FROM version WHERE major=$1 AND minor=$2;`,

	GetItem: `SELECT * FROM items WHERE major=$1 AND minor=$2;`,

	GetLinks: `SELECT * FROM links WHERE major=$1 AND minor=$2;`,

	GetLink: `SELECT * FROM links WHERE major=$1 AND minor=$2 AND linked_major=$3 AND linked_minor=$4;`,
}
