import { ipc } from "photon-lib-js";

var input = document.getElementById("input");
var output = document.getElementById("output");
var button = document.querySelector("button");

const send = () => {
  ipc.emit("message", {
    message: input.value,
  });
  input.value = "";
};

button.addEventListener("click", send);

ipc.on("open", () => {
  output.innerHTML += "Status: Connected\n";
});

ipc.on("message", (payload, event) => {
  output.innerHTML += "Server: " + payload.message + "\n";
  output.innerHTML += "Server: " + payload.id + "\n";
  output.innerHTML += "Event: " + event + "\n";
});
