import { useState ,useEffect } from "react";
import RadioGroup from "./RadioGroup";
import CourseCombobox from "./Combobox";
import CheckboxGroup from "./CheckboxGroup";
import FormStyles from "./Form.module.css"
import { Course } from "../../models"
import { SemesterOption } from "../../typing"; 
import semestersData from "../../data/semesters.json";
import { useFormContext } from "../../contexts";

const formFields = [
    { id: "email", label: "Email", type: "email" },
    { id: "pnum", label: "Phone Number", type: "tel" }
];

const Form = () => {
    const [radioOptions, setRadioOptions] = useState<SemesterOption[]>([]);
    const { 
        selectedTerm, 
        selectedCourses, 
        selectedCheckboxes, 
        setSelectedTerm, 
        setSelectedCourses, 
        setSelectedCheckboxes, 
        setQuery 
    } = useFormContext();

    console.log("Form component rendering");

    const handleTermSelected = (termCode : string) => {
        setSelectedTerm(termCode);
        setSelectedCourses(null);
        setQuery("");
    };

    const handleCourseSelected = (course : Course | null) => {
        console.log("Selected course in parent:", course);
        setSelectedCourses(course);
    };

    const handleCheckboxSelected = (courseIds : number[]) => {
        console.log("Selected checkboxes in Form.tsx: ", courseIds)
        setSelectedCheckboxes(courseIds);
    };

    useEffect(() => {
        setRadioOptions(semestersData as SemesterOption[]);
    }, []);


    return (
        <form className={FormStyles.Content}>
            <div className={FormStyles.SectionContent}>
                <h3>Term</h3>
                <RadioGroup options={radioOptions} onChange={handleTermSelected} />
            </div>
            {selectedTerm.length > 0 && (
                <div className={FormStyles.SectionContent}>
                    <h3>Course</h3>
                    <CourseCombobox
                        onChange={handleCourseSelected} />
                </div>
            )}
            { selectedCourses !== null && (
                <div className={FormStyles.SectionContent}>
                    <h3>Section</h3>
                    <CheckboxGroup
                        course={selectedCourses}
                        onChange={handleCheckboxSelected}
                    />
                </div>
            )}
            { selectedCheckboxes?.length > 0 && (
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
            )}
        </form>
    );
}

export default Form;