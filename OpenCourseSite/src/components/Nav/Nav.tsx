import { SelectMenu } from "./SelectMenu"
import NavStyles from "./Nav.module.css"
import Colors from "../../styles/ColorSystem"
import { Link } from "react-router-dom";

interface NavProps {
    setCurrentTheme: (theme: keyof typeof Colors) => void;
}

export const Nav = (props: NavProps) => {
    return (
        <div className={NavStyles.Content}>
            <div className={NavStyles.ContentLeft}>
                <Link to="/" className={NavStyles.Text}>OpenCourse</Link>
                <SelectMenu setCurrentTheme={props.setCurrentTheme} />
            </div>
            <div>
                <Link to="/about" className={NavStyles.Text}>About</Link>
            </div>
        </div>
    )
}