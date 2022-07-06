require('dotenv').config()
const grpc = require("grpc");
import * as protoLoader from '@grpc/proto-loader';
const packageDef = protoLoader.loadSync("proto/gallery.proto", {});
const grpcObject = grpc.loadPackageDefinition(packageDef);
const gallery = grpcObject.gallery;

const server = new grpc.Server();

import { addImage, getImages } from "./actions";

// Without defining protocol
// 0.0.0.0 if you need to access an external IP
// createInsecure => without SSL
server.bind("0.0.0.0:5001", grpc.ServerCredentials.createInsecure());

server.addService(gallery.Gallery.service, {
  addImage: addImage,
  getImages: getImages,
});

server.start();
