import { Radio as HeadlessRadio, RadioGroup as HeadlessRadioGroup } from '@headlessui/react'
import RadioGroupStyles from "./RadioGroup.module.css";
import { SemesterOption } from "../../typing";
import { useFormContext } from "../../contexts";

interface RadioGroupProps {
    options: SemesterOption[];
    onChange: (value: string) => void;
}

const RadioGroup: React.FC<RadioGroupProps> = ({ options }) => {
    const { selectedTerm, setSelectedTerm } = useFormContext();
    
    const handleOnChange = (option: SemesterOption) => {
        setSelectedTerm(option.value);
    };

    return (
        <HeadlessRadioGroup
            value={selectedTerm || null}
            onChange={handleOnChange}
            aria-label="Term"
            className={RadioGroupStyles.RadioGroup}
        >
            {options.map((option) => (
                <HeadlessRadio
                    key={option.id}
                    value={option}
                    className={`${RadioGroupStyles.Radio} ${(selectedTerm === option.value) ? RadioGroupStyles.Checked : ""}`}
                >
                <div className={RadioGroupStyles.RadioContent}>
                    <div className={RadioGroupStyles.RadioLabel}>
                        <p>{option.label}</p>
                    </div>
                </div>
                </HeadlessRadio>
            ))}
        </HeadlessRadioGroup>
    );
}

export default RadioGroup;