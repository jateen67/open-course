import { Radio as HeadlessRadio, RadioGroup as HeadlessRadioGroup } from '@headlessui/react'
import { useState } from "react";
import RadioGroupStyles from "./RadioGroup.module.css";
import { SemesterOption } from "../../typing";

interface RadioGroupProps {
    options: SemesterOption[];
    onChange: (value: string) => void;
}

const RadioGroup: React.FC<RadioGroupProps> = ({ options, onChange }) => {
    const [selected, setSelected] = useState<SemesterOption | null>(null);
    
    const handleOnChange = (option: SemesterOption) => {
        setSelected(option);
        onChange(option.value);
    }

    return (
        <HeadlessRadioGroup
            value={selected}
            onChange={handleOnChange}
            aria-label="Term"
            className={RadioGroupStyles.RadioGroup}
        >
            {options.map((option) => (
                <HeadlessRadio
                    key={option.id}
                    value={option}
                    className={`${RadioGroupStyles.Radio} ${(selected?.id === option.id) ? RadioGroupStyles.Checked : ""}`}
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