/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("z3q69ih6a71khj1")

  collection.name = "gametable"

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "03fe9bru",
    "name": "cards",
    "type": "json",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "maxSize": 2000000
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("z3q69ih6a71khj1")

  collection.name = "usertables"

  // remove
  collection.schema.removeField("03fe9bru")

  return dao.saveCollection(collection)
})
