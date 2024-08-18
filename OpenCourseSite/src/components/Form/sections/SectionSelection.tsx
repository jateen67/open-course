import CheckboxGroup from "../../CheckboxGroup/CheckboxGroup";
import { useFormContext } from "../../../contexts";
import FormStyles from "../Form.module.css";

const SectionSelection = () => {
    const { setSelectedCheckboxes } = useFormContext();

    const handleCheckboxSelected = (courseIds: number[]) => {
        setSelectedCheckboxes(courseIds);
    };

    return (
        <div className={FormStyles.SectionContent}>
            <h3>Section</h3>
            <CheckboxGroup
                onChange={handleCheckboxSelected}
            />
        </div>
    );
};

export default SectionSelection;