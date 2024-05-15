/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("0xhrpsmzoeig48k")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "b6a60gqs",
    "name": "pot",
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

  // remove
  collection.schema.removeField("b6a60gqs")

  return dao.saveCollection(collection)
})
