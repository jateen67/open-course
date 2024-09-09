import { Radio as HeadlessRadio, RadioGroup as HeadlessRadioGroup } from '@headlessui/react'
import RadioGroupStyles from "./RadioGroup.module.css";
import { SemesterOption } from "../../typing";
import { useFormContext } from "../../contexts";

interface RadioGroupProps {
    options: SemesterOption[];
    onChange: (termCode: string) => void;
}

export const RadioGroup: React.FC<RadioGroupProps> = ({ options, onChange }) => {
    const { selectedTerm, setSelectedTerm } = useFormContext();

    const handleOnChange = (id: string) => {
        setSelectedTerm(id);
        onChange(id);
    };

    return (
        <HeadlessRadioGroup
            value={selectedTerm}
            onChange={handleOnChange}
            aria-label="Term"
            className={RadioGroupStyles.RadioGroup}
        >
            {options.map((option) => (
                <HeadlessRadio
                    key={option.id}
                    value={option.id}
                    className={`${RadioGroupStyles.Radio} ${(selectedTerm === option.id) ? RadioGroupStyles.Checked : ""}`}
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