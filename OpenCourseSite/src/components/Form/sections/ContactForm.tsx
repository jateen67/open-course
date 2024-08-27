import ContactFormStyles from "./ContactForm.module.css";
import { Field, Label, Input, Transition } from "@headlessui/react";
import { useOrderContext } from "contexts";
import { useState } from "react";
//import { useForm } from "react-hook-form";

const inputs = [
  {
    id: "phone",
    label: "Phone Number",
    type: "text",
    pattern: "^[0-9]{10}$",
    required: true,
    errorMessage: "Please enter a valid phone number.",
  },
];

const ContactForm = () => {
  const { setPhone } = useOrderContext();
  const [errors, setErrors] = useState<{ [key: string]: string }>({});

  const handleChange = (id: string, value: string) => {
    const inputConfig = inputs.find((input) => input.id === id);
    if (inputConfig) {
      const regex = new RegExp(inputConfig.pattern);
      if (!regex.test(value)) {
        setErrors((prevErrors) => ({
          ...prevErrors,
          [id]: inputConfig.errorMessage,
        }));
      } else {
        setErrors((prevErrors) => {
          const newErrors = { ...prevErrors };
          delete newErrors[id];
          return newErrors;
        });
        if (id === "phone") setPhone(value);
      }
    }
  };

  return (
    <>
      {inputs.map((input) => (
        <Field key={input.id}>
          <Label className={ContactFormStyles.Label} htmlFor={input.id}>
            {input.label}
          </Label>
          <Input
            className={ContactFormStyles.Input}
            id={input.id}
            type={input.type}
            pattern={input.pattern}
            required
            onChange={(event) => handleChange(input.id, event.target.value)}
          />
          <Transition
            show={Boolean(errors[input.id])}
            as="div"
            enter="transition-opacity duration-200"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="transition-opacity duration-200"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <span className={ContactFormStyles.Error}>{errors[input.id]}</span>
          </Transition>
        </Field>
      ))}
    </>
  );
};

export default ContactForm;
