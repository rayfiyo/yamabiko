import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Main from "./Main";
import Timeline from "./Timeline";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/timeline" element={<Timeline />} />
      </Routes>
    </Router>
  );
};

export default App;
