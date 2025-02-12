// src/App.jsx

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Root from "./Root";
import Timeline from "./Timeline";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Root />} />
        <Route path="/timeline" element={<Timeline />} />
      </Routes>
    </Router>
  );
};

export default App;
