/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("1hdnpj2b1o55f92")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "ol1mvlow",
    "name": "players",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "_pb_users_auth_",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("1hdnpj2b1o55f92")

  // remove
  collection.schema.removeField("ol1mvlow")

  return dao.saveCollection(collection)
})
