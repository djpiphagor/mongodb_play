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
    brand: 'vw',
    model: 'polo',
    engine: 1,
    engine_power: 100.5,
    number_plate: 'A001AA198',
  },
 {
    brand: 'ford',
    model: 'mustang',
    engine: 2,
    engine_power: 500.6,
    number_plate: 'B001BB198',
  },
]);
