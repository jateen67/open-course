import { Nav } from "../../components/Nav/Nav";
import { useState, useEffect } from "react";
import Colors, { applyTheme } from '../../styles/ColorSystem';
import MainContentStyles from "./MainContent.module.css"

export const AboutPage: React.FC = () => {
    const [currentTheme, setCurrentTheme] = useState<keyof typeof Colors>("blue");

    useEffect(() => {
        applyTheme(Colors[currentTheme]);
    }, [currentTheme]);

    return (
        <>
            <Nav setCurrentTheme={setCurrentTheme} />
            <div>
                <p>Text</p>
            </div>
        </>
    );
};