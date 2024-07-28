import { SelectMenu } from "../SelectMenu/SelectMenu"
import NavStyles from "./Nav.module.css"
import Colors from "../../styles/ColorSystem"

interface NavProps {
    setCurrentTheme: (theme: keyof typeof Colors) => void;
}

export const Nav = (props: NavProps) => {
    return (
        <div className={NavStyles.Content}>
            <div className={NavStyles.ContentLeft}>
                <p className={NavStyles.Text}>OpenCourse</p>
                <SelectMenu setCurrentTheme={props.setCurrentTheme} />
            </div>
            <div>
                <p className={NavStyles.Text}>About</p>
            </div>
        </div>
    )
}