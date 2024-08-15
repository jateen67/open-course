import ContactInfoStyles from "./ContactInfo.module.css"
import { Field, Label, Input } from "@headlessui/react"
import { useOrderContext } from "contexts";

const formFields = [
    { id: "email", label: "Email", type: "email" },
    { id: "phone", label: "Phone Number", type: "text" }
];

const ContactInfo = () => {
    const { setEmail, setPhone } = useOrderContext();

    const handleChange = (id: string, value: string) => {
        if (id === "email") setEmail(value);
        if (id === "phone") setPhone(value);
    };

    return (
        <>
            {formFields.map(field => (
                <Field key={field.id}>
                    <Label className={ContactInfoStyles.Label} htmlFor={field.id}>{field.label}</Label>
                    <Input 
                        className={ContactInfoStyles.Input} 
                        id={field.id} 
                        type={field.type}
                        onChange={(event) => handleChange(field.id, event.target.value)}
                    />
                </Field>
            ))}
        </>
    );
};

export default ContactInfo;