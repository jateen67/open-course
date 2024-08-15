import { FormProvider } from "contexts";
import Form from "../Form/Form";
import Header from "../Header/Header";
import MainContentStyles from "./MainContent.module.css"

export const MainContent = () => {
    return (
        <div className={MainContentStyles.Content}>
            <Header />
            <FormProvider>
                <Form />
            </FormProvider>
        </div>
    );
}