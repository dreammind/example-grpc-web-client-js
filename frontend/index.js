import {grpc} from "grpc-web-client";

import {Echo} from "../pb/echo_pb_service.js";
import {Request} from "../pb/echo_pb.js";

function EchoCall(value) {
  const req = new Request();
  req.setValue(value);

  grpc.invoke(Echo.Call, {
    request: req,
    host: "https://localhost:9090",
    onMessage: (message) => {
      console.log("onMessage", message.toObject());
      alert(message.getValue());
    },
    onEnd: (code, msg, trailers) => {
      console.log("onEnd", code, msg, trailers);
    }
  });
}

EchoCall("Hello grpc-web-client");
