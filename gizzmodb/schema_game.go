package gizzmodb

import (
	"time"

	"github.com/centretown/xray/gizzmodb/model"
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
item BIGINT, itemn BIGINT,			
major SMALLINT,
minor SMALLINT,
patch SMALLINT,
extension SMALLINT,
PRIMARY KEY (item,itemn,major,minor));`,

		`CREATE TABLE items (
class VARCHAR(80) NOT NULL,
classn INTEGER,
content TEXT,
major BIGINT, minor BIGINT,			
origin BIGINT, originn BIGINT,
created TIMESTAMP, updated TIMESTAMP,
encoding SMALLINT,
PRIMARY KEY (major,minor));`,

		"CREATE INDEX items_class ON items (class,classn,origin,originn);",
		"CREATE INDEX items_classn ON items (classn,class,origin,originn);",
		"CREATE INDEX items_origin_t ON items (origin,originn,class,classn);",
		"CREATE INDEX items_origin_c ON items (origin,originn,classn,class);",

		`CREATE TABLE tags (
name VARCHAR(32) NOT NULL,
weight DOUBLE PRECISION,
major BIGINT, minor BIGINT,
PRIMARY KEY (name,major,minor));`,

		"CREATE INDEX tags_id ON tags (major,minor,weight);",
		"CREATE INDEX tags_name ON tags (name,weight,major,minor);",

		`CREATE TABLE links (
major BIGINT, minor BIGINT,
linked BIGINT, linkedn BIGINT,
repeated SMALLINT, weight DOUBLE PRECISION,
PRIMARY KEY (major,minor,linked,linkedn));`,

		// "CREATE UNIQUE INDEX links_link ON links (linked,linkedn,major,minor);",
		"CREATE INDEX links_weight ON links (major,minor,weight);",
		"CREATE INDEX links_weight_link ON links (linked,linkedn,weight);",
	},

	InsertVersion: `INSERT INTO version (item, itemn, major, minor, patch, extension) 
VALUES (:item,:itemn,:major,:minor,:patch,:extension);`,

	InsertItem: `INSERT INTO items (class, classn, content, encoding, major, minor, origin, originn, created, updated)
VALUES (:class,:classn,:content,:encoding,:major,:minor,:origin,:originn,:created,:updated);`,

	InsertLink: `INSERT INTO links (major, minor, linked, linkedn, repeated, weight)
VALUES (:major,:minor,:linked,:linkedn,:repeated,:weight);`,

	InsertTag: `INSERT INTO tags (name, weight, major, minor)
VALUES (:name,:weight,:major,:minor);`,

	GetVersion: `SELECT * FROM version WHERE item=$1 AND itemn=$2 AND major=$3 AND minor=$4;`,

	GetVersions: `SELECT item FROM version;`,

	GetItem: `SELECT * FROM items WHERE major=$1 AND minor=$2;`,

	GetLinks: `SELECT * FROM links WHERE major=$1 AND minor=$2;`,

	GetLink: `SELECT * FROM links WHERE major=$1 AND minor=$2 AND linked=$3 AND linkedn=$4;`,
}
