import ContactFormStyles from "./ContactForm.module.css"
import { Field, Label, Input } from "@headlessui/react"
import { useOrderContext } from "contexts";
import { useForm } from "react-hook-form";

const inputs = [
    { 
        id: "email", 
        label: "Email", 
        type: "email", 
        pattern: "^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$",
        required: true,
        errorMessage: "Please enter a valid email."
    },
    { 
        id: "phone", 
        label: "Phone Number", 
        type: "text", 
        pattern: "",
        required: true,
        errorMessage: "Please enter a valid phone number."
    }
];

const ContactForm = () => {
    const { setEmail, setPhone } = useOrderContext();

    const handleChange = (id: string, value: string) => {
        if (id === "email") setEmail(value);
        if (id === "phone") setPhone(value);
    };

    return (
        <>
            {inputs.map(input => (
                <Field key={input.id}>
                    <Label className={ContactFormStyles.Label} htmlFor={input.id}>{input.label}</Label>
                    <Input 
                        className={ContactFormStyles.Input} 
                        id={input.id} 
                        type={input.type}
                        pattern={input.pattern}
                        required
                        onChange={(event) => handleChange(input.id, event.target.value)}
                    />
                </Field>
            ))}
        </>
    );
};

export default ContactForm;