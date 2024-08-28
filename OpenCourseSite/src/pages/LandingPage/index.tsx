import { useEffect, useState } from "react";
import { Nav } from "../../components/Nav/Nav";
import Colors, { applyTheme } from "../../styles/ColorSystem";
import { MainContent } from "../../components/MainContent/MainContent";
import { Footer } from "components/Footer";

export const LandingPage: React.FC = () => {
  const [currentTheme] = useState<keyof typeof Colors>("burgundy");

  useEffect(() => {
    applyTheme(Colors[currentTheme]);
  }, [currentTheme]);

  return (
    <>
      <Nav />
      <MainContent />
      <Footer />
    </>
  );
};
