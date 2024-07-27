import { SelectMenu } from "./SelectMenu/SelectMenu"
import Colors from "../styles/ColorSystem"

interface NavProps {
    setCurrentTheme: (theme: keyof typeof Colors) => void;
}

export default function Nav(props: NavProps) {
    return (
        <div className="nav">
            <div className="nav-left">
                <a>OpenCourse</a>
                <SelectMenu setCurrentTheme={props.setCurrentTheme} />
            </div>
            <div>
                <a>About</a>
            </div>
        </div>
    )
}