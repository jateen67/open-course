import CourseCombobox from "../../Combobox/Combobox";
import { Course } from "../../../models";
import { useFormContext } from "../../../contexts";
import FormStyles from "../Form.module.css";

const CourseSelection = () => {
    const { selectedTerm, setSelectedCourses } = useFormContext();

    const handleCourseSelected = (course: Course | null) => {
        setSelectedCourses(course);
    };

    return (
        <div className={FormStyles.SectionContent}>
            <h3>Course</h3>
            <CourseCombobox
                onChange={handleCourseSelected}
            />
        </div>
    );
};

export default CourseSelection;