import RadioGroup from "./RadioGroup";
import FormStyles from "./Form.module.css"
import CourseCombobox from "./Combobox";
import CheckboxGroup from "./CheckboxGroup";

const radioOptions = [
    { id: "r1", value: "fall", label: "Fall 2023" },
    { id: "r2", value: "winter", label: "Winter 2023" },
    { id: "r3", value: "summer", label: "Summer 2024" },
];

const sectionOptions = [
    { sectionType: "lec", sectionNum: "001", crn: 2343, days: ["Tuesday, Thursday"], times: ["11:07 – 23:04"]},
    { sectionType: "lec", sectionNum: "001", crn: 2343, days: ["Tuesday, Thursday"], times: ["11:07 – 23:04"]},
]

const formFields = [
    { id: "name", label: "Name", type: "text" },
    { id: "email", label: "Email", type: "email" },
    { id: "pnum", label: "Phone Number", type: "tel" }
];

const Form = () => (
    <form className={FormStyles.Content}>
        <div className={FormStyles.SectionContent}>
            <h3>Term</h3>
            <RadioGroup options={radioOptions} />
        </div>
        
        <div className={FormStyles.SectionContent}>
            <h3>Course</h3>
            <CourseCombobox />
        </div>
        <div className={FormStyles.SectionContent}>
            <h3>Section</h3>
            <CheckboxGroup sections={sectionOptions} />
        </div>
        <div className={FormStyles.SectionContent}>
            <h3>Contact Info</h3>
            {formFields.map(field => (
                <fieldset key={field.id} className={FormStyles.Fieldset}>
                <label className={FormStyles.Label} htmlFor={field.id}>{field.label}</label>
                <input className={FormStyles.Input} id={field.id} type={field.type} />
                </fieldset>
            ))}
            <button className={FormStyles.Button}>Checkout</button>
        </div>
    </form>
);

export default Form;