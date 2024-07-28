import { Form } from "../Form/Form";
import Header from "../Header/Header";
import MainContentStyles from "./MainContent.module.css"

export const MainContent = () => {
    return (
        <div className={MainContentStyles.Content}>
            <Header />
            <Form />
        </div>
    );
}