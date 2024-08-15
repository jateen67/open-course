import { useState ,useEffect } from "react";
import { Fieldset } from "@headlessui/react"
import RadioGroup from "./RadioGroup";
import CourseCombobox from "./Combobox";
import CheckboxGroup from "./CheckboxGroup";
import ContactInfo from "./ContactInfo";
import Button from "./Button"
import FormStyles from "./Form.module.css"
import { Course } from "../../models"
import { SemesterOption } from "../../typing"; 
import semestersData from "../../data/semesters.json";
import { useFormContext } from "../../contexts";

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

    const handleTermSelected = (termCode: string) => {
        setSelectedTerm(termCode);
        setSelectedCourses(null);
        setQuery("");
    };

    const handleCourseSelected = (course : Course | null) => {
        setSelectedCourses(course);
    };

    const handleCheckboxSelected = (courseIds : number[]) => {
        setSelectedCheckboxes(courseIds);
    };

    useEffect(() => {
        setRadioOptions(semestersData as SemesterOption[]);
    }, []);


    return (
        <Fieldset className={FormStyles.Content}>
            <div className={FormStyles.SectionContent}>
                <h3>Term</h3>
                <RadioGroup options={radioOptions} onChange={handleTermSelected} />
            </div>
            {selectedTerm && (
                <div className={FormStyles.SectionContent}>
                    <h3>Course</h3>
                    <CourseCombobox
                        onChange={handleCourseSelected} />
                </div>
            )}
            {selectedCourses && (
                <div className={FormStyles.SectionContent}>
                    <h3>Section</h3>
                    <CheckboxGroup
                        onChange={handleCheckboxSelected}
                    />
                </div>
            )}
            {selectedCheckboxes.length > 0 && (
                <div className={FormStyles.SectionContent}>
                    <h3>Contact Info</h3>
                    <ContactInfo />
                    <Button />
                </div>
            )}
        </Fieldset>
    );
}

export default Form;