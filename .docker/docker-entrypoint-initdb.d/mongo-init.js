db = db.getSiblingDB(process.env.MONGO_INITDB_DATABASE);

rs.initiate();

db.createCollection("dev");