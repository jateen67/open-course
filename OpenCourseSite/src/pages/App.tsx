import { useEffect, useState } from "react";
import Test from "../components/test";
import { Nav } from "../components/Nav/Nav"
import Colors, { applyTheme } from '../styles/ColorSystem';
import { MainContent } from "../components/MainContent/MainContent";

export default function App() {
  const [currentTheme, setCurrentTheme] = useState<keyof typeof Colors>("blue");

  useEffect(() => {
    applyTheme(Colors[currentTheme]);
  }, [currentTheme]);

  return (
    <>
      <Nav setCurrentTheme={setCurrentTheme} />
      <MainContent />
      <Test />
    </>
  );
}
