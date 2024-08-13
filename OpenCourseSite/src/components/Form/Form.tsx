import { useState ,useEffect } from "react";
import RadioGroup from "./RadioGroup";
import FormStyles from "./Form.module.css"
import CourseCombobox from "./Combobox";
import CheckboxGroup from "./CheckboxGroup";
import { Course, SemesterOption } from "../../typing"; 
import semestersData from "../../data/semesters.json";

const formFields = [
    { id: "name", label: "Name", type: "text" },
    { id: "email", label: "Email", type: "email" },
    { id: "pnum", label: "Phone Number", type: "tel" }
];

const Form = () => {
    const [radioOptions, setRadioOptions] = useState<SemesterOption[]>([]);
    const [checkboxSelected, setCheckboxSelected] = useState(false);
    const [selectedTerm, setSelectedTerm] = useState("");
    const [query, setQuery] = useState<string>("");
    const [selectedCourses, setSelectedCourses] = useState<Course[]>([]);

    const handleTermSelected = (term: string) => {
        setSelectedTerm(term);
        setSelectedCourses([]);
        setQuery("");
    };

    const handleCourseSelected = (courses: Course[]) => {
        setSelectedCourses(courses);
    };

    const handleCheckboxSelected = () => {
        setCheckboxSelected(true);
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
                        term={selectedTerm}
                        query={query}
                        onQueryChange={setQuery}
                        onChange={handleCourseSelected} />
                </div>
            )}
            { selectedCourses.length > 0 && (
                <div className={FormStyles.SectionContent}>
                    <h3>Section</h3>
                    <CheckboxGroup courses={selectedCourses} onChange={handleCheckboxSelected} />
                </div>
            )}
            { checkboxSelected && (
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