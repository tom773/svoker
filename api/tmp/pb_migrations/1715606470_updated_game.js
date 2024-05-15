/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("0xhrpsmzoeig48k")

  // update
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

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("0xhrpsmzoeig48k")

  // update
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
        "predeal",
        "preflop",
        "preturn",
        "preriver",
        "postriver"
      ]
    }
  }))

  return dao.saveCollection(collection)
})
