service: {
  name: "test"
  mode: "debug"
  mongodb: {
    username: ""
    hosts: ["127.0.0.1:27017"]
  }
  // logging: {
  //   level: "error"
  // }
  http: {
    port: 8787
  }
  grpc: {
    features: {
      healthCheck: true
    }
  }
}
