/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("0xhrpsmzoeig48k")

  // remove
  collection.schema.removeField("cdb2yyyy")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "nf6nne3z",
    "name": "action",
    "type": "number",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "noDecimal": false
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("0xhrpsmzoeig48k")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "cdb2yyyy",
    "name": "action",
    "type": "select",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "-1",
        "0",
        "1",
        "2",
        "3"
      ]
    }
  }))

  // remove
  collection.schema.removeField("nf6nne3z")

  return dao.saveCollection(collection)
})
