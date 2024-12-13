import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./Home";
import Timeline from "./Timeline";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/timeline" element={<Timeline />} />
      </Routes>
    </Router>
  );
};

export default App;
