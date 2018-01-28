## grpc-web-hacker-news
An example app implementing a Hacker News reader. This example aims to demonstrate usage of grpc-web with React. It additionally shows how to integrate with Redux.

### Running
To start both the Go backend server and the frontend React application, run the following:
```bash
./start.sh
```

The backend server is running on `http://localhost:8900` while the frontend will by default start on `http://localhost:3000`

## Notable setup points

### Disable TSLint for protobuf generated classes
https://github.com/easyCZ/grpc-web-hacker-news/blob/master/app/tslint.json#L4
