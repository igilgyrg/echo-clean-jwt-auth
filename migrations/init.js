conn = new Mongo();
db = conn.getDB("todo");

(async function () {
        let userResult = await db.users.insertOne({
            "email": "ig.pomazkov@gmail.com",
            "password": "$2a$10$qBLtELEVDECprsPT9gpz4uJPR1Sq22Jn/YCQnOpdFuOMiCr/1jWNa",
            "created_at": ISODate("2022-08-06T00:00:00.000Z")
        });
        await db.tasks.insert({
            "title": "Нотификации испаниский TIMATE",
            "user_id": userResult.insertedId,
            "description": "Добавлить нотификации на испанском языке TIMATE",
            "is_complete": true,
            "created_at": ISODate("2022-09-21T20:00:00.000Z")
        });
    }
)();


