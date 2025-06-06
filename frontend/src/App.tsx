import React from "react";
import { Link, Route, BrowserRouter as Router, Routes } from "react-router-dom";
import "./App.css";
import About from "./screen/about/index";
import Home from "./screen/home/index";

const DEFAULT_CATALOG = [
  {
    name: "Home",
    path: "/",
  },
  {
    name: "Screener",
    path: "/screener",
  },
  {
    name: "Analitics",
    path: "/analitics",
  },
  {
    name: "Comunity",
    path: "/Comunity",
  },
  {
    name: "News",
    path: "/News",
  },
  {
    name: "Dao",
    path: "/dao",
  },
];

const NavContainer = () => {
  const [state, setState] = React.useState(DEFAULT_CATALOG);
  const [selected, setSelected] = React.useState(DEFAULT_CATALOG[0].name);

  return (
    <div className="flex row">
      {state.map((item, index) => (
        <Link key={index} to={item.path}>
          <div
            onClick={() => setSelected(item.name)}
            className="px-4 px-2 mx-2 hover:text-green-700 hover:shadow-lg hover:scale-105 transition-all duration-300"
          >
            <p>{item.name}</p>
          </div>
        </Link>
      ))}
    </div>
  );
};

function App() {
  return (
    <div className="main">
      <Router>
        <NavContainer />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
