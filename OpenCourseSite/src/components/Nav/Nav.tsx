import NavStyles from "./Nav.module.css";
import { Link } from "react-router-dom";

export const Nav = () => {
  return (
    <div className={NavStyles.Content}>
      <div className={NavStyles.ContentLeft}>
        <Link to="/" className={NavStyles.Text}>
          ConUAlerts
        </Link>
      </div>
      <div>
        <Link to="/about" className={NavStyles.Text}>
          About
        </Link>
      </div>
    </div>
  );
};
