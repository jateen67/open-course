import { useState, useEffect } from "react";
import Checkbox from "./Checkbox";
import CheckboxGroupStyles from "./CheckboxGroup.module.css"
import { Course } from "../../models"
import { useFormContext } from "contexts";

interface CheckboxGroupProps {
    onChange: (courseIds: number[]) => void;
}

const headers = [
    { label: "Section", containerClass: CheckboxGroupStyles.SectionContainer },
    { label: "CRN", containerClass: CheckboxGroupStyles.CRNContainer },
    { label: "Days", containerClass: CheckboxGroupStyles.DayContainer },
    { label: "Time", containerClass: CheckboxGroupStyles.TimeContainer }
];

const fetchCourseInfo = async (termCode: string, courseId: number): Promise<Course[]> => {
    console.log("selectedTerm:", termCode);
    console.log("courseId:", courseId);

    const response = await fetch(`http://localhost:8081/course/${termCode}/${courseId}`);
    if (!response.ok) {
        throw new Error("Failed to fetch course sections");
    }
    return response.json();
};

function formatTime(time: string): string {
    const [hour, minute] = time.split('.');

    return `${hour}:${minute}`;
};

export const CheckboxGroup: React.FC<CheckboxGroupProps> = ({ onChange }) => {
    const { selectedTerm, selectedCourses, selectedCheckboxes, setSelectedCheckboxes } = useFormContext();
    const [sections, setSections] = useState<Course[]>([]);

    useEffect(() => {
        if (!selectedCourses) {
            setSections([]);
            return;
        }

        const fetchData = async () => {
            try {
                const data = await fetchCourseInfo(selectedTerm, selectedCourses.courseId);
                setSections(data);
            } catch (error) {
                console.error("Error fetching course sections:", error);
                setSections([]);
            }
        };

        fetchData();
    }, [selectedTerm, selectedCourses]);

    useEffect(() => {
        console.log("Selected Sections Updated:", selectedCheckboxes);
        onChange(selectedCheckboxes);
    }, [selectedCheckboxes, onChange]);

    const handleOnChange = (checked: boolean, classNumber: number) => {
        setSelectedCheckboxes((prevSelectedCheckboxes: number[]) => {
            const updatedSections = checked
                ? [...prevSelectedCheckboxes, classNumber]
                : prevSelectedCheckboxes.filter((id : number) => id !== classNumber);
    
            return updatedSections;
        });
    };

    return (
        <div className={CheckboxGroupStyles.Container}>
            <div className={CheckboxGroupStyles.HeaderRow}>
                <div className={CheckboxGroupStyles.BoxContainer}></div>
                {headers.map(({ label, containerClass }) => (
                    <div className={containerClass} key={label}>
                        <p className={CheckboxGroupStyles.Heading}>{label}</p>
                    </div>
                ))}
            </div>
            {sections.map(({ classNumber, componentCode, section, classStartTime, classEndTime, ...days }) => (
                <Checkbox
                    key={classNumber}
                    section={`${componentCode} ${section}`}
                    classNumber={classNumber}
                    {...days}
                    times={`${formatTime(classStartTime)} – ${formatTime(classEndTime)}`}
                    onChange={(checked: boolean) => handleOnChange(checked, classNumber)}
                />
            ))}
        </div>
    );
}