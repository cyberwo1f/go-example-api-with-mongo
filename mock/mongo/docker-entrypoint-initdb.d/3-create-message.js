db.createCollection("message", {
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["id", "userId", "message"],
            properties: {
                id: {
                    bsonType: "int",
                    description: "must be int and is required and unique",
                },
                userId: {
                    bsonType: "int",
                    description: "must be int and is required and unique",
                },
                message: {
                    bsonType: "string",
                    description: "must be a string and is required"
                },
            }
        }
    }
});
db.message.createIndex({id: 1}, { unique: true })