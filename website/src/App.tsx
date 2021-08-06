import { BrowserRouter } from "react-router-dom";
import "./App.css";
import { Navbar } from "./components";
import { Routes } from "./routes";

function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <Routes />
    </BrowserRouter>
  );
}

export default App;
