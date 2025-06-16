import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import "./App.css";
import HeaderLogin from "./component/header/header.login";
import FirstLogin from "./component/section/first/first.login"
import About from "./screen/about/index";
import Home from "./screen/home/index";
import SecondLogin from "./component/section/second/second.login";
import ThreeLogin from "./component/section/three/three.login"



const firstLoad = true
function App() {
  if ( firstLoad )
  {
    return (
      <>
        <HeaderLogin />
        <FirstLogin />
        <SecondLogin />
        <ThreeLogin/>
      </>
    )
  }
  return (
    <div className="main">
      <Router>
        <HeaderLogin />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
