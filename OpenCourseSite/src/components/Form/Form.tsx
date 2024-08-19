import { Fieldset } from "@headlessui/react";
import {
  TermSelection,
  CourseSelection,
  SectionSelection,
  ContactForm,
} from "./sections";
import FormStyles from "./Form.module.css";
import { useFormContext } from "../../contexts";
//import { useOrderContext } from "../../contexts";
import { Button } from "components/Button";
import { useForm } from "react-hook-form";

const Form = () => {
  const { selectedTerm, selectedCourses, selectedCheckboxes } =
    useFormContext();
  const { handleSubmit } = useForm();
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const onSubmit = (data: any) => console.log(data);

  return (
    <Fieldset className={FormStyles.Content} onSubmit={handleSubmit(onSubmit)}>
      <TermSelection />
      {selectedTerm && <CourseSelection />}
      {selectedCourses && <SectionSelection />}
      {selectedCheckboxes.length > 0 && <ContactForm />}
      <Button />
    </Fieldset>
  );
};

export default Form;
