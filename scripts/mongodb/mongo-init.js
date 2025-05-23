db.createUser(
        {
            user: "dbuser",
            pwd: "dbpassword",
            roles: [
                {
                    role: "readWrite",
                    db: "garage"
                }
            ]
        }
);

db = db.getSiblingDB('garage');

db.createCollection('cars');

db.cars.insertMany([
 {
    brand: 'VW',
    model: 'Polo',
    engine: 1,
    engine_power: 100.5,
    number_plate: '123OOO',
  },
 {
    brand: 'Ford',
    model: 'Mustang',
    engine: 2,
    engine_power: 500.6,
    number_plate: '456AAA',
  },
]);

db.cars.createIndex({ number_plate: 1 }, { unique: true });
