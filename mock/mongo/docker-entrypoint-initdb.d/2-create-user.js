db.createCollection("user", {
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["id", "name"],
            properties: {
                id: {
                    bsonType: "int",
                    description: "must be int and is required and unique",
                },
                name: {
                    bsonType: "string",
                    description: "must be a string and is required"
                },
            }
        }
    }
});
db.user.createIndex({id: 1}, { unique: true })