import { BrowserRouter, Route, Routes } from "react-router-dom";
import HomePage from "./pages/Home";
import Navbar from "./components/Navbar";
import TimelinePage from "./pages/Timeline";

function App() {
  const isLogin = false;
  return (
    <div>
      <BrowserRouter>
        <Navbar />
        <Routes>
          {isLogin ? (
            <Route path="/" element={<TimelinePage />} />
          ) : (
            <Route path="/" element={<HomePage />} />
          )}
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
