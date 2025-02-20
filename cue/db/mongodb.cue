package db

#Mongodb: {
  uri?: string
  username?: string
  password?: string
  dbName?: string
  hosts!: [string, ...string]
  options: {
    replicaSet?: string
  }
}

mongodb: #Mongodb
