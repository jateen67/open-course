
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { LandingPage } from "./LandingPage";
import { AboutPage } from "./AboutPage";

export default function App() {
  return (
    <>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<LandingPage />} />
            <Route path="/about" element={<AboutPage />} />
          </Routes>
        </BrowserRouter>
    </>
  );
}
