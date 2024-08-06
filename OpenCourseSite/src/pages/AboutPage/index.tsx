import { Nav } from "../../components/Nav/Nav";
import { useState, useEffect } from "react";
import Colors, { applyTheme } from '../../styles/ColorSystem';
import MainContentStyles from "../../components/MainContent/MainContent.module.css"

export const AboutPage: React.FC = () => {
    const [currentTheme, setCurrentTheme] = useState<keyof typeof Colors>("red");

    useEffect(() => {
        applyTheme(Colors[currentTheme]);
    }, [currentTheme]);

    return (
        <>
            <Nav setCurrentTheme={setCurrentTheme} />
            <div className={MainContentStyles.Content}>
                <h3>About OpenCourse</h3>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris et nisi eget ligula scelerisque euismod sed ac metus.
                    Praesent vel nisi vel erat luctus congue. Vestibulum euismod mollis est at hendrerit. Curabitur euismod mauris risus,
                    non faucibus eros accumsan vel. In interdum augue ex, at bibendum urna finibus in. Praesent tellus lacus, hendrerit
                    vitae luctus non, facilisis a risus. Proin consectetur ligula et tortor tempor vulputate. Sed posuere vestibulum
                    ex, a venenatis purus dapibus id. Mauris aliquet enim vel ornare scelerisque. Nam convallis est justo, eu porttitor
                    arcu scelerisque posuere. Ut ac rhoncus eros. Morbi ac nulla tempor, suscipit felis vitae, euismod libero.
                </p>
            </div>
        </>
    );
};