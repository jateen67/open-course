import { useState, useEffect } from "react";
import RadioGroup from "../../RadioGroup/RadioGroup";
import { SemesterOption } from "../../../typing"; 
import semestersData from "../../../data/semesters.json";
import { useFormContext } from "../../../contexts";
import FormStyles from "../Form.module.css";

const TermSelection = () => {
    const [radioOptions, setRadioOptions] = useState<SemesterOption[]>([]);
    const { setSelectedTerm, setSelectedCourses, setQuery } = useFormContext();

    const handleTermSelected = (termCode: string) => {
        setSelectedTerm(termCode);
        setSelectedCourses(null);
        setQuery("");
    };

    useEffect(() => {
        setRadioOptions(semestersData as SemesterOption[]);
    }, []);

    return (
        <div className={FormStyles.SectionContent}>
            <h3>Term</h3>
            <RadioGroup options={radioOptions} onChange={handleTermSelected} />
        </div>
    );
};

export default TermSelection;