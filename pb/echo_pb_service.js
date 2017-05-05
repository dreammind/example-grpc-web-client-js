// package: pb
// file: pb/echo.proto

var jspb = require("google-protobuf");
var pb_echo_pb = require("../pb/echo_pb");
var Echo = {
  serviceName: "pb.Echo"
};
Echo.Call = {
  methodName: "Call",
  service: Echo,
  requestStream: false,
  responseStream: false,
  requestType: pb_echo_pb.Request,
  responseType: pb_echo_pb.Response
};
module.exports = {
  Echo: Echo,
};

