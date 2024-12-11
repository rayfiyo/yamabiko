import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";

// Importing the Bootstrap CSS
import "bootstrap/dist/css/bootstrap.min.css";

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// https://codesandbox.io/p/sandbox/github/react-bootstrap/code-sandbox-examples/tree/master/basic?file=%2Fsrc%2Findex.js%3A1%2C1-9%2C1
// ReactDOM.render(<App />, document.getElementById('root'));
