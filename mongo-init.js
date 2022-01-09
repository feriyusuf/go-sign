// Execute this script from docker-compose.yml to generate database

db.createUser(
    {
        user: "feriyusuf",
        pwd: "p4ssw0rd",
        roles: [
            {
                role: "readWrite",
                db: "go_sign"
            }
        ]
    }
);