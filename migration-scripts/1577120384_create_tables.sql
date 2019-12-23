CREATE TABLE "item"
(
    "id"          SERIAL NOT NULL,
    "name"        TEXT    NOT NULL DEFAULT '',
    "picture"     TEXT    NOT NULL DEFAULT '',
    "description" TEXT    NOT NULL DEFAULT '',
    PRIMARY KEY ("id")
);

CREATE TABLE "key_value_field"
(
    "id"      SERIAL NOT NULL,
    "item_id" INTEGER NOT NULL REFERENCES item(id),
    "name"    TEXT    NOT NULL DEFAULT '',
    "value"   TEXT    NOT NULL DEFAULT '',
    PRIMARY KEY ("id")
);
CREATE INDEX "item_id" ON "key_value_field" ("item_id");
