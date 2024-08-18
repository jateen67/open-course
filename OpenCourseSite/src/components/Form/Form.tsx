import { Fieldset } from "@headlessui/react";
import { TermSelection, CourseSelection, SectionSelection, ContactForm } from "./sections";
import FormStyles from "./Form.module.css";
import { useFormContext } from "../../contexts";

const Form = () => {
    const { selectedTerm, selectedCourses, selectedCheckboxes } = useFormContext();

    return (
        <Fieldset className={FormStyles.Content}>
            <TermSelection />
            {selectedTerm && <CourseSelection />}
            {selectedCourses && <SectionSelection />}
            {selectedCheckboxes.length > 0 && <ContactForm />}
        </Fieldset>
    );
};

export default Form;